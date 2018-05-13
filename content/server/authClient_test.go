package server

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
	username, err := target.Validate(token)

	// Then
	assert.Equal(t, nil, err)
	assert.Equal(t, testUsername, username)
	assert.Equal(t, testAddress+"/session/"+token, poster.LastUrl)
	assert.Equal(t, nil, poster.LastBody)
}

func getValidResponse() *http.Response {
	user := &user{Username: testUsername}
	buffer := &bytes.Buffer{}
	codecs.JSON.Encode(buffer, user)
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
	username, err := target.Validate(token)

	// Then
	assert.Equal(t, "", username)
	assert.NotNil(t, err)
	assert.Equal(t, "could not validate user", err.Error())
}
func getInvalidResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusForbidden,
	}
}
