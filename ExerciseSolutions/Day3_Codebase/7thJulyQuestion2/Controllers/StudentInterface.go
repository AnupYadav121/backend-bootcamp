package Controllers

import (
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type InterfaceStudent interface {
	CreateStudent()
	GetStudents()
	FindStudent()
	GetStudentInfo()
	UpdateStudent()
	DeleteStudent()
}

type AnyStudent struct {
	db Utils.DB
}

func NewStudent() *AnyStudent {
	return &AnyStudent{}
}

func (anyStudent *AnyStudent) CreateStudent(c *gin.Context) {
	var Input Models.Student
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	errr := Input.ValidateStudent()
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errr.Error()})
		return
	}

	anyStudent.db.DoCreate(&Input)
	c.JSON(http.StatusOK, gin.H{"Students": Input})
}

func (anyStudent *AnyStudent) GetStudents(c *gin.Context) {
	var Students []Models.Student
	anyStudent.db.DoFind(&Students)

	c.JSON(http.StatusOK, gin.H{"Students": &Students})
}

func (anyStudent *AnyStudent) FindStudent(c *gin.Context) {
	var Student Models.Student
	err := anyStudent.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Student": Student})
}

func (anyStudent *AnyStudent) GetStudentInfo(c *gin.Context) {
	var Student Models.Student
	err := anyStudent.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student Not Found"})
		return
	}

	var SubjectMarks []Models.SubjectMarks
	anyStudent.db.MyMarks(c.Param("id"), &SubjectMarks)

	MyInfo := Models.StudentInfo{ID: Student.ID, FirstName: Student.FirstName, LastName: Student.LastName, DOB: Student.DOB, Address: Student.Address, Marks: SubjectMarks}
	c.JSON(http.StatusOK, gin.H{"My Complete Info": MyInfo})
}

func (anyStudent *AnyStudent) UpdateStudent(c *gin.Context) {
	var Student Models.Student
	err := anyStudent.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student id to be updated record not found"})
		return
	}

	var NewStudent Models.UpdatedStudent
	errNew := c.ShouldBindJSON(&NewStudent)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "New student request body is not correct"})
		return
	}

	errr := NewStudent.ValidateUpdatedStudent()
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errr.Error()})
		return
	}

	anyStudent.db.DoUpdate(&Student, NewStudent)
	c.JSON(http.StatusOK, NewStudent)
}

func (anyStudent *AnyStudent) DeleteStudent(c *gin.Context) {
	var Student Models.Student
	err := anyStudent.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student id to be deleted, record not found"})
		return
	}

	var SubjectMarks []Models.SubjectMarks
	errNew := anyStudent.db.IsMyMark(c.Param("id"), &SubjectMarks)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student Id is related to Subjects Marks. Sorry !, Can not delete it"})
		return
	}

	newErr := anyStudent.db.DoDelete(&Student)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the Student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}
