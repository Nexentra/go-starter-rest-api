package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userName string) (string, error) {
    // Create a new token object
    token := jwt.New(jwt.SigningMethodHS256)

    // Set token claims
    claims := token.Claims.(jwt.MapClaims)
    claims["user_name"] = userName
    claims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix() // Token will expire in 3 days

    // Sign the token with the secret key
    secretKey := []byte(os.Getenv("JWT_SECRET"))
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
