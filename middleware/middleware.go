package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header not provided")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, "Invalid authorization token format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("UbjV&dxxc3dk6wTU"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid authorization token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, "invalid authorization token")
			c.Abort()
			return
		}

		expTime, ok := claims["exp"].(float64)
		if !ok || float64(time.Now().Unix()) > expTime {
			c.JSON(http.StatusUnauthorized, "Expired authorization token")
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, "Invalid subject claim in token")
			c.Abort()
			return
		}

		c.Set("sub", sub)
		c.Next()
	}
}