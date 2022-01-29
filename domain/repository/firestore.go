package repository

import "github.com/MoriTomo7315/go-user-rest-api/domain/model"

type FirestoreRepository interface {
	GetUsers() ([]*model.User, error)
	GetUserById(id string) (*model.User, error)
	CreateUser(user *model.User) (string, error)
	UpdateUser(user *model.User) error
	// DeleteUser(id string) (bool, error)
}
