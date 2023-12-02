package otphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"gobus/entities"
	"gobus/services/interfaces"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis")
	}

}

type otpUser struct {
	Otp  string         `json:"otp"`
	User *entities.User `json:"user"`
}
type OtpHandler struct {
	user interfaces.UserService
}

func (oh *OtpHandler) GenerateOTP(c *gin.Context) {
	user := &entities.User{}
	c.BindJSON(user)
	otp := generateRandomOTP(6)
	otpData := otpUser{
		Otp:  otp,
		User: user,
	}
	data, err := json.Marshal(otpData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to marshal data",
		})
		return
	}
	if err := rdb.Set(ctx, user.Email, data, 5*time.Minute).Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't share data to redis-otp" + err.Error(),
		})
		return
	}
	// if err := rdb.Set(rdb.Context(), user.Email, user, 5*time.Minute).Err(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "couldn't share data to redis-userdata",
	// 	})
	// 	return
	// }

	if err = sendOTPEmail(user.Email, otp); err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "couldn't send otp" + err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "otp has been sent to " + user.Email,
	})

}

// Function to generate a random OTP of the specified length
func generateRandomOTP(length int) string {
	characters := "0123456789"
	otp := make([]byte, length)

	for i := range otp {
		otp[i] = characters[rand.Intn(len(characters))]
	}
	fmt.Print(string(otp))
	return string(otp)
}

func sendOTPEmail(recipientEmail, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Your OTP")

	m.SetBody("text/plain", "Your OTP: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

type verifyOTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func (oh *OtpHandler) VerifyOTP(c *gin.Context) {
	emailotp := &verifyOTP{}
	c.BindJSON(emailotp)
	serializedData, err := rdb.Get(ctx, emailotp.Email).Result()
	if err != nil {
		log.Print("Unable get from redis")
		return
	}
	var retrievedStruct *otpUser
	err = json.Unmarshal([]byte(serializedData), &retrievedStruct)
	if err != nil {
		log.Print("Unable unmarshal the data")
		return
	}

	if emailotp.OTP != retrievedStruct.Otp {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP expired or not valid",
		})
		return
	}

	user, err := oh.user.RegisterUser(retrievedStruct.User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func NewotpHandler(userService interfaces.UserService) *OtpHandler {
	return &OtpHandler{
		user: userService,
	}
}
