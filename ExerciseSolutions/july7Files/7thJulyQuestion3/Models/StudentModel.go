// Models/StudentModel.go
package Models

type Student struct {
	Id        uint   `json:"id"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
	DOB       int    `json:"dob"`
	Address   string `json:"address"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
}

func (s *Student) TableName() string {
	return "student"
}
