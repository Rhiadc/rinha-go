package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/rinha/domain"
)

type Server struct {
	service *domain.Service
}

func NewServer(service *domain.Service) *Server {
	server := &Server{service: service}
	r := chi.NewRouter()
	server.router(r)
	s := &http.Server{Addr: "0.0.0.0:3333", Handler: r}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return server
}

func (s *Server) router(r *chi.Mux) {
	pessoaHandler := NewPessoaHandler(s.service)
	r.Post("/pessoas", pessoaHandler.Create)

	//r.Post("/create-repo", repoHandler.CreateRepo)

}
