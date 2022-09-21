package repository

import (
	"context"
	"day6/internal/dto"
	"day6/internal/model"
	pkgdto "day6/pkg/dto"
	"day6/pkg/util"
	"strings"

	"gorm.io/gorm"
)

type Users interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Users, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Users, error)
	FindByEmail(ctx context.Context, email *string) (*model.Users, error)
	ExistByEmail(ctx context.Context, email *string) (bool, error)
	Save(ctx context.Context, user *dto.CreateUsersRequestBody) (model.Users, error)
	Edit(ctx context.Context, oldUser *model.Users, updateData *dto.UpdateUsersRequestBody) (*model.Users, error)
	Destroy(ctx context.Context, user *model.Users) (*model.Users, error)
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
func (u *users) FindByEmail(ctx context.Context, email *string) (*model.Users, error) {
	var data model.Users
	err := u.DB.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (u *users) ExistByEmail(ctx context.Context, email *string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := u.DB.WithContext(ctx).Model(&model.Users{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (u *users) Save(ctx context.Context, user *dto.CreateUsersRequestBody) (model.Users, error) {
	//TODO implement me
	newUser := model.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.DB.WithContext(ctx).Save(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (u *users) Edit(ctx context.Context, oldUser *model.Users, updateData *dto.UpdateUsersRequestBody) (*model.Users, error) {

	if updateData.Name != nil {
		oldUser.Name = *updateData.Name
	}
	if updateData.Email != nil {
		oldUser.Email = *updateData.Email
	}
	if updateData.Password != nil {
		hashedPassword, err := util.HashPassword(*updateData.Password)
		if err != nil {
			return nil, err
		}
		oldUser.Password = hashedPassword
	}

	if err := u.DB.
		WithContext(ctx).
		Save(oldUser).
		Find(oldUser).
		Error; err != nil {
		return nil, err
	}

	return oldUser, nil
}

func (u *users) Destroy(ctx context.Context, user *model.Users) (*model.Users, error) {

	if err := u.DB.WithContext(ctx).Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
