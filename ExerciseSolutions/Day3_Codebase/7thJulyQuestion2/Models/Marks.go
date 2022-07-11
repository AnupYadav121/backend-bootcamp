package Models

import (
	"errors"
)

var (
	ErrInvalidID   = errors.New("invalid ID")
	ErrStudentID   = errors.New("student id is invalid")
	ErrSubjectName = errors.New("subject name is missing")
	ErrMarks       = errors.New("marks provided is invalid")
	ErrGrade       = errors.New("grade provided is invalid")
)

type SubjectMarks struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	StudentId uint   `json:"student_id" gorm:"foreign_key" gorm:"index"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	Grade     string `json:"grade"`
}

func (marks *SubjectMarks) ValidateMarks() error {
	switch {
	case marks.ID < 0:
		return ErrInvalidID
	case marks.Marks > 100 || marks.Marks < 0:
		return ErrMarks
	case len(marks.Grade) > 1 || int(marks.Grade[0]) < 65 || int(marks.Grade[0]) > 90:
		return ErrGrade
	case marks.StudentId <= 0:
		return ErrStudentID
	case marks.Subject == "":
		return ErrSubjectName
	default:
		return nil
	}
}

type UpdatedSubjectMarks struct {
	StudentId uint   `json:"student_id" gorm:"foreign_key" gorm:"index"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	Grade     string `json:"grade"`
}

func (uMarks *UpdatedSubjectMarks) ValidateUpdatedMarks() error {
	switch {
	case uMarks.Marks > 100 || uMarks.Marks < 0:
		return ErrMarks
	case len(uMarks.Grade) > 1 || int(uMarks.Grade[0]) < 65 || int(uMarks.Grade[0]) > 90:
		return ErrGrade
	case uMarks.StudentId <= 0:
		return ErrStudentID
	case uMarks.Subject == "":
		return ErrSubjectName
	default:
		return nil
	}
}
