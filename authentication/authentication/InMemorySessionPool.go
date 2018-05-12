package authentication

import "errors"

type inMemorySessionPool map[string]*session

var _ SessionPool = &inMemorySessionPool{}

func NewInMemorySessionPool() inMemorySessionPool {
	return make(inMemorySessionPool)
}

func (sp inMemorySessionPool) GetUsername(sessionId string) (string, error) {
	if v, ok := sp[sessionId]; ok {
		return v.Username, nil
	} else {
		return "", errors.New("not valid SessionId")
	}
}

func (sp inMemorySessionPool) IsSessionIdValid(sessionId string) bool {
	if v, ok := sp[sessionId]; ok {
		return v.isActive()
	} else {
		return false
	}
}

func (sp inMemorySessionPool) CreateSession(username string) (string, error) {
	session, err := createSession(username)
	if err != nil {
		return "", errors.New("could not create session for user: " + username)
	}
	sp[session.SessionId] = session
	return session.SessionId, nil
}
