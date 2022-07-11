package Models

import (
	"errors"
)

var (
	ErrInvalidIDD     = errors.New("invalid id")
	ErrMissingFName   = errors.New("student first name is invalid")
	ErrMissingDOB     = errors.New("student DOB is invalid")
	ErrMissingAddress = errors.New("student Address is invalid")
)

type Student struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

func (student *Student) ValidateStudent() error {
	switch {
	case student.ID < 0:
		return ErrInvalidIDD
	case student.Address == "":
		return ErrMissingAddress
	case student.FirstName == "":
		return ErrMissingFName
	case student.DOB == "":
		return ErrMissingDOB
	default:
		return nil
	}
}

type StudentInfo struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	DOB       string         `json:"dob"`
	Address   string         `json:"address"`
	Marks     []SubjectMarks `json:"marks"`
}

func (studentInfo *StudentInfo) ValidateStudentInfo() error {
	switch {
	case studentInfo.ID < 0:
		return ErrInvalidIDD
	case studentInfo.Address == "":
		return ErrMissingAddress
	case studentInfo.FirstName == "":
		return ErrMissingFName
	case studentInfo.DOB == "":
		return ErrMissingDOB
	default:
		return nil
	}
}

type UpdatedStudent struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

func (updatedStudent *UpdatedStudent) ValidateUpdatedStudent() error {
	switch {
	case updatedStudent.Address == "":
		return ErrMissingAddress
	case updatedStudent.FirstName == "":
		return ErrMissingFName
	case updatedStudent.DOB == "":
		return ErrMissingDOB
	default:
		return nil
	}
}
