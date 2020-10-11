module cmd

replace github.com/nqbao/learn-go/chatserver/client => ../../client

replace github.com/nqbao/learn-go/chatserver/tui => ../

go 1.14

require (
	github.com/marcusolsson/tui-go v0.4.0 // indirect
	github.com/nqbao/learn-go v0.0.0-20200624064038-8d7dcb73f398 // indirect
	github.com/nqbao/learn-go/chatserver/client v0.0.0-00010101000000-000000000000 // indirect
	github.com/nqbao/learn-go/chatserver/tui v0.0.0-00010101000000-000000000000 // indirect
)
