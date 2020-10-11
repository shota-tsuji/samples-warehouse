package main

import (
	//"github.com/nqbao/learn-go/chatserver/server"
	"github.com/shota-tsuji/samples-warehouse/go-sandbox/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	s.Start()
}
