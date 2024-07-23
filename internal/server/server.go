package server

import "net/http"

type IServer interface {
	SetPort(port int)
	StartServer(handler *http.Handler) error
}
