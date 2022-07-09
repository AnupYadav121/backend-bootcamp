package Controllers

import (
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var Input Models.Student
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	Student := Models.Student{FirstName: Input.FirstName, LastName: Input.LastName, DOB: Input.DOB, Address: Input.Address, SubjectMarksID: Input.SubjectMarksID}
	Utils.DoCreate(&Student)
	c.JSON(http.StatusOK, gin.H{"Students": Student})
}

func GetStudents(c *gin.Context) {
	var Students []Models.Student
	Utils.DoFind(&Students)
	c.JSON(http.StatusOK, gin.H{"Students": Students})
}

func FindStudent(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Students": Student})
}

func UpdateStudent(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student id to be updated record not found"})
		return
	}
	var NewStudent Models.UpdatedStudent
	errNew := c.ShouldBindJSON(&NewStudent)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "New student request body is not correct"})
		return
	}

	Utils.DoUpdate(&Student, NewStudent)
	c.JSON(http.StatusOK, NewStudent)
}

func DeleteStudent(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student id to be deleted, record not found"})
		return
	}

	newErr := Utils.DoDelete(&Student)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the Student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}

func CreateSubjectMarks(c *gin.Context) {
	var Input Models.SubjectMarks
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	SubjectMarks := Models.SubjectMarks{Subject: Input.Subject, Marks: Input.Marks, Grade: Input.Grade}
	Utils.DoCreateSubject(&SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"SubjectMarkss": SubjectMarks})
}

func GetSubjectMarks(c *gin.Context) {
	var SubjectMarkss []Models.SubjectMarks
	Utils.DoFindSubject(&SubjectMarkss)
	c.JSON(http.StatusOK, gin.H{"SubjectMarkss": SubjectMarkss})
}

func FindSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentSubject(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "SubjectMarks Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SubjectMarkss": SubjectMarks})
}

func UpdateSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentSubject(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "SubjectMarks id to be updated record not found"})
		return
	}
	var NewSubjectMarks Models.UpdatedSubjectMarks
	errNew := c.ShouldBindJSON(&NewSubjectMarks)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "New student request body is not correct"})
		return
	}

	Utils.DoUpdateSubject(&SubjectMarks, NewSubjectMarks)
	c.JSON(http.StatusOK, NewSubjectMarks)
}

func DeleteSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentSubject(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "SubjectMarks id to be deleted, record not found"})
		return
	}

	newErr := Utils.DoDeleteSubject(&SubjectMarks)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the SubjectMarks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}
