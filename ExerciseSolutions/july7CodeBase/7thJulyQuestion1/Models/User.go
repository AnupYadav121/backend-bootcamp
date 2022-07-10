// Models/User.go
package Models

import "C"
import (
	"7thJulyQuestion1/Config"
	_ "github.com/go-sql-driver/mysql"
)

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *User) (err error) {

	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	if err = Config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
