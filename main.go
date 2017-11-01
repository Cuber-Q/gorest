package main

import (
	"gorest/server"
	"gorest/handler"
)

func main() {
	restServer := server.RestServer{}
	restServer.SetPort("8081")
	restServer.AddRouter("/hi", &handler.HiHandler{}, "SayHi")
	restServer.Start()
}
