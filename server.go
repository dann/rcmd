package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Server struct {
	Conf *Conf
}

func NewServer(conf *Conf) *Server {
	return &Server{conf}
}

func (s *Server) Run() error {
	log.Printf("Starting Server. address: %s", s.Conf.Addr)
	server := s.initServer()
	return http.ListenAndServe(s.Conf.Addr, server)
}

func (s *Server) initMiddleware(server *martini.ClassicMartini) {
	server.Use(render.Renderer())
	server.Use(martini.Recovery())
}

func (s *Server) initServer() *martini.ClassicMartini {
	server := martini.Classic()
	s.initMiddleware(server)
	s.initRoutes(server)
	return server
}

func (s *Server) initRoutes(server *martini.ClassicMartini) {
	server.Get("/exec", s.createExecuteCommandHandler())
}

func (s *Server) createExecuteCommandHandler() func() (int, string) {
	command := s.Conf.Command
	return func() (int, string) {
		log.Println("Executing command: %s", command)
		commandOutput, err := exec.Command(command).CombinedOutput()
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusOK, string(commandOutput)
	}
}
