package main

import (
	"github.com/DiTo04/spexflix/authentication/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	serverAddr  = os.Getenv("AUTHENTICATION_SERVER")
	auPort      = os.Getenv("AUTHENTICATION_PORT")
	contentAddr = os.Getenv("CONTENT_SERVER")
	contentPort = os.Getenv("CONTENT_PORT")
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
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithInsecure())
	auConnection, err := grpc.Dial(addrAndPort, opt...)
	if err != nil {
		log.Fatal("Could not dial up au service,", err)
	}
	auClient := api.NewAuthenticationClient(auConnection)
	r := mux.NewRouter()
	r.NewRoute().Path("/login").Methods("GET").HandlerFunc(createLoginGetHandler(logger))
	r.NewRoute().Path("/login").Methods("POST").HandlerFunc(createLoginPostHandler(auClient, logger))
	homepageHandler, err := getHomePage("/html-templates/homepage.tmpl", contentAddrAndPort)
	if err != nil {
		log.Fatal("Could not parse homepage")
	}
	r.NewRoute().Path("/").Methods("GET").HandlerFunc(homepageHandler)
	http.ListenAndServe("0.0.0.0:8000", r)
}
