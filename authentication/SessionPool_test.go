package authentication

import (
	"testing"
)

var sessionPool = SessionPoolImpl{}

const username2 = "asd"

func TestSessionPool_CreateSession(t *testing.T) {
	// When
	s := sessionPool.CreateSession(username2)
	switch {
	case !s.IsActive():
		t.Fail()
	case s.Username != username2:
		t.Fail()
	}
}

func TestSessionPoolImpl_GetUsername(t *testing.T) {
	token := sessionPool.CreateSession(username2)
	name, err := sessionPool.GetUsername(token.SessionId)
	switch {
	case err != nil:
		t.Error("Could not get Username")
	case name != username2:
		t.Log("Expexted", username2, "Got:", name)
		t.Fail()
	}
}

func TestSessionPoolImpl_GetUsername2(t *testing.T) {
	sessionPool.CreateSession(username2)
	name, err := sessionPool.GetUsername("not")
	switch {
	case err == nil:
		t.Error("Could get Username")
	case name != "":
		t.Log("Expexted nothing", "Got:", name)
		t.Fail()
	}
}

func TestSessionPool_IsSessionIdValid2(t *testing.T) {
	if sessionPool.IsSessionIdValid(username2 + "not") {
		t.Fail()
	}
}

func TestSessionPool_IsSessionIdValid(t *testing.T) {
	s := sessionPool.CreateSession(username2)
	if !sessionPool.IsSessionIdValid(s.SessionId) {
		t.Fail()
	}
}
