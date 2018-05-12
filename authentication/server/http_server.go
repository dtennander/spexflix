package main

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Authenticator interface {
	Login(username string, password string) (token string, err error)
	AuthenticateSession(token string) (username *string)
}

type server struct {
	auth    Authenticator
	logger  *log.Logger
	codec   codecs.Codec
	address string
	port    string
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (s *server) startServer() {
	s.logger.Print("Starting authentivation service on port: " + s.port)
	router := s.createRouter()
	http.ListenAndServe(s.address+":"+s.port, router)
}

func (s *server) createSessionHandler() func(http.ResponseWriter, *http.Request) {
	logRouteSetup(s.logger, "/session/{token}")
	return func(writer http.ResponseWriter, request *http.Request) {
		sessionToken := mux.Vars(request)["token"]
		s.logger.Print("Handeling session/" + sessionToken)
		username := s.auth.AuthenticateSession(sessionToken)
		if username == nil {
			s.logger.Print("Could not authenticate!")
			http.Error(writer, "Could not Authenticate!", http.StatusForbidden)
			return
		}
		writer.WriteHeader(http.StatusOK)
		s.codec.Encode(writer, &user{Username: *username})
	}
}

func (s *server) createLoginHandler() func(http.ResponseWriter, *http.Request) {
	logRouteSetup(s.logger, "/login")
	return func(writer http.ResponseWriter, request *http.Request) {
		user := &user{}
		s.codec.Decode(request.Body, user)
		token, err := s.auth.Login(user.Username, user.Password)
		if err != nil {
			s.logger.Print("Wrong username and password!")
			http.Error(writer, "Wrong username and password!", http.StatusForbidden)
			return
		}
		writer.WriteHeader(http.StatusOK)
		s.codec.Encode(writer, token)
	}
}

func logRouteSetup(logger *log.Logger, path string) {
	logger.Print("Setting up " + path + "endpoint.")
}
