package repository

import (
	"context"
	"day6/internal/dto"
	"day6/internal/model"
	"day6/internal/pkg/util"
	pkgdto "day6/pkg/dto"
	"strings"

	"gorm.io/gorm"
)

type Users interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Users, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Users, error)
	ExistByName(ctx context.Context, name string) (bool, error)
	Save(ctx context.Context, rooms *dto.CreateUsersRequestBody) (model.Users, error)
	Edit(ctx context.Context, oldRooms *model.Users, updateData *dto.UpdateUsersRequestBody) (*model.Users, error)
	Destroy(ctx context.Context, rooms *model.Users) (*model.Users, error)
	Login(ctx context.Context, login *dto.LoginUsersRequestBody) (*model.Users, error)
}

type users struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *users {
	return &users{
		db,
	}
}
func (u *users) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Users, *pkgdto.PaginationInfo, error) {
	var users []model.Users
	var count int64
	query := u.DB.WithContext(ctx).Model(&users)
	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? or lower(email) LIKE ?", search, search)
	}
	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, pkgdto.CheckInfoPagination(pagination, count), err
}

func (u *users) FindByID(ctx context.Context, id uint) (model.Users, error) {
	var user model.Users
	q := u.DB.WithContext(ctx).Model(&model.Users{}).Where("id = ?", id)
	err := q.First(&user).Error
	return user, err
}

func (u *users) ExistByName(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) Save(ctx context.Context, rooms *dto.CreateUsersRequestBody) (model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) Edit(ctx context.Context, oldRooms *model.Users, updateData *dto.UpdateUsersRequestBody) (*model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) Destroy(ctx context.Context, rooms *model.Users) (*model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) Login(ctx context.Context, login *dto.LoginUsersRequestBody) (*model.Users, error) {
	//TODO implement me
	var user model.Users
	q := u.DB.WithContext(ctx).Model(&model.Users{}).Where("email = ? and password = ?", login.Email, login.Password)
	err := q.First(&user).Error
	if err != nil {
		return nil, err
	}
	token, err := util.CreateJWTToken(util.CreateJWTClaims(user.Name, user.Email, user.ID))
	if err != nil {
		return nil, err
	}
	user.Token = token
	if e := u.DB.WithContext(ctx).Save(user).Find(&user).Error; e != nil {
		return nil, e
	}
	return &user, err
}
