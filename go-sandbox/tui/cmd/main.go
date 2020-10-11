package main

import (
	"flag"
	//"github.com/nqbao/learn-go/chatserver/client"
	"github.com/shota-tsuji/samples-warehouse/go-sandbox/client"
	"github.com/shota-tsuji/samples-warehouse/go-sandbox/tui"
	//"github.com/nqbao/learn-go/chatserver/tui"
	"log"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")
	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()
	tui.StartUi(client)
}
