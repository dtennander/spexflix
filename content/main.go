package main

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/DiTo04/spexflix/content/content"
	"github.com/DiTo04/spexflix/content/server"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	authAddress   = os.Getenv("AUTHENTICATION_SERVER")
	authPort      = os.Getenv("AUTHENTICATION_PORT")
	serverAddress = os.Getenv("SERVER_ADDRESS")
	serverPort    = os.Getenv("SERVER_PORT")
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	auClient := &server.AuthClient{
		Codec:       codecs.JSON,
		Logger:      logger,
		AuthAddress: authAddress + ":" + authPort,
		Client:      &http.Client{Timeout: 1 * time.Second},
	}
	provider := &content.Provider{}
	server.New(
		provider, auClient, logger, serverAddress, serverPort).
		StartServer()
}
