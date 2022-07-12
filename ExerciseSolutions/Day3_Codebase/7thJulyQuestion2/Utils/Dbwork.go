package Utils

import (
	"7thJulyQuestion2/Config"
	"7thJulyQuestion2/Models"
)

type InterfaceDB interface {
	IsPresent(id string, Student *Models.Student) (*Models.Student, error)
	DoCreate(Student *Models.Student)
	DoFind(Student *[]Models.Student)
	DoUpdate(Student *Models.Student, newStudent Models.UpdatedStudent)
	DoDelete(Student *Models.Student) error

	IsPresentMark(id string, SubjectMarks *Models.SubjectMarks) error
	DoCreateMark(SubjectMarks *Models.SubjectMarks)
	DoFindMarks(SubjectMarks *[]Models.SubjectMarks)
	DoUpdateMark(SubjectMarks *Models.SubjectMarks, newSubjectMarks Models.UpdatedSubjectMarks)
	DoDeleteMark(SubjectMarks *Models.SubjectMarks) error
	MyMarks(id string, SubjectMarks *[]Models.SubjectMarks)
	IsMyMark(id string, SubjectMarks *[]Models.SubjectMarks) error
}

type DB struct {
}

func GetDB() InterfaceDB {
	return &DB{}
}

func (db *DB) IsPresent(id string, Student *Models.Student) (*Models.Student, error) {
	err := Config.DB.Where("id = ?", id).First(Student)
	if err != nil {
		return nil, err.Error
	}
	return Student, nil
}

func (db *DB) DoCreate(Student *Models.Student) {
	Config.DB.Create(Student)
}

func (db *DB) DoFind(Student *[]Models.Student) {
	Config.DB.Find(Student)
}

func (db *DB) DoUpdate(Student *Models.Student, newStudent Models.UpdatedStudent) {
	Config.DB.Model(Student).Updates(newStudent)
}

func (db *DB) DoDelete(Student *Models.Student) error {
	return Config.DB.Delete(Student).Error
}

func (db *DB) IsPresentMark(id string, SubjectMarks *Models.SubjectMarks) error {
	return Config.DB.Where("id = ?", id).First(SubjectMarks).Error
}

func (db *DB) DoCreateMark(SubjectMarks *Models.SubjectMarks) {
	Config.DB.Create(SubjectMarks)
}

func (db *DB) DoFindMarks(SubjectMarks *[]Models.SubjectMarks) {
	Config.DB.Find(SubjectMarks)
}

func (db *DB) DoUpdateMark(SubjectMarks *Models.SubjectMarks, newSubjectMarks Models.UpdatedSubjectMarks) {
	Config.DB.Model(SubjectMarks).Updates(newSubjectMarks)
}

func (db *DB) DoDeleteMark(SubjectMarks *Models.SubjectMarks) error {
	return Config.DB.Delete(SubjectMarks).Error
}

func (db *DB) MyMarks(id string, SubjectMarks *[]Models.SubjectMarks) {
	Config.DB.Where("student_id = ?", id).Find(SubjectMarks)
}

func (db *DB) IsMyMark(id string, SubjectMarks *[]Models.SubjectMarks) error {
	return Config.DB.Where("student_id = ?", id).First(SubjectMarks).Error
}
