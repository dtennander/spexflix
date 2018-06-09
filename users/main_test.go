package main

import (
	"bytes"
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluIiwiZXhwIjoxNTYwMTExODg5LCJmZiI6MH0.lS_vLAKgohK2vr7zqnAeap67a5vJbzML_DYToXySh2Y"

var userIds int64 = 1

const email = "tester@karspexet.se"

type mockUsers map[int64]User

func (u mockUsers) queryUsers(emailAdress string) ([]*User, error) {
	if email != emailAdress {
		return make([]*User, 0), errors.New("email: " + emailAdress)
	}
	users := make([]*User, 1)
	user := u[2]
	users[0] = &user
	return users, nil
}

func (u mockUsers) postUser(user *User) (int64, error) {
	userIds = userIds + 1
	return userIds, nil
}

func (u mockUsers) getUser(userId int64) (*User, error) {
	user := u[userId]
	return &user, nil
}

var users = mockUsers{
	2: {
		Id:        2,
		Name:      "testUser",
		Email:     email,
		SpexYears: 10,
	},
}

func TestGetUser(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/users/2", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	requestRecorder := httptest.NewRecorder()

	// when
	target.ServeHTTP(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.Equal(t, "{\"id\":2,\"name\":\"testUser\",\"email\":\"tester@karspexet.se\",\"spex_years\":10}\n", requestRecorder.Body.String())
}

func TestGetUserWithoutAccess(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/users/2", nil)
	req.Header.Add("Authorization", "Bearer "+token+"1")
	requestRecorder := httptest.NewRecorder()

	// when
	target.ServeHTTP(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusUnauthorized, requestRecorder.Code)
}

func TestHealthz(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/healthz", nil)
	requestRecorder := httptest.NewRecorder()

	// when
	target.ServeHTTP(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
}

func TestPostUser(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	body := &bytes.Buffer{}
	codecs.JSON.Encode(body, User{
		Name:      "tester",
		Email:     "tester@spexflix.se",
		SpexYears: 10,
	})
	req := httptest.NewRequest("POST", "/users", body)
	req.Header.Add("Authorization", "Bearer "+token)
	requestRecorder := httptest.NewRecorder()

	// when
	target.ServeHTTP(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	resultingUser := &User{}
	codecs.JSON.Decode(requestRecorder.Body, resultingUser)
	assert.Equal(t, userIds, resultingUser.Id)
}

func TestPostUserUnAuthorized(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	body := &bytes.Buffer{}
	codecs.JSON.Encode(body, User{
		Name:      "tester",
		Email:     "tester@spexflix.se",
		SpexYears: 10,
	})
	req := httptest.NewRequest("POST", "/users", body)
	req.Header.Add("Authorization", "Bearer "+token+"1")
	requestRecorder := httptest.NewRecorder()

	// when
	target.ServeHTTP(requestRecorder, req)

	// Then
	assert.Equal(t, http.StatusUnauthorized, requestRecorder.Code)
}

func TestGetUserFromEmail(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/users?email="+email, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	requestRec := httptest.NewRecorder()

	// When
	target.ServeHTTP(requestRec, req)

	// Then
	var rsp []User
	codecs.JSON.Decode(requestRec.Body, &rsp)
	println(requestRec.Body.String())
	assert.Equal(t, http.StatusOK, requestRec.Code)
	assert.Equal(t, 1, len(rsp))
	assert.NotNil(t, rsp[0])
	assert.Equal(t, email, rsp[0].Email)
}

func TestGetUserFromEmailIsProtected(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/users?email="+email, nil)
	req.Header.Add("Authorization", "Bearer "+token+"1")
	requestRec := httptest.NewRecorder()

	// When
	target.ServeHTTP(requestRec, req)

	// Then
	assert.Equal(t, http.StatusUnauthorized, requestRec.Code)
}

func TestGetIdFromEmail(t *testing.T) {
	controller := &controller{
		jwtSecret: "Secret",
		users:     users,
	}
	target := controller.getRouter()
	req := httptest.NewRequest("GET", "/id?email="+email, nil)
	requestRec := httptest.NewRecorder()

	// When
	target.ServeHTTP(requestRec, req)

	// Then
	assert.Equal(t, http.StatusOK, requestRec.Code)
	assert.Equal(t, "2\n", requestRec.Body.String())
}

func TestUserValidationn(t *testing.T) {
	// Given
	user := &User{
		Name:      "Tester",
		Email:     "tester@spexflix.se",
		SpexYears: 10,
	}
	// When
	err := validateUser(user)
	// Then
	assert.Nil(t, err)
}

func TestUserValidation2(t *testing.T) {
	// Given
	user := &User{
		Email:     "tester@spexflix.se",
		SpexYears: 10,
	}
	// When
	err := validateUser(user)
	// Then
	assert.NotNil(t, err)
}

func TestUserValidation3(t *testing.T) {
	// Given
	user := &User{
		Name:      "Tester",
		SpexYears: 10,
	}
	// When
	err := validateUser(user)
	// Then
	assert.NotNil(t, err)
}

func TestUserValidation4(t *testing.T) {
	// Given
	user := &User{
		Name:  "Tester",
		Email: "tester@spexflix.se",
	}
	// When
	err := validateUser(user)
	// Then
	assert.NotNil(t, err)
}

func TestUserValidation5(t *testing.T) {
	// Given
	user := &User{
		Id:        1,
		Name:      "Tester",
		Email:     "tester@spexflix.se",
		SpexYears: 10,
	}
	// When
	err := validateUser(user)
	// Then
	assert.NotNil(t, err)
}
