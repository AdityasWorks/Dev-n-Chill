package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	r := gin.Default()

	secret := []byte("secretkey")


	r.POST("/login", func(c *gin.Context){
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user" : "testuser",
			"exp" : time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, _ := token.SignedString(secret)
		c.JSON(200,gin.H{"token": tokenString})
	})


	authGroup := r.Group("/api")
	authGroup.Use(func (c *gin.Context){
		tokenStr == c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatus(401);
			return
		}

		claims := jwt.MapClaims{}
		_,err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token){
			if err != nil {
				c.AbortWithStatus(401);
				return
			}
			c.Next()
		})
		authGroup.GET("/environments", func(c *gin.Context){
			c.JSON(200,gin.H{"environments" : []string{"dev", "staging",}})

		})
	})


	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "OK"})
	})
	r.Run(":8080")
}