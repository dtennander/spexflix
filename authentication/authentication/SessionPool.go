package authentication

import "errors"

type SessionPool interface {
	IsSessionIdValid(sessionId string) bool
	GetUsername(sessionId string) (string, error)
	CreateSession(username string) (session *Session, err error)
}

type SessionPoolImpl map[string]*Session

func (sp SessionPoolImpl) GetUsername(sessionId string) (string, error) {
	if v, ok := sp[sessionId]; ok {
		return v.Username, nil
	} else {
		return "", errors.New("Not valid SessionId")
	}
}

func (sp SessionPoolImpl) IsSessionIdValid(sessionId string) bool {
	if v, ok := sp[sessionId]; ok {
		return v.IsActive()
	} else {
		return false
	}
}

func (sp SessionPoolImpl) CreateSession(username string) (session *Session, err error) {
	session, err = CreateSession(username)
	sp[session.SessionId] = session
	return session, nil
}
