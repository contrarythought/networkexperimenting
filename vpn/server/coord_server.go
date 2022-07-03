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

type Coord_Server struct {
	Address    string
	Port       string
	RouteTable map[string]bool
	mu         sync.Mutex
}

type PathResolver struct {
	Handlers map[string]http.HandlerFunc
}

const (
	COORD_PORT = "8080"
	COORD_ADDR = "127.0.0.1"
)

func (c *Coord_Server) NewCoordServer() *Coord_Server {
	return &Coord_Server{
		Address:    COORD_ADDR,
		Port:       COORD_PORT,
		RouteTable: make(map[string]bool),
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

func (c *Coord_Server) GetRouteTable(res http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(c.RouteTable)
	if err != nil {
		log.Fatal(err)
	}
	res.Write(b)
}

func (c *Coord_Server) AddConnection(res http.ResponseWriter, req *http.Request) {
	request := strings.Split(req.URL.Path, "/")
	if c.RouteTable[request[1]] == false {
		c.RouteTable[request[1]] = true
	}
}
