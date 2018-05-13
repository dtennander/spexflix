package server

import (
	"github.com/gorilla/mux"
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
	homepageHandler, err := getHomePage("/html-templates/homepage.tmpl", server.contentClient)
	if err != nil {
		log.Fatal("Could not parse homepage")
	}
	r.NewRoute().Path("/").Methods("GET").HandlerFunc(homepageHandler)
	return r
}
