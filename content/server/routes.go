package server

import "github.com/gorilla/mux"

func (server *server) createRoutes() *mux.Router {
	r := mux.NewRouter()
	r.NewRoute().
		Path("/{token}/content").
		Methods("GET").
		HandlerFunc(server.loggedIn(server.getApiHandler()))
	return r
}
