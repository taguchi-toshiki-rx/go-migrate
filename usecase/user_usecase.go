package usecase

import (
	"context"
	"go-migratre-sample/domain/model"
	"go-migratre-sample/domain/repository"

	"github.com/samber/lo"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, reqData ReqCreateUpdateUser) error
	UpdateUser(ctx context.Context, ID uint, reqData ReqCreateUpdateUser) error
	DeleteUser(ctx context.Context, ID uint) error
	GetUserList(ctx context.Context, limit int, offset int) (ResGetUserList, error)
	GetUserOne(ctx context.Context, ID uint) (ResGetUser, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

type ReqCreateUpdateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ReqGetUserList struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}
type ResGetUserList struct {
	Users []ResGetUser `json:"users"`
}
type ResGetUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *userUsecase) CreateUser(ctx context.Context, reqData ReqCreateUpdateUser) error {

	user, err := model.NewUser(reqData.Name, reqData.Age)
	if err != nil {
		return err
	}

	_, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, ID uint, reqData ReqCreateUpdateUser) error {

	user := model.User{
		Name: reqData.Name,
		Age:  reqData.Age,
	}

	err := u.userRepo.Update(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, ID uint) error {

	err := u.userRepo.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetUserList(ctx context.Context, limit int, offset int) (ResGetUserList, error) {
	var resUsers ResGetUserList

	// limit default is 100
	l := limit
	if l == 0 {
		l = 100
	}

	users, err := u.userRepo.GetList(ctx, l, offset)
	if err != nil {
		return resUsers, err
	}

	resUsers = ResGetUserList{
		Users: lo.Map(users, func(u model.User, _ int) ResGetUser {
			return ResGetUser{
				ID:   u.ID,
				Name: u.Name,
				Age:  u.Age,
			}
		}),
	}

	return resUsers, nil
}

func (u *userUsecase) GetUserOne(ctx context.Context, ID uint) (ResGetUser, error) {
	var resUser ResGetUser

	user, err := u.userRepo.GetOne(ctx, ID)
	if err != nil {
		return resUser, err
	}

	resUser = ResGetUser{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}

	return resUser, nil
}
