package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	Address string
	Db      *gorm.DB
}

func NewAPIServer(port string, db *gorm.DB) *APIServer {
	return &APIServer{
		Address: ":" + port,
		Db:      db,
	}
}
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	log.Println("Listening on port", s.Address)
	return http.ListenAndServe(s.Address, router)
}
