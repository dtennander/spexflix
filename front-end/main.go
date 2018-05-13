package main

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/DiTo04/spexflix/front-end/content_client"
	server2 "github.com/DiTo04/spexflix/front-end/server"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	serverAddr  = os.Getenv("AUTHENTICATION_SERVER")
	auPort      = os.Getenv("AUTHENTICATION_PORT")
	contentAddr = os.Getenv("CONTENT_SERVER")
	contentPort = os.Getenv("CONTENT_PORT")
	port        = os.Getenv("PORT")
)

//This server is the gateway onto Spexflix.
//If you are logged in you get passed to the home-page.
//Else you get the log in screen.

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	addrAndPort := serverAddr + ":" + auPort
	logger.Print("Authentication at: " + addrAndPort)
	contentAddrAndPort := contentAddr + ":" + contentPort
	logger.Print("Content server at: " + contentAddrAndPort)
	auClient := &server2.AuthClient{
		Logger:      logger,
		Codec:       codecs.JSON,
		Client:      &http.Client{Timeout: 1 * time.Second},
		AuthAddress: serverAddr + ":" + auPort,
	}
	contentClient := &content_client.Client{
		Codec:                codecs.JSON,
		ContentServerAddress: contentAddr + ":" + contentPort,
	}
	server := server2.New("0.0.0.0", port, logger, auClient, contentClient)
	server.StartServer()
}
