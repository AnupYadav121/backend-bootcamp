package Controllers

import (
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type InterfaceMarks interface {
	CreateSubjectMarks()
	GetSubjectMarks()
	GetMyMarks()
	FindSubjectMarks()
	UpdateSubjectMarks()
	DeleteSubjectMarks()
}

type Marks struct {
	db Utils.DB
}

func NewMarks() *Marks {
	return &Marks{}
}

func (marks *Marks) CreateSubjectMarks(c *gin.Context) {
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
	errNew := marks.db.IsPresent(strconv.Itoa(int(Input.StudentId)), &Student)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student with id Not Found"})
		return
	}

	marks.db.DoCreateMark(&Input)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": Input})
}

func (marks *Marks) GetSubjectMarks(c *gin.Context) {
	var SubjectMarks []Models.SubjectMarks
	marks.db.DoFindMarks(&SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": SubjectMarks})
}

func (marks *Marks) GetMyMarks(c *gin.Context) {
	var Student Models.Student
	err := marks.db.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student Not Found"})
		return
	}

	var SubjectMarks []Models.SubjectMarks
	marks.db.MyMarks(c.Param("id"), &SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"My SubjectMarks": SubjectMarks})
}

func (marks *Marks) FindSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := marks.db.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "SubjectMarks Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SubjectMark": SubjectMarks})
}

func (marks *Marks) UpdateSubjectMarks(c *gin.Context) {
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
	newErr := marks.db.IsPresent(strconv.Itoa(int(NewSubjectMarks.StudentId)), &Student)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Student with id Not Found"})
		return
	}

	marks.db.DoUpdateMark(&SubjectMarks, NewSubjectMarks)
	c.JSON(http.StatusOK, NewSubjectMarks)
}

func (marks *Marks) DeleteSubjectMarks(c *gin.Context) {
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
