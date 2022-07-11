package Models

type Student struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

type StudentInfo struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	DOB       string         `json:"dob"`
	Address   string         `json:"address"`
	Marks     []SubjectMarks `json:"marks"`
}

type UpdatedStudent struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

type SubjectMarks struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	StudentId uint   `json:"student_id" gorm:"foreign_key" gorm:"index"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	Grade     string `json:"grade"`
}

type UpdatedSubjectMarks struct {
	StudentId uint   `json:"student_id" gorm:"foreign_key" gorm:"index"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	Grade     string `json:"grade"`
}
