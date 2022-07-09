package Utils

import (
	"7thJulyQuestion2/Config"
	"7thJulyQuestion2/Models"
)

func IsPresent(id string, Student *Models.Student) error {
	return Config.DB.Where("id = ?", id).First(Student).Error
}

func DoCreate(Student *Models.Student) {
	Config.DB.Create(Student)
}

func DoFind(Student *[]Models.Student) {
	Config.DB.Find(Student)
}

func DoUpdate(Student *Models.Student, newStudent Models.UpdatedStudent) {
	Config.DB.Model(Student).Updates(newStudent)
}

func DoDelete(Student *Models.Student) error {
	return Config.DB.Delete(Student).Error
}
