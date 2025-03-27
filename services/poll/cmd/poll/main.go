package main

import (
	server2 "vote-broadcast-server/services/poll/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
