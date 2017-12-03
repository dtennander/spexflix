package authentication

import (
	"math/rand"
	"time"
)

type Session struct {
	Username       string
	SessionId      string
	ExpirationDate time.Time
}

func CreateSession(username string) *Session {
	session := Session{
		Username:       username,
		SessionId:      string(rand.Int63()),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	}
	return &session
}

func (s *Session) GetUserName() string {
	return s.Username
}

func (s *Session) GetSessionId() string {
	return s.SessionId
}

func (s *Session) IsActive() bool {
	return s.ExpirationDate.After(time.Now())
}
