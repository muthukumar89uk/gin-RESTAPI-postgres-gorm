package models

import (
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/database"
	_ "gorm.io/driver/postgres"
)

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = database.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	if err = database.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
