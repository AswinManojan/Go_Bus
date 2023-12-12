package otphandlerprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"gobus/entities"
	"gobus/services/interfaces"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

var rdb *redis.Client
var ctx = context.Background()

// InitRedis is used to initialize the Redis client
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

// otpProvider struct is used to define the otp related information
type otpProvider struct {
	Otp      string                    `json:"otp"`
	Provider *entities.ServiceProvider `json:"provider"`
}

// OtpHandler struct is used to define the otp handler.
type OtpHandler struct {
	provider interfaces.ProviderService
}

// GenerateOTP function is used to generate and send the OTP.
func (oh *OtpHandler) GenerateOTP(c *gin.Context) {
	provider := &entities.ServiceProvider{}
	c.BindJSON(provider)
	// Generate a random 6-digit OTP
	// fmt.Println("Reached here")
	otp := generateRandomOTP(6)
	otpData := otpProvider{
		Otp:      otp,
		Provider: provider,
	}
	data, err := json.Marshal(otpData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to marshal data",
		})
		return
	}
	fmt.Print(rdb)
	// Store the OTP in Redis with an expiration time (e.g., 5 minutes)
	if err := rdb.Set(ctx, provider.Email, data, 5*time.Minute).Err(); err != nil {
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

	if err = sendOTPEmail(provider.Email, otp); err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "couldn't send otp" + err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "otp has been sent to " + provider.Email,
	})

}

// generateRandomOTP function to generate a random OTP of the specified length
func generateRandomOTP(length int) string {
	// Define the characters that can be used in the OTP
	characters := "0123456789"
	otp := make([]byte, length)

	// Use a random source to create the OTP
	for i := range otp {
		otp[i] = characters[rand.Intn(len(characters))]
	}
	fmt.Println(otp)
	return string(otp)
}

// sendOTPEmail function is used to send the otp.
func sendOTPEmail(recipientEmail, otp string) error {
	// Create an email message
	m := gomail.NewMessage()
	m.SetHeader("From", "gobusaswin@gmail.com") // Sender's email address GoBus2000
	m.SetHeader("To", recipientEmail)           // Recipient's email address
	m.SetHeader("Subject", "Your OTP")

	// Set the OTP as the email body
	m.SetBody("text/plain", "Your OTP: "+otp)

	// Send the email using an SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "gobusaswin@gmail.com", "zfej mjdj hhzq lxve")

	// Uncomment this line if you're using a secure SSL/TLS connection
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// verifyOTP struct is used to define otp verification related infomations.
type verifyOTP struct {
	Email string `json:"email"` // Exported field with JSON tag
	OTP   string `json:"otp"`   // Exported field with JSON tag
}

// VerifyOTP fucntion is used to verify the OTP.
func (oh *OtpHandler) VerifyOTP(c *gin.Context) {
	// Retrieve the stored OTP from Redis
	emailotp := &verifyOTP{}
	c.BindJSON(emailotp)
	serializedData, err := rdb.Get(ctx, emailotp.Email).Result()
	if err != nil {
		log.Print("Unable get from redis")
		return
	}
	var retrievedStruct *otpProvider
	err = json.Unmarshal([]byte(serializedData), &retrievedStruct) // Deserialize from JSON
	if err != nil {
		log.Print("Unable unmarshal the data")
		return
	}

	// Compare the user-submitted OTP with the stored OTP
	if emailotp.OTP != retrievedStruct.Otp {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP expired or not valid",
		})
		return
	}

	provider, err := oh.provider.RegisterProvider(retrievedStruct.Provider)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": provider,
	})
}

// NewotpHandler function is used to instatiate the OtpHandler
func NewotpHandler(providerService interfaces.ProviderService) *OtpHandler {
	return &OtpHandler{
		provider: providerService,
	}
}
