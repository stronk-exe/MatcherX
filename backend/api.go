package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	store Store
}

type ApiError struct {
	Error string
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("starting api server at", s.addr)
	
	// ** Registering services here **
	userServices := NewUserService(s.store)
	userServices.RegisterRoutes(subrouter)

	postService := NewPostService(s.store)
	postService.RegisterRoutes(subrouter)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}