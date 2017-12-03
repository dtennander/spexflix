package authentication

import (
	"testing"
	"time"
)

const username1 = "user"
const sessionId = "123"

var expirationDate = time.Unix(1111111, 0)

var target = Session{
	username1,
	sessionId,
	expirationDate,
}

func TestSession_GetSessionId(t *testing.T) {
	if target.GetSessionId() != sessionId {
		t.Log("target.GetSessionId did not match")
		t.Fail()
	}
}

func TestSession_GetUserName(t *testing.T) {
	if target.GetUserName() != username1 {
		t.Fail()
	}
}

func TestCreateSession(t *testing.T) {
	newSession := CreateSession(username1)
	switch {
	case newSession.Username != username1:
		t.Log("Usernames do not match!")
		t.Fail()
	case newSession.ExpirationDate.Before(time.Now()):
		t.Log("New session is not active!")
		t.Fail()
	}
}

func TestSession_IsActive(t *testing.T) {
	s := CreateSession("asd")
	if !s.IsActive() {
		t.Log("New Session is not active!")
		t.Fail()
	}

}
