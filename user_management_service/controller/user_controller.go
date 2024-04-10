package controller

import (
	"errors"
	"fmt"
	"github.com/eduhub/requests"
	"github.com/eduhub/requests/validators"
	"github.com/eduhub/service"
	"github.com/eduhub/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserController struct {
	UserService service.UserService
}

func (uc UserController) FetchUser(ctx *gin.Context) {
	value := ctx.GetHeader("Authorization")
	token := strings.Split(value, " ")[1]
	resp, err := uc.UserService.FetchUser(token)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uc UserController) RegisterUser(ctx *gin.Context) {
	signUpRequest := &requests.SignUpRequest{}
	err := ctx.ShouldBindJSON(signUpRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: "Bad Request"})
		return
	}
	err = validators.ValidateRequest(signUpRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	resp, err := uc.UserService.RegisterUser(signUpRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (uc UserController) UpdateUser(ctx *gin.Context) {
	updateRequest := &requests.UpdateRequest{}
	err := ctx.ShouldBindJSON(updateRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: "Bad Request"})
		return
	}
	err = validators.ValidateRequest(updateRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	resp, err := uc.UserService.UpdateUser(updateRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (uc UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: "UN is mandatory field"})
	}
	resp, err := uc.UserService.DeleteUser(username)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.BasicResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (uc UserController) Login(ctx *gin.Context) {
	loginRequest := &requests.LoginRequest{}
	err := ctx.ShouldBindJSON(loginRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.LoginResponse{Success: false})
		return
	}
	fmt.Printf("%v", loginRequest)
	err = uc.ValidateLogin(loginRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.LoginResponse{Success: false})
		return
	}
	resp, err := uc.UserService.Login(loginRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, &requests.LoginResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (uc UserController) ValidateLogin(request *requests.LoginRequest) error {
	if request.Username == util.EmptyString || request.Password == util.EmptyString {
		return errors.New("invalid request")
	}
	return nil
}
