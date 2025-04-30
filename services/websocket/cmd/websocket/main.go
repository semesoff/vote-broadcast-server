package main

import (
	server2 "websocket-service/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
