package main

import (
	"github.com/DiTo04/spexflix/authentication/authentication"
	"github.com/DiTo04/spexflix/authentication/server"
	"github.com/DiTo04/spexflix/common/codecs"
	"log"
	"os"
)

var (
	auPort = os.Getenv("AUTHENTICATION_PORT")
)

func main() {
	if auPort == "" {
		auPort = "8080"
	}
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	auth := newAuthenticator()
	s := server.New(auth, logger, codecs.JSON, "0.0.0.0", auPort)
	s.StartServer()
}

func newAuthenticator() server.Authenticator {
	pool := authentication.NewInMemorySessionPool()
	return authentication.CreateAuthenticator(pool, nil)
}
