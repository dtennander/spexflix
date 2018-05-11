package authentication

import (
	"testing"
)

var sessionPool = inMemorySessionPool{}

const username2 = "asd"

func TestSessionPool_CreateSession(t *testing.T) {
	// When
	token, err := sessionPool.CreateSession(username2)
	if err != nil {
		t.Error("Could not create session.")
	}
	if len(token) < 10 {
		t.Fail()
	}
}

func TestSessionPoolImpl_GetUsername(t *testing.T) {
	token, err := sessionPool.CreateSession(username2)
	if err != nil {
		t.Error("Could not create session.")
	}
	name, err := sessionPool.GetUsername(token)
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
	s, err := sessionPool.CreateSession(username2)
	if err != nil {
		t.Error("Could not create session.")
	}
	if !sessionPool.IsSessionIdValid(s) {
		t.Fail()
	}
}
