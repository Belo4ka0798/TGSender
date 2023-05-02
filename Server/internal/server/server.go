package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	serverHTTP *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(routes chi.Router) {
	s.serverHTTP = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("Server.host"), viper.GetString("Server.port")),
		Handler:      routes,
		ReadTimeout:  viper.GetDuration("Server.readTimeout") * time.Second,
		WriteTimeout: viper.GetDuration("Server.writeTimeout") * time.Second,
	}
}

func (s *Server) Run() error {
	if err := s.serverHTTP.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
