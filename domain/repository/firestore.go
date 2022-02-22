package repository

import (
	"context"

	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
)

type FirestoreRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}
