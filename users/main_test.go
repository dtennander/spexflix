package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluIiwiZXhwIjoxNTI4MzA3NzgzLCJmZiI6MH0.cn4ntJoyDugH6MKJRZwYMMOtq56GB53SzmG9PLRl0O4"

type mockUsers map[int64]User

func (u mockUsers) getUser(userId int64) (*User, error) {
	user := u[userId]
	return &user, nil
}

var users = mockUsers{
	2: {
		Id:        2,
		Name:      "testUser",
		Email:     "tester@karspexet.se",
		SpexYears: 10,
	},
}

func TestGetUser(t *testing.T) {
	controller := &controller{
		jwtSecret: "",
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
		jwtSecret: "",
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
		jwtSecret: "",
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
