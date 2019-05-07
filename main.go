package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"challenge/config"
	"challenge/server"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context){
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
	

}

func main() {
	config.Init("development")
	config := config.GetConfig()
	a := server.App{}
	a.Init()
	a.Run(config.GetString("server.port"))
}