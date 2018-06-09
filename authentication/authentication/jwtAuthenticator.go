package authentication

import (
	"errors"
	jwt2 "github.com/DiTo04/spexflix/common/jwt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type HashStore interface {
	GetHash(i int64) ([]byte, error)
}

type UserService interface {
	GetUserId(email string) (int64, error)
}

type JwtAuthenticator struct {
	Secret          string
	SessionDuration time.Duration
	HashStore       HashStore
	UserService     UserService
}

func (a *JwtAuthenticator) Login(email string, password string) (token string, err error) {
	id, err := a.UserService.GetUserId(email)
	if err != nil {
		return "", err
	}
	if err := a.authenticate(id, password); err != nil {
		return "", errors.New("invalid Password")
	}
	soon := time.Now().Add(a.SessionDuration)
	claims := &jwt2.SessionClaims{
		UserId:         id,
		Username:       email,
		ExpirationTime: soon.Unix(),
		FeatureFlags:   0,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(a.Secret))
}

func (a *JwtAuthenticator) getSecret(token *jwt.Token) (interface{}, error) {
	return []byte(a.Secret), nil
}

func (a *JwtAuthenticator) AuthenticateSession(tokenString string) (username *string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt2.SessionClaims{}, a.getSecret)
	if err != nil || !token.Valid {
		return nil
	}
	name := token.Claims.(*jwt2.SessionClaims).Username
	return &name
}

func (a *JwtAuthenticator) authenticate(id int64, password string) (error) {
	hash, err := a.HashStore.GetHash(id)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
