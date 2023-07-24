package infra

import (
	"context"
	"go-migratre-sample/domain/model"
	"go-migratre-sample/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Userの新規作成
func (r *UserRepository) Create(ctx context.Context, data *model.User) (uint, error) {
	if err := r.Conn.Create(data).Error; err != nil {
		return 0, err
	}

	return data.ID, nil
}

// Update Userの更新
func (r *UserRepository) Update(ctx context.Context, data *model.User) error {
	return r.Conn.Save(data).Error
}

// Delete Userの削除
func (r *UserRepository) Delete(ctx context.Context, userID uint) error {
	return r.Conn.Delete(&model.User{}, userID).Error
}

// GetList Userの複数件取得
func (r *UserRepository) GetList(ctx context.Context, limit int, offset int) ([]model.User, error) {
	resUsers := make([]model.User, 0, limit)

	if err := r.Conn.Limit(limit).Offset(offset).Find(&resUsers).Error; err != nil {
		return resUsers, err
	}

	return resUsers, nil
}

// GetOne Userを1件取得
func (r *UserRepository) GetOne(ctx context.Context, userID uint) (model.User, error) {
	var resUser model.User

	if err := r.Conn.First(&resUser, userID).Error; err != nil {
		return resUser, err
	}

	return resUser, nil
}
