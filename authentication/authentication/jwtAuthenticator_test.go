package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	mocks "github.com/DiTo04/spexflix/authentication/authentication/mock_authentication"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source jwtAuthenticator.go -destination mock_authentication/jwtAuthenticator.go
func TestLogin(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hashStore := mocks.NewMockHashStore(ctrl)
	userService := mocks.NewMockUserService(ctrl)
	target := &JwtAuthenticator{
		SessionDuration: time.Minute,
		Secret:          "Secret",
		HashStore:       hashStore,
		UserService:     userService,
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("kakakaka"), bcrypt.DefaultCost)
	userService.EXPECT().GetUserId("admin").Return(int64(2), nil)
	hashStore.EXPECT().GetHash(int64(2)).Return([]byte(hash), nil)
	// When
	token, err := target.Login("admin", "kakakaka")
	// Then
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthenticate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hashStore := mocks.NewMockHashStore(ctrl)
	userService := mocks.NewMockUserService(ctrl)
	target := &JwtAuthenticator{
		SessionDuration: time.Minute,
		Secret:          "Secret",
		HashStore:       hashStore,
		UserService:     userService,
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("lol"), bcrypt.DefaultCost)
	userService.EXPECT().GetUserId("admin").Return(int64(2), nil)
	hashStore.EXPECT().GetHash(int64(2)).Return([]byte(hash), nil)
	token, _ := target.Login("admin", "lol")
	// When
	resultName := target.AuthenticateSession(token)
	// Then
	assert.Equal(t, "admin", *resultName)
}
