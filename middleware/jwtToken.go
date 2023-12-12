package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtUtil struct is used to define the jwt functions.
type JwtUtil struct{}

// Claims struct is used to define the claim related details.
type Claims struct {
	Email string
	Role  string
	*jwt.StandardClaims
}

// CreateToken function is used to create a token
func (j *JwtUtil) CreateToken(email string, role string) (string, string, error) {
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
	refreshTokenClaims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("111222"))
	if err != nil {
		return "", "", err
	}

	return strToken, refreshTokenString, nil
}

// ValidateToken function is used to validate the token.
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

		// Check if the access token is about to expire
		expirationTime := time.Unix(claims.ExpiresAt, 0)
		if time.Until(expirationTime) < 5*time.Minute {
			// If the token is about to expire, issue a new access token and send it in the response
			newAccessToken, _, err := j.CreateToken(claims.Email, claims.Role)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to generate new access token",
				})
				return
			}
			c.Header("X-New-Access-Token", newAccessToken)
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

// NewJwtUtil function is used to initialize/instatiate the JwtUtil
func NewJwtUtil() *JwtUtil {
	return &JwtUtil{}
}
