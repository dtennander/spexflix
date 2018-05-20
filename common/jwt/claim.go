package jwt

import (
	"time"
	"errors"
)

type SessionClaims struct {
	UserId int64 `json:"id"`
	Username string `json:"name"`
	ExpirationTime int64 `json:"exp"`
	FeatureFlags int `json:"ff"`
}


func (sC *SessionClaims) Valid() error {
	expirationTime := time.Unix(sC.ExpirationTime, 0)
	if time.Now().After(expirationTime) {
		errors.New("session has expired")
	}
	return nil
}
