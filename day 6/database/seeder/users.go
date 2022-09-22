package seeder

import (
	"day6/internal/model"
	"log"
	"time"

	"gorm.io/gorm"
)

func usersSeeder(db *gorm.DB) {
	now := time.Now()
	var users = []model.Users{
		{
			Common:   model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
			Name:     "user1",
			Email:    "user1@gmail.com",
			Password: "$2a$10$bKAc90iPix.ud43bSs6Yi.z8wfH.73K5PXWYLx6DF7DtHeWRUdZXe", //12345678
		},
		{
			Common:   model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
			Name:     "user2",
			Email:    "user2@gmail.com",
			Password: "$2a$10$bKAc90iPix.ud43bSs6Yi.z8wfH.73K5PXWYLx6DF7DtHeWRUdZXe", //12345678
		},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}
