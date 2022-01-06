package server

import "fmt"

type IServer interface {
	Hello(name string) string
	Bye(name string) string
}

type Server struct {
}

func (s Server) Hello(name string) string {
	return fmt.Sprintf("%s:Hello", name)
}

func (s Server) Bye(name string) string {
	return fmt.Sprintf("%s:Bye", name)
}
