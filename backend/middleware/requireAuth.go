package middleware

import (
	"fmt"
	"log"
	"main/db"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	log.Printf("RequireAuth middleware from %s\n", c.Request.URL.Path)

	cookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		log.Printf("Error getting token cookie: %v\n", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := validateToken(cookie.Value)
	if err != nil {
		log.Printf("Error validating token: %v\n", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = setUserInContextFromToken(c, token)
	if err != nil {
		log.Printf("Error setting user: %v\n", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_SALT")), nil
	})
}

func setUserInContextFromToken(c *gin.Context, token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().After(time.Unix(int64(exp), 0)) {
		return fmt.Errorf("token expired or invalid expiration time")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return fmt.Errorf("invalid user ID")
	}

	user, err := db.GetUser(strconv.FormatFloat(sub, 'f', -1, 64))
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	c.Set("user", user)
	return nil
}
