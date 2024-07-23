package server

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	port int
	Mux  *http.ServeMux
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Mux: http.NewServeMux(),
	}
}

func (hs *HttpServer) SetPort(port int) {
	hs.port = port
}

func (hs *HttpServer) StartServer(handler http.Handler) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", hs.port), handler)
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}
