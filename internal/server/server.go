package server

type IServer interface {
	SetPort(port int)
	StartServer() error
	CreateServer() error
}
