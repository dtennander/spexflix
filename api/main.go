package main

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
	"io/ioutil"
	"strings"
)

const apiPrefix = "/api/v1"

var (
	usersServiceUrl = os.Getenv("USER_SERVICE")
	contentServiceUrl = os.Getenv("CONTENT_SERVICE")
	authService = os.Getenv("AUTHENTICATION_SERVICE")
	port        = os.Getenv("PORT")
)

type endpoint func(w http.ResponseWriter, r *http.Request)

func main() {
	logger := log.New(os.Stderr, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	httpClient := &http.Client{
		Timeout: 500 * time.Millisecond,
	}
	auClient := &AuthClient{
		AuthAddress: authService,
		Client:      httpClient,
		Codec:       codecs.JSON,
		Logger:      logger,
	}

	r := mux.NewRouter()
	apiRouter := r.PathPrefix(apiPrefix).Subrouter()

	//Api Calls
	apiRouter.NewRoute().
		Path("/login").
		Methods("POST").
		HandlerFunc(makeHandlePostLogin(auClient, logger))

	apiRouter.PathPrefix("/users").HandlerFunc(proxyTo(usersServiceUrl))
	apiRouter.PathPrefix("/movies").HandlerFunc(proxyTo(contentServiceUrl))

	r.PathPrefix(apiPrefix).Handler(apiRouter)
	r.HandleFunc("/healthz", healthz)
	n := negroni.Classic()
	n.UseHandler(r)
	if port == "" {
		port = "8080"
	}
	logger.Print("Starting api server on:", port)
	http.ListenAndServe("0.0.0.0:"+port, n)
}

func proxyTo(backendAddress string) (func(writer http.ResponseWriter, request *http.Request)) {
	s := strings.Split(backendAddress, "://")
	log.Print(s)
	host := s[1]
	protocol := s[0]
	return func (writer http.ResponseWriter, request *http.Request) {
		url := request.URL
		url.Path = strings.TrimPrefix(url.Path, apiPrefix)
		url.Host = host
		url.Scheme = protocol
		proxyReq, err := http.NewRequest(request.Method, url.String(), request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		proxyReq.Header = request.Header
		client := &http.Client{}
		rsp, err := client.Do(proxyReq)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}
		defer rsp.Body.Close()
		body, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}

		writer.Write(body)
		writer.WriteHeader(rsp.StatusCode)
	}
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
