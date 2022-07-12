package Service

import (
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type InterfaceMarks interface {
	SaveMarks()
	GetMarks()
	GetMyMarks()
	FindMarks()
	UpdateMarks()
	DeleteMarks()
}

type Marks struct {
	db Utils.DB
}

func NewMarks() *Marks {
	return &Marks{}
}

func (marks *Marks) SaveMarks(c *gin.Context) {
	var Input Models.SubjectMarks
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	errr := Input.ValidateMarks()
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errr.Error()})
		return
	}

	var Student Models.Student
	StudentData, errNew := marks.db.IsPresent(strconv.Itoa(int(Input.StudentId)), &Student)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student with id Not Found"})
		fmt.Println(StudentData)
		return
	}

	marks.db.DoCreateMark(&Input)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": Input})
}

func (marks *Marks) GetMarks(c *gin.Context) {
	var SubjectMarks []Models.SubjectMarks
	marks.db.DoFindMarks(&SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": SubjectMarks})
}

func (marks *Marks) GetMyMarks(c *gin.Context) {
	var Student Models.Student
	StudentData, err := marks.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student Not Found"})
		return
	}
	fmt.Println(StudentData)
	var SubjectMarks []Models.SubjectMarks
	marks.db.MyMarks(c.Param("id"), &SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"My SubjectMarks": SubjectMarks})
}

func (marks *Marks) FindMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := marks.db.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "SubjectMarks Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SubjectMark": SubjectMarks})
}

func (marks *Marks) UpdateMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := marks.db.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "SubjectMarks id to be updated record not found"})
		return
	}

	var NewSubjectMarks Models.UpdatedSubjectMarks
	errNew := c.ShouldBindJSON(&NewSubjectMarks)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "New student request body is not correct"})
		return
	}

	errr := NewSubjectMarks.ValidateUpdatedMarks()
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errr.Error()})
		return
	}

	var Student Models.Student
	StudentData, newErr := marks.db.IsPresent(strconv.Itoa(int(NewSubjectMarks.StudentId)), &Student)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student with id Not Found"})
		return
	}
	fmt.Println(StudentData)
	marks.db.DoUpdateMark(&SubjectMarks, NewSubjectMarks)
	c.JSON(http.StatusOK, NewSubjectMarks)
}

func (marks *Marks) DeleteMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := marks.db.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "SubjectMarks id to be deleted, record not found"})
		return
	}

	newErr := marks.db.DoDeleteMark(&SubjectMarks)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the SubjectMarks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}
