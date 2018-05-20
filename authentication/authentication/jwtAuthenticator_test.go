package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	// Given
	username := "admin"
	password := "kakakaka"
	target := &JctAuthenticator{
		SessionDuration: time.Minute,
		Secret:          "Secret"}
	// When
	token, err := target.Login(username, password)
	// Then
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthenticate(t *testing.T) {
	// Given
	username := "admin"
	password := "kakakaka"
	target := &JctAuthenticator{
		SessionDuration: time.Minute,
		Secret:          "Secret",
	}
	token, _ := target.Login(username, password)
	// When
	resultName := target.AuthenticateSession(token)
	// Then
	assert.Equal(t, username, *resultName)
}
