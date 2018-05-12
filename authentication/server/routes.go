package main

import "github.com/gorilla/mux"

func (s *server) createRouter() *mux.Router {
	router := mux.NewRouter()
	router.NewRoute().
		Path("/session/{token}").
		Methods("POST").
		HandlerFunc(s.createSessionHandler())
	router.NewRoute().
		Path("/login").Methods("POST").HandlerFunc(s.createLoginHandler())
	return router
}
