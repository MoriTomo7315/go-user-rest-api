package infrastructure

import (
	"log"

	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
)

// response作成用のinterface
type responseFactory struct{}

func NewResponseFactory() repository.ResponseRepository {
	return &responseFactory{}
}

func (r *responseFactory) GetResponse(status int32, message string, count int64, users []*model.User) (res *model.ResponseModel) {
	log.Printf("INFO start creating response")

	res = &model.ResponseModel{
		Status: status,
		Message: message,
		UserCount: count,
		Users: users,
	}

	log.Printf("INFO end creating response")

	return res
}
