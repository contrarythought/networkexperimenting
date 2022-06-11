package server

import (
	"fmt"
	"log"
	"net"
)

type Node struct {
	private_key string
	public_key  string
	address     string
	port        string
}

type Coord_Server struct {
	port    string
	address string
}

const (
	COORD_PORT = "8080"
	COORD_ADDR = "127.0.0.1"
)

func New_Node(private_key string, public_key string, address string, port string) *Node {
	return &Node{
		private_key: private_key,
		public_key:  public_key,
		address:     address,
		port:        port,
	}
}

func New_coord_server() *Coord_Server {
	return &Coord_Server{
		port:    COORD_PORT,
		address: COORD_ADDR,
	}
}

var err_failed_connection error = fmt.Errorf("connection failed\n")

func (c *Coord_Server) Start_coord_server() {
	l, err := net.Listen("tcp", COORD_ADDR+COORD_PORT)
	if err != nil {
		log.Fatal(err)
	}

	// listen for connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err_failed_connection)
			conn.Close()
		}

		// TODO - handle connections concurrently
		go func(connection net.Conn) {
			fmt.Println(connection.RemoteAddr().String())
		}(conn)
	}
}
