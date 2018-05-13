package server

import (
	"log"
	"net/http"
	"context"
	"io"
	"bytes"
	"github.com/gorilla/mux"
	"time"
	"github.com/DiTo04/spexflix/common/codecs"
)

type TokenValidator interface {
	Validate(token string) (username string, err error)
}

type ContentProvider interface {
	Get(username string) (content io.ReadCloser)
}

type shutdownLambda func(ctx context.Context) error

type server struct {
	contentProvider ContentProvider
	auClient        TokenValidator
	logger          *log.Logger
	address         string
	port            string
	codec			codecs.Codec
	// internal
	shutdownHook shutdownLambda
}

func New(
	contentProvider ContentProvider,
	auClient TokenValidator,
	logger *log.Logger,
	codec codecs.Codec,
	serverAddress string,
	serverPort string) *server {
		return &server{
			contentProvider:contentProvider,
			auClient:auClient,
			logger:logger,
			address:serverAddress,
			port:serverPort,
			codec:codec,
	}
}

func (server *server) StartServer() {
	server.logger.Print("Starting authentivation service on adress: ", server.address + ":" + server.port)
	router := server.createRoutes()
	httpServer := &http.Server{
		Addr:server.address+":"+server.port,
		Handler:router,
	}
	server.shutdownHook = httpServer.Shutdown
	httpServer.ListenAndServe()
}

func (server *server) StopServer(timeout time.Duration)  {
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	if err := server.shutdownHook(ctx); err != nil {
		panic(err)
	}
}

type Content struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func (server *server) getApiHandler() func(w http.ResponseWriter, r *http.Request) {
	server.logger.Print("Setting up content end point")
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Print("Got Request for content!")
		username := r.Context().Value("username").(string)
		c := server.contentProvider.Get(username)
		defer c.Close()
		buff := &bytes.Buffer{}
		buff.ReadFrom(c)
		content := &Content{Content:buff.String(), Username:username}
		server.codec.Encode(w, content)
	}
}

func (server *server) loggedIn(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	server.logger.Print("Setting up Validation middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		token := mux.Vars(r)["token"]
		username, err := server.auClient.Validate(token)
		if err != nil {
			server.logger.Print("Got rejected request for token: ", token)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		server.logger.Print("Got accepted request for token: ", token)
		ctx := context.WithValue(r.Context(), "username", username)
		handler(w, r.WithContext(ctx))
	}
}
