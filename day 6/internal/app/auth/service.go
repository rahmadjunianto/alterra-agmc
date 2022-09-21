package auth

import (
	"context"
	"day6/internal/dto"
	"day6/internal/factory"
	"day6/internal/pkg/util"
	"day6/internal/repository"
	"day6/pkg/constant"
	pkgutil "day6/pkg/util"
	res "day6/pkg/util/response"
	"errors"
)

type service struct {
	UsersRepository repository.Users
}

func NewService(f *factory.Factory) Service {
	return &service{
		UsersRepository: f.UsersRepository,
	}
}

type Service interface {
	Login(ctx context.Context, payload *dto.LoginUsersRequestBody) (*dto.UsersWithJWTResponse, error)
	Register(ctx context.Context, payload *dto.CreateUsersRequestBody) (*dto.UsersWithJWTResponse, error)
}

func (s *service) Login(ctx context.Context, payload *dto.LoginUsersRequestBody) (*dto.UsersWithJWTResponse, error) {
	var result *dto.UsersWithJWTResponse

	data, err := s.UsersRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}
	claims := util.CreateJWTClaims(data.Email, data.Name, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.UsersWithJWTResponse{
		UsersResponse: dto.UsersResponse{
			ID:    data.ID,
			Name:  data.Name,
			Email: data.Email,
		},
		JWT: token,
	}

	return result, nil
}
func (s *service) Register(ctx context.Context, payload *dto.CreateUsersRequestBody) (*dto.UsersWithJWTResponse, error) {
	var result *dto.UsersWithJWTResponse
	isExist, err := s.UsersRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("employee already exists"))
	}

	hashedPassword, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword
	data, err := s.UsersRepository.Save(ctx, payload)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	claims := util.CreateJWTClaims(data.Email, data.Name, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}
	result = &dto.UsersWithJWTResponse{
		UsersResponse: dto.UsersResponse{
			ID:    data.ID,
			Name:  data.Name,
			Email: data.Email,
		},
		JWT: token,
	}

	return result, nil
}
