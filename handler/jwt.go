package handler

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// generates token
func GenerateToken(userID uint) (string, error) {
	// membuat objek payload data / claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 3).Unix(),
		"iat": time.Now().Unix(),
		"userID": userID,
	}

	// generate new token dengan menyisipkan data userID
	// NewWithClaims creates a new Token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SignedString creates and returns a complete, signed JWT.
// The token is signed using the SigningMethod specified in the token.
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}