package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	usename  = "testUser"
	password = "testPassword"
)

type loginMock struct {
	token        string
	err          error
	lastUsername string
	lastPassword string
}

func (m *loginMock) Login(username string, password string) (token string, err error) {
	m.lastPassword = password
	m.lastUsername = username
	return m.token, m.err
}

func TestLoginEndPoint(t *testing.T) {
	// Given
	auClient := &loginMock{token: token, err: nil}
	target := makeHandlePostLogin(auClient, log.New(os.Stderr, "TEST", 0))

	requestDate := &bytes.Buffer{}
	requestDate.WriteString("{\"username\":\"" + usename + "\",\"password\":\"" + password + "\"}")
	req := httptest.NewRequest("POST", "/api/v1/login", requestDate)
	requestRecorder := httptest.NewRecorder()

	// when
	target(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.Equal(t, "\""+token+"\"\n", requestRecorder.Body.String())
	assert.Equal(t, usename, auClient.lastUsername)
	assert.Equal(t, password, auClient.lastPassword)
}
