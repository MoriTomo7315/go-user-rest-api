package repository

import "github.com/MoriTomo7315/go-user-rest-api/domain/model"

type ResponseRepository interface {
	GetResponse(status int32, message string, count int64, users []*model.User) *model.ResponseModel
}
