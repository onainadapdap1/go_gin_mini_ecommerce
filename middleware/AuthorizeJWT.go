package middleware

import (
	"fmt"
	"go_gin_mini_ecommerce/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthorizeJWT -> to authorize JWT token
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found",
			})
		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := handler.ValidateToken(tokenString); err != nil {
			fmt.Println("token", tokenString, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token",
			})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "No valid signing method",
				})
			} else {
				if token.Valid {
					c.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "token error",
					})
				}
			}
		}

	}
}