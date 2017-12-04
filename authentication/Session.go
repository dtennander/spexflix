package authentication

import (
	"time"
	"crypto/rand"
	"encoding/base64"
)

type Session struct {
	Username       string
	SessionId      string
	ExpirationDate time.Time
}

func CreateSession(username string) (*Session, error) {
	token, err := GenerateRandomString(32)
	if err != nil {
		return nil, err
	}
	session := Session{
		Username:       username,
		SessionId:      string(token),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	}
	return &session, nil
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

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}