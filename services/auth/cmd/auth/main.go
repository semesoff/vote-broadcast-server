package main

import (
	server2 "vote-broadcast-server/services/auth/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
