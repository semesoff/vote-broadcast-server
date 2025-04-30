package main

import (
	server2 "vote-service/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
