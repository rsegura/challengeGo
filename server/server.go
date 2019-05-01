package server

import (
	"challenge/config"
	"net/http"
	"log"
)

func Init(){
	log.Println("Init Server")
	config := config.GetConfig()
	/*r :=SetupRouter()
	r.Run(config.GetString("server.port"))*/

	r := SetupRouterWithoutFramework()

	http.ListenAndServe(config.GetString("server.port"), r)
}