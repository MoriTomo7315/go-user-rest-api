package util

import (
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
)

func GetResponse(status int32, message string, count int64, users []*model.User) *model.ResponseModel {

	res := &model.ResponseModel{
		Status:    status,
		Message:   message,
		UserCount: count,
		Users:     users,
	}

	return res
}
