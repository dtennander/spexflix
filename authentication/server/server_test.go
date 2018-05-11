package main

import (
	"github.com/DiTo04/spexflix/authentication/api"
	"log"
	"os"
	"testing"
	"github.com/pkg/errors"
)

const (
	token        = "123"
	testUser     = "testUser"
	testPassword = "testPassword"
)

var logger = log.New(os.Stdout, "INFO: ", 0)

type authenticatorMock struct {
}

func (a authenticatorMock) Login(username string, password string) (string, error) {
	if username == testUser && password == testPassword {
		return token, nil
	} else {
		return "", errors.New("Error")
	}
}

func (a authenticatorMock) AuthenticateSession(sessionToken string) (username *string) {
	if sessionToken == token {
		user := testUser
		return &user
	}
	return nil
}

// A valid user should be authenticated
func TestAuService_Authenticate(t *testing.T) {
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, logger)
	req := &api.AuRequest{SessionToken: token}
	rsp, err := target.Authenticate(nil, req)
	switch {
	case err != nil:
		t.Error("Error from Authenticate(), " +  err.Error())
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
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, logger)
	req := &api.AuRequest{SessionToken: "321"}
	rsp, err := target.Authenticate(nil, req)
	switch {
	case err != nil:
		t.Error("Error from Authenticate(), " + err.Error())
	case rsp.IsAuthenticated:
		t.Log("rsp.IsAuthenticated is True")
		t.Fail()
	case rsp.Username != "":
		t.Log("Did recive a user")
		t.Fail()
	}
}

// Test User should authenticate
func TestAuService_Login(t *testing.T) {
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, logger)
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
	}
}

// Test User should Not authenticate
func TestAuService_Login2(t *testing.T) {
	authenticatorMock := authenticatorMock{}
	target := createAuService(authenticatorMock, logger)
	req := &api.LoginRequest{Username: testUser, Password: testPassword + "wrong password"}
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
	}
}
