package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"../db"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token_str := c.Request.Header.Get("Authorization")
		access_token := strings.SplitAfter(access_token_str, " ")[1]
		if access_token == "" {
			log.Println("Access token not found in request!")
			c.AbortWithStatus(401)
			return
		}

		log.Print(access_token)
		user, err := db.GetCache().HMGet(access_token, "user_id", "user_email").Result()

		if user == nil || err != nil {
			log.Printf("Can't get access token from Redis %s", access_token)
			c.AbortWithStatus(401)
			return
		}
		for _, v := range user {
			log.Print(v)
		}

		c.Next()
	}
}
