package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Node struct {
	Address     string
	Port        string
	Route_table map[Conn_info]bool
	Listener    net.Listener
}

func New_Node(address string, port string, ip_table map[Conn_info]bool) *Node {
	return &Node{
		Address:     address,
		Port:        port,
		Route_table: ip_table,
		Listener:    nil,
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

/*
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

	// TODO - generate public key...map[address]public_keyo 
	// add new node info to route table
	// this is a placeholder
	network.Head.Route_table[address] = port
}
*/

type Coord_Server struct {
	Port    string
	Address string
	IPTable map[Conn_info]bool
	mu sync.Mutex
}

type Conn_info struct {
	IP string
	Port string
}

const (
	COORD_PORT = "8080"
	COORD_ADDR = "127.0.0.1"
)

func New_coord_server() *Coord_Server {
	return &Coord_Server{
		Port:    COORD_PORT,
		Address: COORD_ADDR,
		IPTable: make(map[Conn_info]bool),
	}
}

var err_failed_connection error = fmt.Errorf("connection failed\n")

func (c *Coord_Server) Start_coord_server() map[Conn_info]bool {
	l, err := net.Listen("tcp", COORD_ADDR+":"+COORD_PORT)
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

		var IP, port string

		go func(connection net.Conn) {
			c.mu.Lock()
			defer c.mu.Unlock()
			conn_str := strings.Split(connection.RemoteAddr().String(), ":")
			
			IP = conn_str[0]
			port = conn_str[1]

			conn_info := Conn_info {
				IP: IP,
				Port: port,
			}

			// store connection's ip address into the server's ip table if it is a new connection
			if c.IPTable[conn_info] == false {
				c.IPTable[conn_info] = true	
			} 
		}(conn)

		return c.IPTable
	}
}
