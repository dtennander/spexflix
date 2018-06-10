package server

import (
	"testing"
	"net/http/httptest"
	"github.com/golang/mock/gomock"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/DiTo04/spexflix/common/codecs"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluIiwiZXhwIjoxNTYwMTExODg5LCJmZiI6MH0.lS_vLAKgohK2vr7zqnAeap67a5vJbzML_DYToXySh2Y"

func initTarget(t *testing.T) (http.Handler, *MockStorageService) {
	ctrl := gomock.NewController(t)
	storageService := NewMockStorageService(ctrl)
	controller := &controller{
		storageService: storageService,
		jwtSecret: "Secret",
	}
	defer ctrl.Finish()
	return controller.CreateRoutes(), storageService
}

//go:generate mockgen -package server -destination controller_mocks.go -source controller.go
func TestListYears(t *testing.T) {
	// Given
	const year = "2015"
	target, storageService := initTarget(t)
	req := httptest.NewRequest("GET", "/movies", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	rsp := httptest.NewRecorder()
	storageService.EXPECT().GetYears(gomock.Any()).Return([]Year{{Year: year}}, nil)

	// When
	target.ServeHTTP(rsp, req)

	// Then
	var years []Year
	codecs.JSON.Decode(rsp.Body, &years)
	assert.Equal(t, 1, len(years))
	assert.Equal(t, year, years[0].Year)
}

func TestListContent(t *testing.T) {
	// Given
	const year = "2015"
	target, storageService := initTarget(t)
	req := httptest.NewRequest("GET", "/movies/" + year, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	rsp := httptest.NewRecorder()
	expextedResult := Movie{
		Name: "trailer",
		Description: "asd",
		Uri: "uriToCloudStorage",
	}
	storageService.EXPECT().GetContent(gomock.Any(), year).Return([]Movie{expextedResult}, nil)

	// When
	target.ServeHTTP(rsp, req)

	// Then
	var years []Movie
	codecs.JSON.Decode(rsp.Body, &years)
	assert.Equal(t, 1, len(years))
	assert.Equal(t, expextedResult, years[0])
}

func TestListYearsUnAuthorized(t *testing.T) {
	// Given
	const year = "2015"
	target, _ := initTarget(t)
	req := httptest.NewRequest("GET", "/movies", nil)
	rsp := httptest.NewRecorder()

	// When
	target.ServeHTTP(rsp, req)

	// Then
	assert.Equal(t, http.StatusUnauthorized, rsp.Code)
}

func TestListContentUnAuthorized(t *testing.T) {
	// Given
	const year = "2015"
	target, _ := initTarget(t)
	req := httptest.NewRequest("GET", "/movies/" + year, nil)
	rsp := httptest.NewRecorder()

	// When
	target.ServeHTTP(rsp, req)

	// Then
	assert.Equal(t, http.StatusUnauthorized, rsp.Code)
}
