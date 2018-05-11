package authentication

import (
	"testing"
	"time"
)

const username1 = "user"
const sessionId = "123"

var expirationDate = time.Unix(1111111, 0)

var target = session{
	username1,
	sessionId,
	expirationDate,
}

func TestSession_GetSessionId(t *testing.T) {
	if target.getSessionId() != sessionId {
		t.Log("target.getSessionId did not match")
		t.Fail()
	}
}

func TestSession_GetUserName(t *testing.T) {
	if target.getUserName() != username1 {
		t.Fail()
	}
}

func TestCreateSession(t *testing.T) {
	newSession, err := createSession(username1)
	switch {
	case err != nil:
		t.Log("Could not create session. Error: " + err.Error())
		t.Fail()
	case newSession.Username != username1:
		t.Log("Usernames do not match!")
		t.Fail()
	case newSession.ExpirationDate.Before(time.Now()):
		t.Log("New session is not active!")
		t.Fail()
	}
}

func TestSession_IsActive(t *testing.T) {
	s, err := createSession("asd")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if !s.isActive() {
		t.Log("New session is not active!")
		t.Fail()
	}

}
