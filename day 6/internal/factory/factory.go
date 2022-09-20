package factory

import (
	"day6/database"
	"day6/internal/repository"
)

type Factory struct {
	UsersRepository repository.Users
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUsersRepository(db),
	}
}
