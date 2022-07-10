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

func IsPresentSubject(id string, SubjectMarks *Models.SubjectMarks) error {
	return Config.DB.Where("id = ?", id).First(SubjectMarks).Error
}

func DoCreateSubject(SubjectMarks *Models.SubjectMarks) {
	Config.DB.Create(SubjectMarks)
}

func DoFindSubjects(SubjectMarks *[]Models.SubjectMarks) {
	Config.DB.Find(SubjectMarks)
}

func DoUpdateSubject(SubjectMarks *Models.SubjectMarks, newSubjectMarks Models.UpdatedSubjectMarks) {
	Config.DB.Model(SubjectMarks).Updates(newSubjectMarks)
}

func DoDeleteSubject(SubjectMarks *Models.SubjectMarks) error {
	return Config.DB.Delete(SubjectMarks).Error
}
