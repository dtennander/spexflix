package server

import (
	"bytes"
	"context"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io"
	"log"
	"net/http"
	"time"
)

type TokenValidator interface {
	Validate(token string) (username string, err error)
}

type ContentProvider interface {
	Get(username string) (content io.ReadCloser)
}

type shutdownLambda func(ctx context.Context) error

type server struct {
	contentProvider     ContentProvider
	validatorMiddleware negroni.Handler
	logger              *log.Logger
	address             string
	port                string
	codec               codecs.Codec
	// internal
	shutdownHook shutdownLambda
}

func New(
	contentProvider ContentProvider,
	validatorMiddleware negroni.Handler,
	logger *log.Logger,
	codec codecs.Codec,
	serverAddress string,
	serverPort string) *server {
	return &server{
		contentProvider:     contentProvider,
		validatorMiddleware: validatorMiddleware,
		logger:              logger,
		address:             serverAddress,
		port:                serverPort,
		codec:               codec,
	}
}

func (server *server) StartServer() {
	server.logger.Print("Starting authentivation service on adress: ", server.address+":"+server.port)
	router := server.createRoutes()
	httpServer := &http.Server{
		Addr:    server.address + ":" + server.port,
		Handler: router,
	}
	server.shutdownHook = httpServer.Shutdown
	httpServer.ListenAndServe()
}

func (server *server) StopServer(timeout time.Duration) {
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	if err := server.shutdownHook(ctx); err != nil {
		panic(err)
	}
}

func (server *server) createRoutes() *negroni.Negroni {
	r := mux.NewRouter()
	getContentEndpoint := negroni.New(
		server.validatorMiddleware,
		negroni.Wrap(http.HandlerFunc(server.getContent)),
	)
	r.NewRoute().
		Path("/{token}/content").
		Methods("GET").
		Handler(getContentEndpoint)
	r.NewRoute().
		Path("/healthz").
		Methods("GET").
		HandlerFunc(server.healthz)
	n := negroni.Classic()
	n.UseHandler(r)
	return n
}

type Content struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func (server *server) getContent(w http.ResponseWriter, r *http.Request) {
	server.logger.Print("Got Request for content!")
	username := r.Context().Value("username").(string)
	c := server.contentProvider.Get(username)
	defer c.Close()
	buff := &bytes.Buffer{}
	buff.ReadFrom(c)
	content := &Content{Content: buff.String(), Username: username}
	server.codec.Encode(w, content)
}

func (server *server) healthz(w http.ResponseWriter, r *http.Request) {
	content := server.contentProvider.Get("health")
	if content != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no content"))
	}
}
