package repository

import (
	"context"
	"go-migratre-sample/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, data *model.User) (uint, error)
	Update(ctx context.Context, data *model.User) error
	Delete(ctx context.Context, userID uint) error
	GetList(ctx context.Context, limit int, offset int) ([]model.User, error)
	GetOne(ctx context.Context, userID uint) (model.User, error)
}
