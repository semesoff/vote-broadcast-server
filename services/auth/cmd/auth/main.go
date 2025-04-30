package main

import (
	server2 "auth-service/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
