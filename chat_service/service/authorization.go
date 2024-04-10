package service

import (
	"github.com/eduhub/external/metadome_users"
	"github.com/eduhub/helper"
	"sync"
)

type AuthService interface {
	AuthoriseAndFetchUser(token string) (*helper.UserResponse, error)
}

var authService *AuthServiceImpl
var once sync.Once

type AuthServiceImpl struct {
}

func (a AuthServiceImpl) AuthoriseAndFetchUser(token string) (*helper.UserResponse, error) {
	user, err := metadome_users.FetchUserDetails(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAuthServiceInstance() AuthService {
	once.Do(func() {
		authService = &AuthServiceImpl{}
	})
	return authService
}
