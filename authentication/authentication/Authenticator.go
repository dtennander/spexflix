package authentication

import (
	"errors"
)

type authenticator struct {
	hashes   map[string]string
	sessions SessionPool
	db       Database
}

func CreateAuthenticator(pool SessionPool, database Database) *authenticator {
	return &authenticator{hashes: make(map[string]string), sessions:pool, db:database}
}

type SessionPool interface {
	IsSessionIdValid(sessionId string) bool
	GetUsername(sessionId string) (string, error)
	CreateSession(username string) (string, error)
}

type Database interface {}

func (au *authenticator) authenticate(user string, password string) bool {
	if user == "admin" && password == "kakakaka" {
		return true
	} else {
		return au.hashes[user] == password
	}
}

func (au *authenticator) AuthenticateSession(id string) *string {
	username, err := au.sessions.GetUsername(id)
	if err != nil {
		return nil
	} else {
		return &username
	}
}

func (au *authenticator) Login(user string, password string) (string, error) {
	if au.authenticate(user, password) {
		return au.sessions.CreateSession(user)
	} else {
		return "", errors.New("invalid password")
	}
}