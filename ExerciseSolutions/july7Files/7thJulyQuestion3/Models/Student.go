// Models/Student.go
package Models

import "C"
import (
	"7thJulyQuestion3/Config"
	_ "github.com/go-sql-driver/mysql"
)

// CreateStudent ... Insert New data
func CreateStudent(Student *Student) (err error) {
	if err = Config.DB.Create(Student).Error; err != nil {
		return err
	}
	return nil
}

// GetStudentByID ... Fetch only one Student by Id
func GetStudentByID(Student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(Student).Error; err != nil {
		return err
	}
	return nil
}

// GetAllStudents Fetch all Student data
func GetAllStudents(Student *[]Student) (err error) {
	if err = Config.DB.Find(Student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent ... Update Student
func UpdateStudent(Student *Student) (err error) {

	if err = Config.DB.Save(Student).Error; err != nil {
		return err
	}
	return nil
}

// DeleteStudent ... Delete Student
func DeleteStudent(Student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(Student).Error; err != nil {
		return err
	}
	if err = Config.DB.Where("id = ?", id).Delete(Student).Error; err != nil {
		return err
	}
	return nil
}
