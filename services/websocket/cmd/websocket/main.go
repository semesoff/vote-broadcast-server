package main

import (
	server2 "vote-broadcast-server/services/websocket/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
