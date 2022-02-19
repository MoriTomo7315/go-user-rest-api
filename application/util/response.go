package util

import (
	"log"

	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	logger "github.com/MoriTomo7315/go-user-rest-api/infrastructure/gcplogger"
)

func GetResponse(status int32, message string, count int64, users []*model.User) *model.ResponseModel {
	log.Printf(logger.InfoLogEntry("start creating response"))

	res := &model.ResponseModel{
		Status:    status,
		Message:   message,
		UserCount: count,
		Users:     users,
	}

	log.Printf(logger.InfoLogEntry("end creating response"))

	return res
}
