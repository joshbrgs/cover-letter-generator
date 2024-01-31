package server

import "time"

type Server struct {
	Host    string
	Port    string
	Timeout time.Duration
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
    return func(s *Server) {
        s.Host = host
    }
}

func WithPort(port string) ServerOption {
return func(s *Server) {
        s.Port = port
    }
}

func WithTimeout(timeout time.Duration) ServerOption {
return func(s *Server) {
        s.Timeout = timeout
    }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{
        Host: "localhost",
        Port: "8080",
        Timeout: 30 * time.Second,
    }

    for _, opt := range opts {
        opt(s)
    }

    return s
}