package server

import (
	"cloud.google.com/go/storage"
	"time"
	"github.com/DiTo04/spexflix/common/codecs"
	"os"
)

type CloudStorageUrlSigner struct {
	KeyFilePath string
}

type jsonKey struct {
	PrivateKey string `json:"private_key"`
	ClientEmail string `json:"client_email"`
}


func (s *CloudStorageUrlSigner) createUrl(file *storage.ObjectAttrs, method string) (string, error) {
	f, err := os.Open(s.KeyFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	jsonKey := jsonKey{}
	codecs.JSON.Decode(f, &jsonKey)
	options := &storage.SignedURLOptions{
		GoogleAccessID: jsonKey.ClientEmail,
		PrivateKey:     []byte(jsonKey.PrivateKey),
		Method:         method,
		Expires:        time.Now().Add(6 * time.Hour),
	}
	return storage.SignedURL(file.Bucket, file.Name, options)
}
