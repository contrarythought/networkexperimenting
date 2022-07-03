package node

type Node struct {
	IP      string
	Port    string
	RouteTable map[string]bool
}

func NewNode(IP, Port string, RouteTable map[string]bool) *Node {
	return &Node{
		IP: IP,
		Port: Port,
		RouteTable: RouteTable,
	}
}

