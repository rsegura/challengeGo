package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"challenge/config"
	"challenge/server"
	"challenge/db"
)

//var db = make(map[string]string)


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
	db.Init()
	server.Init()
	/*config := config.GetConfig()
	r := setupRouter()
	
	// Listen and Server in 0.0.0.0:8080
	r.Run(config.GetString("server.port"))*/
}