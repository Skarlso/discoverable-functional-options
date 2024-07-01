package pkg

import "fmt"

type Server struct {
	Name    string
	Address string
	Port    int
}

type ServerOpts struct{}

func (ServerOpts) WithName(name string) ServerOptFn {
	return func(s *Server) {
		s.Name = name
	}
}

func (ServerOpts) WithAddress(address string) ServerOptFn {
	return func(s *Server) {
		s.Address = address
	}
}

func (ServerOpts) WithPort(port int) ServerOptFn {
	return func(s *Server) {
		s.Port = port
	}
}

type ServerOptFn func(*Server)

func NewServer(opts ...ServerOptFn) *Server {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	return s
}

func (s *Server) Start() {
	fmt.Printf("started server with name: %s; address: %s; port: %d\n", s.Name, s.Address, s.Port)
}
