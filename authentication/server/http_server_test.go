package server

import (
	"bytes"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/urfave/negroni"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	token         = "123"
	testUser      = "testUser"
	testPassword  = "testPassword"
	port          = "8080"
	address       = "0.0.0.0"
	ServerAddress = "http://" + address + ":" + port
)

type authenticatorMock struct {
}

func (a authenticatorMock) Login(username string, password string) (string, error) {
	if username == testUser && password == testPassword {
		return token, nil
	} else {
		return "", errors.New("error")
	}
}

func (a authenticatorMock) AuthenticateSession(sessionToken string) (username *string) {
	if sessionToken == token {
		user := testUser
		return &user
	}
	return nil
}

func setUp() (*server, *http.Client) {
	authenticatorMock := authenticatorMock{}
	logger := log.New(os.Stdout, "INFO: ", 0)
	s := &server{
		auth:       authenticatorMock,
		codec:      codecs.JSON,
		logger:     logger,
		port:       port,
		address:    address,
		middleware: negroni.Classic(),
	}
	go s.StartServer()
	time.Sleep(100 * time.Millisecond)
	return s, &http.Client{Timeout: 1 * time.Second}
}

func TestSessionEndpointSuccess(t *testing.T) {
	// Given
	server, client := setUp()
	defer server.StopServer(1 * time.Second)

	req, err := http.NewRequest("POST", ServerAddress+"/session/"+token, nil)
	if err != nil {
		t.Fatal("Could not create request: " + err.Error())
	}
	// When
	res, err := client.Do(req)
	// Then
	if err != nil {
		t.Fatal("Could not send request: " + err.Error())
	}
	validateResponse(
		t, res, http.StatusOK, "{\"username\":\""+testUser+"\"}\n")
}

func validateResponse(t *testing.T, res *http.Response, statusCode int, expectedResponse string) {
	if res.StatusCode != statusCode {
		t.Log("Wrong status code:", res.StatusCode)
		t.Fail()
	}
	responseBody := toString(res.Body)
	if responseBody != expectedResponse {
		t.Log("Wrong response:", responseBody)
		t.Fail()
	}
}

func toString(reader io.Reader) string {
	buff := new(bytes.Buffer)
	buff.ReadFrom(reader)
	return buff.String()
}

func TestSessionEndpointFailiure(t *testing.T) {
	// Given
	server, client := setUp()
	defer server.StopServer(1 * time.Second)

	req, err := http.NewRequest("POST", ServerAddress+"/session/"+token+"1", nil)
	if err != nil {
		t.Fatal("Could not create request: " + err.Error())
	}
	// When
	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Could not send request: " + err.Error())
	}
	// Then
	validateResponse(
		t, res, http.StatusForbidden, "Could not Authenticate!\n")
}

func TestLoginSuccess(t *testing.T) {
	// Given
	server, client := setUp()
	defer server.StopServer(1 * time.Second)

	buffer := new(bytes.Buffer)
	codecs.JSON.Encode(buffer, user{Username: testUser, Password: testPassword})
	req, err := http.NewRequest("POST", ServerAddress+"/login", buffer)
	if err != nil {
		t.Fatal("Could not create request: " + err.Error())
	}
	// When
	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Could not send request: " + err.Error())
	}
	// Then
	validateResponse(
		t, res, http.StatusOK, "\""+token+"\"\n")
}

func TestLoginFailure(t *testing.T) {
	// Given
	server, client := setUp()
	defer server.StopServer(1 * time.Second)

	buffer := new(bytes.Buffer)
	wrongPassword := testPassword + "1"
	codecs.JSON.Encode(buffer, user{Username: testUser, Password: wrongPassword})
	req, err := http.NewRequest("POST", ServerAddress+"/login", buffer)
	if err != nil {
		t.Fatal("Could not create request: " + err.Error())
	}
	// When
	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Could not send request: " + err.Error())
	}
	// Then
	validateResponse(
		t, res, http.StatusForbidden, "Wrong username and password!\n")
}
