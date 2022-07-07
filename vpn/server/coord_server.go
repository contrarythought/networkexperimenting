package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
)

type CoordServer struct {
	Address    string
	Port       string
	RouteTable map[NodeEntry]bool
	mu         sync.Mutex	
}

type NodeEntry struct {
	PrivateIP string
	PublicIP string
}

type PathResolver struct {
	Handlers map[string]http.HandlerFunc
}

const (
	COORD_PORT = "8080"
	COORD_ADDR = "127.0.0.1"
)

func (c *CoordServer) NewCoordServer() *CoordServer {
	return &CoordServer{
		Address:    COORD_ADDR,
		Port:       COORD_PORT,
		RouteTable: make(map[NodeEntry]bool),
	}
}

func (p *PathResolver) NewPathResolver() *PathResolver {
	return &PathResolver{
		Handlers: make(map[string]http.HandlerFunc),
	}
}

func (p *PathResolver) AddPath(request string, handler http.HandlerFunc) {
	p.Handlers[request] = handler
}

func (p *PathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	for pattern, handler := range p.Handlers {
		if matched, err := path.Match(pattern, check); matched && err == nil {
			handler(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	http.NotFound(res, req)
}

func (c *CoordServer) GetRouteTable(res http.ResponseWriter, req *http.Request) {
	b, err := json.MarshalIndent(c.RouteTable, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	res.Write(b)
}

func (c *CoordServer) AddConnection(res http.ResponseWriter, req *http.Request) {
	request := strings.Split(req.URL.Path, "/")
	publicIP := req.RemoteAddr
	node := NodeEntry {
		PrivateIP: request[1],
		PublicIP: publicIP,
	}
	if c.RouteTable[node] == false {
		c.RouteTable[node] = true
	}
}
