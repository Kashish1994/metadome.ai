package main

import (
	"fmt"
	"github.com/eduhub/configs"
	"github.com/eduhub/controller"
	"github.com/eduhub/helper"

	"github.com/eduhub/middleware"

	"github.com/eduhub/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	pass, _ := helper.HashPassword("password")
	fmt.Println(pass)

	db, err := configs.SetupDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middleware.ValidateToken())

	baseRouter := router.Group("/metadome-api")
	userRouter := baseRouter.Group("/user")

	// Init Service
	userService := service.InitUserService(db)

	// Register Routes
	userRouter.POST("/", controller.UserController{UserService: userService}.RegisterUser)
	userRouter.PUT("/", controller.UserController{UserService: userService}.UpdateUser)
	userRouter.DELETE("/", controller.UserController{UserService: userService}.DeleteUser)
	userRouter.POST("/login", controller.UserController{UserService: userService}.Login)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
