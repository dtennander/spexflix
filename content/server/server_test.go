package server

import (
	"io"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"time"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
)
const testPort = "8282"
const address = "0.0.0.0"
const testContent = "testContent"

func setUpServerTest() (*server, *auMockClient, *contentMockProvider, *http.Client) {
	auMockClient := &auMockClient{}
	contentMockProvider := &contentMockProvider{content:testContent}
	server := New(
		contentMockProvider,
		auMockClient,
		log.New(os.Stdout, "TEST: ", 0),
		codecs.JSON,
		address,
		testPort)
	go server.StartServer()
	time.Sleep(100 * time.Millisecond)
	httpClient := &http.Client{Timeout:1*time.Second}
	return server, auMockClient, contentMockProvider, httpClient
}

type auMockClient struct {
	lastToken string
	username string
	error error
}

func (c *auMockClient) Validate(token string) (username string, err error) {
	c.lastToken = token
	return c.username, c.error
}

type contentMockProvider struct {
	lastUser string
	content string
}

func (p *contentMockProvider) Get(username string) (content io.ReadCloser) {
	p.lastUser = username
	buff := &bytes.Buffer{}
	buff.WriteString(p.content)
	return ioutil.NopCloser(buff)
}

func TestGetContentWithValidUser(t *testing.T) {
	// Given
	target, auClient, contentProvider, httpClient := setUpServerTest()
	defer target.StopServer(1 * time.Second)
	auClient.username = testUsername
	auClient.error = nil

	req, err := http.NewRequest("GET", getURL(token), nil)
	if err != nil {
		t.Fatal(err)
	}
	// When
	response, err := httpClient.Do(req)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	buffer := &bytes.Buffer{}
	buffer.ReadFrom(response.Body)
	assert.Equal(t,
		"{\"username\":\"" + testUsername + "\",\"content\":\"" + testContent + "\"}\n",
		buffer.String())
	assert.Equal(t, token, auClient.lastToken)
	assert.Equal(t, testUsername, contentProvider.lastUser)
}

func getURL(token string) string {
	return "http://" + address + ":" + testPort + "/" + token + "/content"
}

func TestGetContentWithInvalidUser(t *testing.T) {
	// Given
	target, auClient, contentProvider, httpClient := setUpServerTest()
	defer target.StopServer(1 * time.Second)
	auClient.username = ""
	auClient.error = errors.New("invalid user")

	req, err := http.NewRequest("GET", getURL(token + "error"), nil)
	if err != nil {
		t.Fatal(err)
	}
	// When
	response, err := httpClient.Do(req)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, http.StatusForbidden, response.StatusCode)
	buffer := &bytes.Buffer{}
	buffer.ReadFrom(response.Body)
	assert.Equal(t, "invalid user\n", buffer.String())
	assert.Equal(t, "", contentProvider.lastUser)
}