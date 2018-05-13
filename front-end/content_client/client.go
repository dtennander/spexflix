package content_client

import (
	"net/http"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/pkg/errors"
)

type contentView struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	contentServerAddress string
	codec codecs.Codec
}

func (c *Client) Get(token string) (content interface{}, err error) {
	targetUri := c.contentServerAddress + "/" + token + "/content"
	rsp, err := http.Get(targetUri)
	if err != nil {
		return nil, err
	} else if (rsp.StatusCode != http.StatusOK) {
		return nil, errors.New("Could not access content server.")
	}
	content = &contentView{}
	c.codec.Decode(rsp.Body, content)
	return content, nil
}


