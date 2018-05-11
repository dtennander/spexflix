package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type session struct {
	Username       string
	SessionId      string
	ExpirationDate time.Time
}

func createSession(username string) (*session, error) {
	token, err := generateRandomString(32)
	if err != nil {
		return nil, err
	}
	session := session{
		Username:       username,
		SessionId:      string(token),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	}
	return &session, nil
}

func (s *session) getUserName() string {
	return s.Username
}

func (s *session) getSessionId() string {
	return s.SessionId
}

func (s *session) isActive() bool {
	return s.ExpirationDate.After(time.Now())
}

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
