package server

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/DiTo04/spexflix/content/content"
	"log"
	"net/http"
	"context"
)

type contentBody struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type server struct {
	contentProvider *content.Provider
	auClient        *AuthClient
	logger          *log.Logger
	codec           codecs.Codec
	address         string
	port            string
}

func (server *server) startServer() {
	server.logger.Print("Starting authentivation service on port:", server.port)
	router := server.createRoutes()
	http.ListenAndServe(server.address+":"+server.port, router)
}

func (server *server) getApiHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Print("Got Request for content!")
		username := r.Context().Value("username").(string)
		c := server.contentProvider.GetContentForUser(username)
		responseBody := &contentBody{Username: username, Content: c}
		w.WriteHeader(http.StatusOK)
		server.codec.Encode(w, responseBody)
	}
}

func (server *server) loggedIn(
	handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("sessionToken")
		username, err := server.auClient.Validate(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		handler(w, r.WithContext(ctx))
	}
}
