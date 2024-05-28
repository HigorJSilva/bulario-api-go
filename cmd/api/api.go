package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/higorjsilva/goapi/service/medicine"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer  {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	medicineHandler := medicine.NewHandler()
	medicineHandler.RegisterRoutes(subrouter)

	log.Println("Server Runnnig", s.addr)
	
	return http.ListenAndServe(s.addr, router)
}