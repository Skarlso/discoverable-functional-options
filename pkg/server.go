package pkg

import "fmt"

// Server contains the server's setup values.
type Server struct {
	Name    string
	Address string
	Port    int
}

// ServerOpts groups all valid options for the server together.
type ServerOpts struct{}

// WithName adds an options to define the name of the server.
func (ServerOpts) WithName(name string) ServerOptFn {
	return func(s *Server) {
		s.Name = name
	}
}

// WithAddress adds the address of the server.
func (ServerOpts) WithAddress(address string) ServerOptFn {
	return func(s *Server) {
		s.Address = address
	}
}

// WithPort defines the port on which the server should run.
func (ServerOpts) WithPort(port int) ServerOptFn {
	return func(s *Server) {
		s.Port = port
	}
}

// ServerOptFn defines the function type to configure a server.
type ServerOptFn func(*Server)

// NewServer creates a new configured server.
func NewServer(opts ...ServerOptFn) *Server {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	return s
}

// Start starts the server and outputs the configured options.
func (s *Server) Start() {
	fmt.Printf("started server with name: %s; address: %s; port: %d\n", s.Name, s.Address, s.Port)
}
