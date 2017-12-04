package main

import (
	"github.com/DiTo04/spexflix/authentication"
	"github.com/DiTo04/spexflix/authentication/api"
	"testing"
	"time"
	"log"
	"os"
)

const (
	token        = "123"
	testUser     = "testUser"
	testPassword = "testPassword"
)

var logger *log.Logger = log.New(os.Stdout, "INFO: ", 0)

type sessionPoolMock struct {
	isValidId   bool
	createdUser string
}

func (*sessionPoolMock) GetUsername(sessionId string) (string, error) {
	return "testUser", nil
}

func (sp *sessionPoolMock) IsSessionIdValid(sessionId string) bool {
	return sp.isValidId
}
func (sp *sessionPoolMock) CreateSession(username string) (session *authentication.Session, err error) {
	sp.createdUser = username
	return &authentication.Session{
		Username:       username,
		ExpirationDate: time.Now().Add(time.Hour),
		SessionId:      token,
	}, nil
}

type authenticatorMock struct {
	shouldAuthenticateTestUser bool
}

func (a authenticatorMock) Authenticate(user string, password string) bool {
	if user == testUser && password == testPassword {
		return a.shouldAuthenticateTestUser
	} else {
		return false
	}
}

// A valid user should be authenticated
func TestAuService_Authenticate(t *testing.T) {
	sessionPoolMock := &sessionPoolMock{isValidId: true}
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, sessionPoolMock, logger)
	req := &api.AuRequest{SessionToken: token}
	rsp, err := target.Authenticate(nil, req)
	switch {
	case err != nil:
		t.Error("Error from Authenticate(), %v", err)
	case !rsp.IsAuthenticated:
		t.Log("rsp.IsAuthenticated is False")
		t.Fail()
	case rsp.Username != testUser:
		t.Log("Did not recive testUser")
		t.Fail()
	}
}

// A Nonexistent Session should not authenticate
func TestAuService_Authenticate2(t *testing.T) {
	sessionPoolMock := &sessionPoolMock{isValidId: false}
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, sessionPoolMock, logger)
	req := &api.AuRequest{SessionToken: token}
	rsp, err := target.Authenticate(nil, req)
	switch {
	case err != nil:
		t.Error("Error from Authenticate(), %v", err)
	case rsp.IsAuthenticated:
		t.Log("rsp.IsAuthenticated is True")
		t.Fail()
	case rsp.Username == testUser:
		t.Log("Did not recive testUser")
		t.Fail()
	}
}

// Test User should authenticate
func TestAuService_Login(t *testing.T) {
	sessionPoolMock := &sessionPoolMock{}
	authenticatorMock := authenticatorMock{shouldAuthenticateTestUser: true}
	target := createAuService(authenticatorMock, sessionPoolMock, logger)
	req := &api.LoginRequest{Username: testUser, Password: testPassword}
	rsp, err := target.Login(nil, req)
	switch {
	case err != nil:
		t.Error("Error on login")
	case !rsp.IsAuthenticated:
		t.Log("Did not authenticate")
		t.Fail()
	case rsp.SessionToken != token:
		t.Log("Got wrong token")
		t.Fail()
	case sessionPoolMock.createdUser != testUser:
		t.Log("did not create session for testUser")
		t.Fail()
	}
}

// Test User should Not authenticate
func TestAuService_Login2(t *testing.T) {
	sessionPoolMock := &sessionPoolMock{}
	authenticatorMock := authenticatorMock{shouldAuthenticateTestUser: false}
	target := createAuService(authenticatorMock, sessionPoolMock, logger)
	req := &api.LoginRequest{Username: testUser, Password: testPassword}
	rsp, err := target.Login(nil, req)
	switch {
	case err != nil:
		t.Error("Error on login")
	case rsp.IsAuthenticated:
		t.Log("Did authenticate")
		t.Fail()
	case rsp.SessionToken != "":
		t.Log("Got a token")
		t.Fail()
	case sessionPoolMock.createdUser != "":
		t.Log("Did create user session for", sessionPoolMock.createdUser)
		t.Fail()
	}
}
