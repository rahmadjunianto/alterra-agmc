package database

import (
	"day2/config"
	"day2/models"
	"fmt"
)

func GetUsers() (*[]models.Users, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return &users, nil
}
func GetUserById(id *int) (*models.Users, error) {
	var user models.Users
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return &user, nil
}
func CreateUsers(user *models.Users) (*models.Users, error) {
	if e := config.DB.
		Create(&user).
		Error; e != nil {
		return nil, e
	}
	return user, nil
}
func UpdateUserById(id *int, data *models.Users) (*models.Users, error) {
	var user models.Users
	e := config.DB.First(&user, id).Updates(&data)
	if e.RowsAffected < 1 {
		return nil, fmt.Errorf("row with id=%v  not found", *id)
	} else if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}
func DeleteUserById(id *int) error {
	var user models.Users
	e := config.DB.First(&user, id).Delete(&user)
	if e.RowsAffected < 1 {
		return fmt.Errorf("row with id=%v  not found", *id)
	} else if e.Error != nil {
		return e.Error
	}
	return nil
}
