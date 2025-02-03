package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

type APIServer struct {
	Address string
	Db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		Address: addr,
		Db:      db,
	}
}
 func (s *APIServer) Run() error{
	mux:=http.NewServeMux()
	srv:=&http.Server{
		Addr: s.Address,
		Handler: mux,
	}
	fmt.Println("Server run ",s.Address)
	return srv.ListenAndServe()
 }
