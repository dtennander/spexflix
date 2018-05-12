package server

import (
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"log"
	"net/http"
	"io"
)

type user struct {
	Username string `json:"username"`
}

type Poster interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type AuthClient struct {
	Client      Poster
	AuthAddress string
	Logger      *log.Logger
	Codec       codecs.Codec
}

func (c *AuthClient) Validate(token string) (string, error) {
	rsp, err := c.Client.Post(c.AuthAddress+"/session/"+token, "", nil)
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
