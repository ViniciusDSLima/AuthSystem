package api

import (
	"github.com/ViniciusDSLima/AuthSystem/config"
	_ "github.com/ViniciusDSLima/AuthSystem/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type ServerApi struct {
	addr string
}

type Server struct {
	router *mux.Router
}

func NewApiServer(addr string) *ServerApi {

	return &ServerApi{
		addr: ":" + addr,
	}
}

func (s *ServerApi) Start() error {
	err := config.ConnectDatabase()
	if err != nil {
		return err
	}

	container, err := NewDependencyContainer()
	if err != nil {
		return err
	}

	router := s.initializeServicesUserService(container)

	log.Println("Started Api server at", s.addr)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
	})

	return http.ListenAndServe(s.addr, c.Handler(router))
}
