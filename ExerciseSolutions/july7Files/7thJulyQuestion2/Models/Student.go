package Models

type Subject struct {
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
	Grade   string `json:"grade"`
}

type Student struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	DOB          string    `json:"dob"`
	Address      string    `json:"address"`
	SubjectMarks []Subject `json:"subjectMarks"`
}

type InputStudent struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	DOB          string    `json:"dob"`
	Address      string    `json:"address"`
	SubjectMarks []Subject `json:"subjectMarks"`
}

type UpdatedStudent struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	DOB          string    `json:"dob"`
	Address      string    `json:"address"`
	SubjectMarks []Subject `json:"subjectMarks"`
}
