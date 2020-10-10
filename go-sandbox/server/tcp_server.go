package server

import (
	"errors"
	"github.com/nqbao/learn-go/chatserver/protocol"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

type TcpChatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
}

var (
	UnknownClient = errors.New("Unknown client")
)

func NewServer() *TcpChatServer {
	return &TcpChatServer{
		mutex: &sync.Mutex{},
	}
}

func (s *TcpChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err == nil {
		s.listener = l
	}
	log.Printf("Listening on %v", address)

	return err
}

func (s *TcpChatServer) Close() {
	s.listener.Close()
}

func (s *TcpChatServer) Start() {
	for {
		// XXX: need a way to break the loop
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
		} else {
			client := s.accept(conn) // handle connection
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %v",
		conn.RemoteAddr().String(), len(s.clients)+1)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)

	return client
}

func (s *TcpChatServer) serve(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)
	for {
		cmd, err := cmdReader.Read()
		if err != nil && err != io.EOF {
			log.Printf("Read error: %v", err)
		}
		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.SendCommand:
				go s.Broadcast(protocol.MessageCommand{
					Message: v.Message,
					Name:    client.name,
				})
			case protocol.NameCommand:
				client.name = v.Name
			}
		}
		if err == io.EOF {
			break
		}
	}
}