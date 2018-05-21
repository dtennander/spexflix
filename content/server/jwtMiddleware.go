package server

import (
	"context"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type user struct {
	Username string `json:"username"`
}

type Poster interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type JwtMiddleware struct {
	Client      Poster
	AuthAddress string
	Codec       codecs.Codec
}

func (c *JwtMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := mux.Vars(r)["token"]
	username, err := c.validate(token)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}
	ctx := context.WithValue(r.Context(), "username", username)
	next(rw, r.WithContext(ctx))
}

func (c *JwtMiddleware) validate(token string) (string, error) {
	rsp, err := c.Client.Post("http://"+c.AuthAddress+"/session/"+token, "", nil)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode != http.StatusOK {
		return "", errors.New("could not validate user")
	}
	user := &user{}
	err = c.Codec.Decode(rsp.Body, user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
