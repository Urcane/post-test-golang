package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/service/user_service"
	"net/http"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.RegisterRequest{}
	helper.ReadFromRequestBody(request, &userRegisterRequest)

	userResponse := controller.UserService.Register(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse := controller.UserService.Login(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
