package server

import (
	"bytes"
	"context"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

const testPort = "8282"
const address = "0.0.0.0"

var testContent = "testContent"

func setUpServerTest() (*server, *auMockClient, *contentMockProvider, *http.Client) {
	auMockClient := &auMockClient{}
	contentMockProvider := &contentMockProvider{content: &testContent}
	server := New(
		contentMockProvider,
		auMockClient,
		log.New(os.Stdout, "TEST: ", 0),
		codecs.JSON,
		address,
		testPort)
	go server.StartServer()
	time.Sleep(100 * time.Millisecond)
	httpClient := &http.Client{Timeout: 1 * time.Second}
	return server, auMockClient, contentMockProvider, httpClient
}

type auMockClient struct {
	username string
	error    error
}

func (c *auMockClient) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if c.error != nil {
		http.Error(rw, c.error.Error(), http.StatusForbidden)
		return
	}
	ctx := context.WithValue(r.Context(), "username", c.username)
	next(rw, r.WithContext(ctx))
}

type contentMockProvider struct {
	lastUser string
	content  *string
}

func (p *contentMockProvider) Get(username string) (content io.ReadCloser) {
	p.lastUser = username
	buff := &bytes.Buffer{}
	buff.WriteString(*p.content)
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
		"{\"username\":\""+testUsername+"\",\"content\":\""+testContent+"\"}\n",
		buffer.String())
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

	req, err := http.NewRequest("GET", getURL(token+"error"), nil)
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

func TestHealthz(t *testing.T) {
	// Given
	target, _, contentProvider, httpClient := setUpServerTest()
	defer target.StopServer(1 * time.Second)
	content := "content"
	contentProvider.content = &content
	req, err := http.NewRequest("GET", "http://"+address+":"+testPort+"/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	// When
	rsp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	// Then
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
}

func TestFailingHealthz(t *testing.T) {
	// Given
	target, _, contentProvider, httpClient := setUpServerTest()
	defer target.StopServer(1 * time.Second)
	contentProvider.content = nil
	req, err := http.NewRequest("GET", "http://"+address+":"+testPort+"/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	// When
	rsp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	// Then
	assert.Equal(t, http.StatusInternalServerError, rsp.StatusCode)
}
