package util

import (
	"log"
	
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
)


func GetResponse(status int32, message string, count int64, users []*model.User) *model.ResponseModel {
	log.Printf("INFO start creating response")

	res := &model.ResponseModel{
		Status: status,
		Message: message,
		UserCount: count,
		Users: users,
	}

	log.Printf("INFO end creating response")

	return res
}