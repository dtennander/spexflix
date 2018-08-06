package server

import (
	"context"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Authenticator interface {
	Login(username string, password string) (token string, err error)
	AuthenticateSession(token string) (username *string)
}

type Middleware interface {
	UseHandler(handler http.Handler)
	http.Handler
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type server struct {
	auth       Authenticator
	logger     *log.Logger
	codec      codecs.Codec
	address    string
	port       string
	httpServer *http.Server
	middleware Middleware
}

func New(auth Authenticator, logger *log.Logger, codec codecs.Codec, address string, port string, middleware Middleware) *server {
	return &server{
		auth:       auth,
		logger:     logger,
		codec:      codec,
		address:    address,
		port:       port,
		middleware: middleware,
	}
}

func (s *server) StartServer() {
	s.logger.Print("Starting authentivation service on port: " + s.port)
	router := s.createRouter()
	s.httpServer = &http.Server{
		Handler: router,
		Addr:    s.address + ":" + s.port,
	}
	s.httpServer.ListenAndServe()
}

func (s *server) StopServer(timeout time.Duration) {
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	if err := s.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func (s *server) createRouter() http.Handler {
	router := mux.NewRouter()
	router.
		Path("/session/{token}").
		Methods("POST").
		HandlerFunc(s.handlePostSession)
	router.
		Path("/login").
		Methods("POST").
		HandlerFunc(s.handlePostLogin)
	router.
		Path("/healthz").
		HandlerFunc(s.handleHealthz)

	s.middleware.UseHandler(router)
	return s.middleware
}

func (s *server) handlePostSession(writer http.ResponseWriter, request *http.Request) {
	sessionToken := mux.Vars(request)["token"]
	s.logger.Print("Handeling session/" + sessionToken)
	username := s.auth.AuthenticateSession(sessionToken)
	if username == nil {
		s.logger.Print("Could not authenticate!")
		http.Error(writer, "Could not Authenticate!", http.StatusForbidden)
		return
	}
	s.codec.Encode(writer, &user{Username: *username})
}

func (s *server) handlePostLogin(writer http.ResponseWriter, request *http.Request) {
	user := &user{}
	if err := s.codec.Decode(request.Body, user); err != nil {
		s.logger.Print("Could not decode message!")
		http.Error(writer, "Bad request!", http.StatusBadRequest)
		return
	}
	token, err := s.auth.Login(user.Username, user.Password)
	if err != nil {
		s.logger.Print(err)
		http.Error(writer, "Wrong username and password!", http.StatusForbidden)
		return
	}
	s.codec.Encode(writer, token)
}

func (s *server) handleHealthz(writer http.ResponseWriter, request *http.Request) {
	s.codec.Encode(writer, "Everything is fine!")
}
