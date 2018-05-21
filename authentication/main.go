package main

import (
	"github.com/DiTo04/spexflix/authentication/authentication"
	"github.com/DiTo04/spexflix/authentication/server"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/urfave/negroni"
	"log"
	"os"
	"time"
)

var (
	auPort = os.Getenv("AUTHENTICATION_PORT")
	jwtSecret = os.Getenv("JWT_SECRET")
)

func main() {
	if auPort == "" {
		auPort = "8080"
		jwtSecret = "IN DEVELOP I HOPE"
	}
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	auth := newAuthenticator()
	middleware := negroni.Classic()
	s := server.New(auth, logger, codecs.JSON, "0.0.0.0", auPort, middleware)
	s.StartServer()
}

func newAuthenticator() server.Authenticator {
	au := &authentication.JctAuthenticator{
		Secret:          jwtSecret,
		SessionDuration: 7 * 24 * time.Hour,
	}
	return au
}
