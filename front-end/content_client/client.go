package content_client

import (
	"bytes"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type contentView struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	ContentServerAddress string
	Codec                codecs.Codec
	Logger               *log.Logger
}

func (c *Client) Get(token string) (content interface{}, err error) {
	targetUri := "http://" + c.ContentServerAddress + "/" + token + "/content"
	rsp, err := http.Get(targetUri)
	if err != nil {
		return nil, err
	} else if rsp.StatusCode != http.StatusOK {
		buffer := &bytes.Buffer{}
		buffer.ReadFrom(rsp.Body)
		c.Logger.Print("Recived error message from content server: " + buffer.String())
		return nil, errors.New("Could not access content server.")
	}
	content = &contentView{}
	c.Codec.Decode(rsp.Body, content)
	return content, nil
}
