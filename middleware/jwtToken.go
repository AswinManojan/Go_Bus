package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtUtil struct{}

type Claims struct {
	Email string
	Role  string
	*jwt.StandardClaims
}

func (j *JwtUtil) CreateToken(email string, role string) (string, error) {
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte("111222"))
	if err != nil {
		panic("Error creating token")
	}
	return strToken, nil
}
func (j *JwtUtil) ValidateToken(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token not valid",
			})
			return
		}
		tokenString = string([]byte(tokenString[7:]))
		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("111222"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token not valid",
			})
			return
		}
		if claims.Role != role || !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "user-admin mismatch-- unauthorized acces-- access denied",
			})
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}


func NewJwtUtil() *JwtUtil {
	return &JwtUtil{}
}
