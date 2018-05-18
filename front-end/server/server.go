package server

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

type server struct {
	http.Server
	logger        *log.Logger
	auClient      Loginer
	contentClient ContentGetter
}

func New(address string, port string, logger *log.Logger, auClient Loginer, contentClient ContentGetter) *server {
	httpServer := http.Server{Addr: address + ":" + port}
	return &server{
		Server:        httpServer,
		logger:        logger,
		auClient:      auClient,
		contentClient: contentClient,
	}
}

func (server *server) StartServer() {
	server.Handler = server.GetRouter()
	server.ListenAndServe()
}

func (server *server) GetRouter() http.Handler {
	r := mux.NewRouter()
	r.NewRoute().
		Path("/login").
		Methods("GET").
		HandlerFunc(createLoginGetHandler(server.logger))
	r.NewRoute().
		Path("/login").
		Methods("POST").
		HandlerFunc(createLoginPostHandler(server.auClient, server.logger))
	homepageHandler, err := server.getHomePage("/html-templates/homepage.tmpl")
	if err != nil {
		log.Fatal("Could not parse homepage")
	}
	r.NewRoute().Path("/browse").Methods("GET").HandlerFunc(homepageHandler)
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(r)
	return n
}
