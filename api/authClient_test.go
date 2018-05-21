package main

import (
	"bytes"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/DiTo04/spexflix/common/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

const (
	testUsername = "testUser"
	testPassword = "testPassword"
	testAddress  = "1.1.1.1:11111111"
	token        = "19y6195ii515i1g51i5"
)

var poster *mocks.MockHttpClient

func setUp(authResponse *http.Response, authError error) *AuthClient {
	poster = &mocks.MockHttpClient{
		Response: authResponse,
		Error:    authError,
	}
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	return &AuthClient{
		Client:      poster,
		Logger:      logger,
		Codec:       codecs.JSON,
		AuthAddress: testAddress,
	}
}

func TestValidateSuccessful(t *testing.T) {
	// Given: a valid response.
	resp := getValidResponse()
	target := setUp(resp, nil)

	// When
	token2, err := target.Login(testUsername, testPassword)

	// Then
	assert.Equal(t, nil, err)
	assert.Equal(t, token, token2)
	assert.Equal(t, testAddress+"/login", poster.LastUrl)
	buff := &bytes.Buffer{}
	buff.ReadFrom(poster.LastBody)
	assert.Equal(t, "{\"username\":\""+testUsername+"\",\"password\":\""+testPassword+"\"}\n", buff.String())
}

func getValidResponse() *http.Response {
	buffer := &bytes.Buffer{}
	buffer.WriteString(token + "\n")
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       mocks.NopCloser{Reader: buffer},
	}
	return resp
}

func TestValidateUnAuthenticated(t *testing.T) {
	// Given: an invalid request
	rsp := getInvalidResponse()
	target := setUp(rsp, nil)

	// When
	token, err := target.Login(testUsername, testPassword+"1")

	// Then
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
	assert.Equal(t, "wrong username or password", err.Error())
}
func getInvalidResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusForbidden,
	}
}
