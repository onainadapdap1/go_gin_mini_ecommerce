package handler

import (
	"fmt"
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

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}