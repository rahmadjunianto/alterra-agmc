package seeder

import (
	"day4/config"
	"day4/models"
	"gorm.io/gorm"
	"log"
)

type seed struct {
	DB *gorm.DB
}

func NewUserSeeder() *seed {
	config.InitDB()
	return &seed{DB: config.DB}
}

func (s *seed) Seed() {
	users := []models.Users{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "user1",
			Email:    "user1@mail.com",
			Password: "user1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "user2",
			Email:    "user2@mail.com",
			Password: "user2",
		},
	}
	if err := s.DB.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}

func (s *seed) Delete() {
	s.DB.Exec("DELETE FROM users")
}
