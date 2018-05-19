package server

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func (server *server) createRoutes() *negroni.Negroni {
	r := mux.NewRouter()
	loggedInMux := mux.NewRouter()
	loggedInMux.NewRoute().
		Path("/{token}/content").
		Methods("GET").
		HandlerFunc(server.getApiHandler())
	r.NewRoute().
		Path("/healthz").
		Methods("GET").
		HandlerFunc(server.healthz)

	r.PathPrefix("/{token}/").Handler(negroni.New(
		negroni.HandlerFunc(server.checkLoggedIn),
		negroni.Wrap(loggedInMux)))

	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
