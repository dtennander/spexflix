package authentication

import (
	"errors"
	jwt2 "github.com/DiTo04/spexflix/common/jwt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JctAuthenticator struct {
	Secret          string
	SessionDuration time.Duration
}

func (a *JctAuthenticator) Login(username string, password string) (token string, err error) {
	if !a.authenticate(username, password) {
		return "", errors.New("invalid password")
	}
	soon := time.Now().Add(a.SessionDuration)
	claims := &jwt2.SessionClaims{
		UserId:         1, //TODO
		Username:       username,
		ExpirationTime: soon.Unix(),
		FeatureFlags:   0,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(a.Secret))
}

func (a *JctAuthenticator) getSecret(token *jwt.Token) (interface{}, error) {
	return []byte(a.Secret), nil
}

func (a *JctAuthenticator) AuthenticateSession(tokenString string) (username *string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt2.SessionClaims{}, a.getSecret)
	if err != nil || !token.Valid {
		return nil
	}
	name := token.Claims.(*jwt2.SessionClaims).Username
	return &name
}

func (jctAuthenticator *JctAuthenticator) authenticate(username string, password string) bool {
	return username == "admin" && password == "kakakaka"
}