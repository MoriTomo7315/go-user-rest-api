package infrastructure

import (
	"log"

	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
)

// error response用interface
type errorHandling struct{}

func NewErrorHandling() repository.ErrorRepository {
	return &errorHandling{}
}

func (e *errorHandling) GetErrorResponse(code int, errMsg string) (res *model.ErrorResponse) {
	log.Printf("INFO start creating error response")

	res = &model.ErrorResponse{
		Code:    code,
		Message: errMsg,
	}

	log.Printf("INFO end creating error response")

	return res
}
