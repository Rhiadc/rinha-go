package api

import (
	"net/http"

	"github.com/rhiadc/rinha/domain"
)

type PessoaHandler struct {
	service *domain.Service
}

func NewPessoaHandler(service *domain.Service) *PessoaHandler {
	return &PessoaHandler{service: service}
}

/*
{
    "apelido" : "josé",
    "nome" : "José Roberto",
    "nascimento" : "2000-10-01",
    "stack" : ["C#", "Node", "Oracle"]
}
*/
type PessoaRequest struct {
	apelido    string `binding:"Required;MaxSize(32)"`
	nome       string `binding:"Required;MaxSize(100)"`
	nascimento string
	stack      []string `binding:"MaxSize(32)"`
}

func (pr PessoaHandler) Create(w http.ResponseWriter, r *http.Request)    {}
func (pr PessoaHandler) Get(w http.ResponseWriter, r *http.Request)       {}
func (pr PessoaHandler) GetByTerm(w http.ResponseWriter, r *http.Request) {}
func (pr PessoaHandler) Count(w http.ResponseWriter, r *http.Request)     {}
