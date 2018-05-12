package server

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/DiTo04/spexflix/content/content"
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
	auClient := &AuthClient{
		Codec:       codecs.JSON,
		Logger:      logger,
		AuthAddress: authAddress + ":" + authPort,
		Client:      &http.Client{Timeout: 1 * time.Second},
	}
	provider := &content.Provider{}
	server := &server{
		contentProvider: provider,
		auClient:        auClient,
		codec:           codecs.JSON,
		logger:          logger,
		address:         serverAddress,
		port:            serverPort,
	}
	server.startServer()
}
