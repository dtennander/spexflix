package main

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	authAddress = os.Getenv("AUTHENTICATION_SERVER")
	authPort    = os.Getenv("AUTHENTICATION_PORT")
	port        = os.Getenv("PORT")
)

type endpoint func(w http.ResponseWriter, r *http.Request)

func main() {
	logger := log.New(os.Stderr, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	httpClient := &http.Client{
		Timeout: 500 * time.Millisecond,
	}
	auClient := &AuthClient{
		AuthAddress: "http://" + authAddress + ":" + authPort,
		Client:      httpClient,
		Codec:       codecs.JSON,
		Logger:      logger,
	}

	// health calls
	r := mux.NewRouter()
	r.NewRoute().
		Path("/healthz").
		Methods("GET").
		HandlerFunc(healthz)
	// Api calls
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.NewRoute().
		Path("/login").
		Methods("POST").
		HandlerFunc(makeHandlePostLogin(auClient, logger))

	r.PathPrefix("/api/v1").Handler(apiRouter)
	n := negroni.Classic()
	n.UseHandler(r)
	if port == "" {
		port = "8080"
	}
	logger.Print("Starting api server on:", port)
	http.ListenAndServe("0.0.0.0:"+port, n)
}

func healthz(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Happy"))
}

type Logginer interface {
	Login(username string, password string) (token string, err error)
}

func makeHandlePostLogin(auClient Logginer, logger *log.Logger) endpoint {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &user{}
		if err := codecs.JSON.Decode(r.Body, user); err != nil {
			http.Error(w, "Could not parse form", http.StatusBadRequest)
			return
		}
		token, err := auClient.Login(user.Username, user.Password)
		if err != nil {
			logger.Print(err.Error())
			http.Error(w, "Could not authenticate", http.StatusInternalServerError)
			return
		}
		codecs.JSON.Encode(w, token)
	}
}
