package server

import (
	"net"
	"sync"
)

type Node struct {
	Address     string
	Port        string
	Next        *Node
	Prev        *Node
	Route_table map[string]string
	Listener    net.Listener
}

func New_Node(address string, port string) *Node {
	return &Node{
		Address:  address,
		Port:     port,
		Listener: nil,
	}
}

type Network struct {
	mu   sync.Mutex
	Head *Node
}

func New_Network() *Network {
	return &Network{
		Head: nil,
	}
}

func (network *Network) Add_Node_to_Network(address string, port string) {
	network.mu.Lock()
	defer network.mu.Unlock()

	Node_to_Add := New_Node(address, port)

	if network.Head == nil {
		network.Head = Node_to_Add
	} else {
		Node_to_Add.Next = network.Head
		network.Head = Node_to_Add
	}
}

/*
type Coord_Server struct {
	port    string
	address string
}
*/

const (
	COORD_PORT = "8080"
	COORD_ADDR = "127.0.0.1"
)

/*
OLD IDEA

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
*/
