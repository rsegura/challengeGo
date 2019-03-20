package server

import "challenge/config"

func Init(){
	config := config.GetConfig()
	r :=SetupRouter()

	r.Run(config.GetString("server.port"))
}