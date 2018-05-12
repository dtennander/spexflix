package mocks

import (
	"io"
	"net/http"
)

type MockHttpClient struct {
	LastUrl         string
	LastContentType string
	LastBody        io.Reader
	Response        *http.Response
	Error           error
}

func (mC *MockHttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	mC.LastUrl = url
	mC.LastContentType = contentType
	mC.LastBody = body
	return mC.Response, mC.Error
}