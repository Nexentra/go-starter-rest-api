package jwt

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
        tokenString := c.Request.Header["Authorization"][0]
        if tokenString == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // Verify JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Check signing algorithm
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            // Retrieve secret key for signing
            secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
            return secretKey, nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // Set user ID in context
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        userName, ok := claims["user_name"].(string)
        if !ok {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        c.Set("user_name", userName)
        
        // Continue with request
        c.Next()
    }
}