package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/MoriTomo7315/go-user-rest-api/application"
)

type UserController interface {
	HandlerHttpRequest(w http.ResponseWriter, r *http.Request)
	HandlerHttpRequestWithParameter(w http.ResponseWriter, r *http.Request)
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
	case http.MethodGet:
		/*
			全Userを取得する
		*/
		uc.userApplication.GetUsers(w, r)
	case http.MethodPost:
		/*
			Userを作成する
		*/
		uc.userApplication.CreateUser(w, r)
	default:
		/*
			GET, POST, DELETE以外のhttp methodは許可しない
		*/
		w.WriteHeader(405)
	}
	log.Printf("INFO [/api/users] END ===========")
}

func (uc *userController) HandlerHttpRequestWithParameter(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO [/api/users/] START ===========")
	userId := strings.TrimPrefix(r.URL.Path, "/api/users/")
	switch r.Method {
	case http.MethodGet:
		if len(userId) == 0 {
			/*
				/api/users/でリクエストが来た場合は/api/userと同じ処理をする（リダイレクト的役割）
			*/
			uc.userApplication.GetUsers(w, r)
		} else {
			/*
				/api/users/:userIdでリクエストが来た場合はidがuserIdのuser情報を返す
			*/
			uc.userApplication.GetUserById(w, r, userId)
		}
	case http.MethodPost:
		if len(userId) == 0 {
			/*
				/api/users/でリクエストが来た場合は/api/userと同じ処理をする（リダイレクト的役割）
			*/
			uc.userApplication.CreateUser(w, r)
		} else {
			/*
				/api/users/:userIdでリクエストが来た場合はidがuserIdのuser情報を更新する
			*/
			uc.userApplication.UpdateUser(w, r, userId)
		}
	case http.MethodDelete:
		if len(userId) == 0 {
			w.WriteHeader(400)
		} else {
			uc.userApplication.DeleteUser(w, r, userId)
		}
	default:
		/*
			GET, POST, DELETE以外のhttp methodは許可しない
		*/
		w.WriteHeader(405)
	}
	log.Printf("INFO [/api/users/] END ===========")
}
