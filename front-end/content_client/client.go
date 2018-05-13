package content_client

import (
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/pkg/errors"
	"net/http"
)

type contentView struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	ContentServerAddress string
	Codec                codecs.Codec
}

func (c *Client) Get(token string) (content interface{}, err error) {
	targetUri := c.ContentServerAddress + "/" + token + "/content"
	rsp, err := http.Get(targetUri)
	if err != nil {
		return nil, err
	} else if rsp.StatusCode != http.StatusOK {
		return nil, errors.New("Could not access content server.")
	}
	content = &contentView{}
	c.Codec.Decode(rsp.Body, content)
	return content, nil
}
