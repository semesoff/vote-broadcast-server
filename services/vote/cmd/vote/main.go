package main

import (
	server2 "vote-broadcast-server/services/vote/pkg/server"
)

func main() {
	server := server2.NewServerManager()
	server.Start()
}
