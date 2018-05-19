package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *server) createRouter() http.Handler {
	router := mux.NewRouter()
	router.NewRoute().
		Path("/session/{token}").
		Methods("POST").
		HandlerFunc(s.handlePostSession)
	router.NewRoute().
		Path("/login").
		Methods("POST").
		HandlerFunc(s.handlePostLogin)

	s.middleware.UseHandler(router)
	return s.middleware
}
