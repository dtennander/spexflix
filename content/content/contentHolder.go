package content

import (
	"bytes"
	"io"
)

type Provider struct {
}

const contentMessage = "Välkommen till Spexflix!"

type closableBuffer struct {
	bytes.Buffer
}

func (*closableBuffer) Close() error {
	// No need to close string buffer.
	return nil
}

func (c *Provider) Get(user string) io.ReadCloser {
	buffer := &closableBuffer{}
	buffer.WriteString(contentMessage)
	return buffer
}
