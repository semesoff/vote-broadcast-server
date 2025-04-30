package main

import (
	server2 "poll-service/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
