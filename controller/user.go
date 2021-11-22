package controller

import (
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/application"
)

type UserController interface {
	HandlerHttpRequest(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userApplication application.UserApplication
}

func NewUserController(ua application.UserApplication) UserController {
	return &userController{
		userApplication: ua,
	}
}

func (uc *userController) HandlerHttpRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO [/api/users] START ===========")
	switch r.Method {
	case "GET":
		uc.userApplication.GetUserById(w, r)
	default:
		w.WriteHeader(405)
	}
	log.Printf("INFO [/api/users] END ===========")
}
