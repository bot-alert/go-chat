package main

import "chat-application/server"

func main() {
	chatServer := server.Server{
		Port: 8080,
	}
	chatServer.Start()
}
