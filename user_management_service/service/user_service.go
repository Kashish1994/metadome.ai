package service

import (
	"errors"
	"github.com/eduhub/helper"
	"github.com/eduhub/models"
	"github.com/eduhub/requests"
	"github.com/eduhub/requests/validators"
	"gorm.io/gorm"
	"sync"
)

type UserService interface {
	Login(userRequest *requests.LoginRequest) (*requests.LoginResponse, error)
	GetUser(email string) (*models.User, error)
	RegisterUser(request *requests.SignUpRequest) (*requests.BasicResponse, error)
	UpdateUser(request *requests.UpdateRequest) (*requests.BasicResponse, error)
	DeleteUser(username string) (*requests.BasicResponse, error)
}
type UserServiceImpl struct {
	Db *gorm.DB
}

var userService *UserServiceImpl
var once sync.Once

func (u *UserServiceImpl) DeleteUser(username string) (*requests.BasicResponse, error) {
	user, err := u.GetUser(username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	u.Db.Delete(user)
	return &requests.BasicResponse{
		Success: true,
		Message: "User successfully deleted",
	}, nil
}
func (u *UserServiceImpl) UpdateUser(request *requests.UpdateRequest) (*requests.BasicResponse, error) {
	user, err := u.GetUser(request.UserName)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	err = validators.ValidateRequest(request)
	if err != nil {
		return nil, err
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Phone = request.Phone
	user.Address = request.Address
	u.Db.Save(user)

	return &requests.BasicResponse{
		Success: true,
		Message: "User successfully updated",
	}, nil
}

func (u *UserServiceImpl) RegisterUser(request *requests.SignUpRequest) (*requests.BasicResponse, error) {
	user, err := u.GetUser(request.UserName)
	if err != nil {
		return nil, err
	}
	if user.ID != 0 {
		return nil, errors.New("user already exist")
	}
	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	err = u.Db.Create(&models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Phone:     request.Phone,
		Address:   request.Address,
		Password:  password,
	}).Error

	if err != nil {
		return nil, err
	}
	return &requests.BasicResponse{
		Success: true,
		Message: "User successfully registered",
	}, nil
}

func (u *UserServiceImpl) Login(userRequest *requests.LoginRequest) (*requests.LoginResponse, error) {
	user, err := u.GetUser(userRequest.Username)
	if err != nil {
		return nil, err
	}
	errr := helper.VerifyPassword(user.Password, userRequest.Password)
	if errr != nil {
		return &requests.LoginResponse{
			Message: "Invalid password",
			Success: false,
		}, nil
	}
	token, _ := helper.GenerateToken(user.UserName)
	return &requests.LoginResponse{
		Token:   token,
		Success: true,
		Message: "success",
	}, nil
}

func (u *UserServiceImpl) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	err := u.Db.Last(user, "user_name = ? ", username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func InitUserService(Db *gorm.DB) UserService {
	once.Do(func() {
		userService = &UserServiceImpl{Db: Db}
	})
	return userService
}

func GetUserServiceExistingInstance() UserService {
	return userService
}
