package server

import (
	"bytes"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"io"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
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

func (c *AuthClient) Login(username string, password string) (token string, err error) {
	buff := &bytes.Buffer{}
	c.Codec.Encode(buff, &user{Username: username, Password: password})
	rsp, err := c.Client.Post(c.AuthAddress+"/login", "application/json", buff)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode != http.StatusOK {
		return "", errors.New("wrong username or password")
	}
	respBuffer := &bytes.Buffer{}
	respBuffer.ReadFrom(rsp.Body)
	token = strings.Trim(respBuffer.String(), "\n")
	return token, nil
}
