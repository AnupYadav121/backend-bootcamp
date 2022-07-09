package Models

type SubjectMarks struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
	Grade   string `json:"grade"`
}

type Student struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DOB            string `json:"dob"`
	Address        string `json:"address"`
	SubjectMarksID uint   `json:"subjectMarksID" gorm:"foreign_key"`
}

type InputStudent struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DOB            string `json:"dob"`
	Address        string `json:"address"`
	SubjectMarksID uint   `json:"subjectMarksID" gorm:"foreign_key"`
}

type UpdatedStudent struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DOB            string `json:"dob"`
	Address        string `json:"address"`
	SubjectMarksID uint   `json:"subjectMarksID" gorm:"foreign_key"`
}

type UpdatedSubjectMarks struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
	Grade   string `json:"grade"`
}
