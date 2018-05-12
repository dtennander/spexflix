package server

import "github.com/gorilla/mux"

func (server *server) createRoutes() *mux.Router {
	r := mux.NewRouter()
	r.NewRoute().
		Methods("GET").
		Path("/content/{token}").
		HandlerFunc(server.loggedIn(server.getApiHandler()))
	return r
}
