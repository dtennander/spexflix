package codecs

import (
	"encoding/json"
	"io"
)

type Codec interface {
	Encode(w io.Writer, source interface{}) error
	Decode(r io.Reader, target interface{}) error
}

type jsonCodec struct{}

func (jsonCodec) Encode(w io.Writer, source interface{}) error {
	return json.NewEncoder(w).Encode(source)
}

func (jsonCodec) Decode(r io.Reader, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

var JSON = jsonCodec{}
