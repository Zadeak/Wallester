package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) InitRouter() {
	s.Router = mux.NewRouter()
}

func (s *Server) StartServer() {
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
