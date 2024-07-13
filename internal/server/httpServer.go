package server

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	port   int
	server http.Server
}

func NewHttpServer() IServer {
	return &HttpServer{}
}

func (hs *HttpServer) SetPort(port int) {
	hs.port = port
}

func (hs *HttpServer) CreateServer() error {
	if hs.port == 0 {
		return fmt.Errorf("port not set")
	}

	hs.server = http.Server{
		Addr: fmt.Sprintf(":%d", hs.port),
	}

	return nil
}

func (hs *HttpServer) StartServer() error {
	err := hs.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}
