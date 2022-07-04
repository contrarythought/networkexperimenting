package client

import (
	"log"
	"net"
	"strconv"
	"vpn/node"
)

// figure out flag package and design skeleton version cli

func StartCmd() {
	IPAddr, port := GetClientAddr()
	node := node.NewNode(IPAddr, port)
	// form request to coord server
	// Send request to coord server to get routetable

}

func GetClientAddr() (IP string, Port string) {
	conn, err := net.Dial("udp", "8.8.8.8:80") // can use any address because it's udp
	if err != nil {
		log.Fatal(err)
	}
	clientAddr := conn.LocalAddr().(*net.UDPAddr)
	return clientAddr.IP.String(), strconv.Itoa(clientAddr.Port)
}
