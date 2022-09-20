package users

import (
	"context"
	"day6/internal/dto"
	"day6/internal/factory"
	"day6/internal/repository"
	"day6/pkg/constant"
	pkgdto "day6/pkg/dto"
	res "day6/pkg/util/response"
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
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UsersResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UsersResponse, error)
	Store(ctx context.Context, payload *dto.CreateUsersRequestBody) (*dto.UsersResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateUsersRequestBody) (*dto.UsersResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UsersWithCUDResponse, error)
	Login(ctx context.Context, payload *dto.LoginUsersRequestBody) (*dto.LoginUsersResponse, error)
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UsersResponse], error) {
	//TODO implement me
	users, info, err := s.UsersRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var data []dto.UsersResponse

	for _, user := range users {
		data = append(data, dto.UsersResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	result := new(pkgdto.SearchGetResponse[dto.UsersResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UsersResponse, error) {
	//TODO implement me
	user, err := s.UsersRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.UsersResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UsersResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UsersResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateUsersRequestBody) (*dto.UsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateUsersRequestBody) (*dto.UsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UsersWithCUDResponse, error) {
	//TODO implement me
	panic("implement me")
}
func (s *service) Login(ctx context.Context, payload *dto.LoginUsersRequestBody) (*dto.LoginUsersResponse, error) {
	user, err := s.UsersRepository.Login(ctx, payload)

	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.LoginUsersResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.LoginUsersResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result := &dto.LoginUsersResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
	return result, nil
}
