package repository

import "github.com/MoriTomo7315/go-user-rest-api/domain/model"

type ErrorRepository interface {
	GetErrorResponse(string) *model.ErrorResponse
}
