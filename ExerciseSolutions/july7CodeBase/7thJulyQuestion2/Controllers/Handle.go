package Controllers

import (
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

func CreateStudent(c *gin.Context) {
	var Input Models.Student
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Utils.DoCreate(&Input)
	c.JSON(http.StatusOK, gin.H{"Students": Input})
}

func GetStudents(c *gin.Context) {
	var Students []Models.Student
	Utils.DoFind(&Students)

	c.JSON(http.StatusOK, gin.H{"Students": &Students})
}

func FindStudent(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Student": Student})
}

func GetStudentInfo(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student Not Found"})
		return
	}

	var SubjectMarks []Models.SubjectMarks
	Utils.MyMarks(c.Param("id"), &SubjectMarks)

	MyInfo := Models.StudentInfo{ID: Student.ID, FirstName: Student.FirstName, LastName: Student.LastName, DOB: Student.DOB, Address: Student.Address, Marks: SubjectMarks}
	c.JSON(http.StatusOK, gin.H{"My Complete Info": MyInfo})
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

	var SubjectMarks []Models.SubjectMarks
	errNew := Utils.IsMyMark(c.Param("id"), &SubjectMarks)
	if errNew == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student Id is related to Subjects Marks. Sorry !, Can not delete it"})
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

	var Student Models.Student
	errNew := Utils.IsPresent(strconv.Itoa(int(Input.StudentId)), &Student)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student with id Not Found"})
		return
	}

	Utils.DoCreateMark(&Input)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": Input})
}

func GetSubjectMarks(c *gin.Context) {
	var SubjectMarks []Models.SubjectMarks
	Utils.DoFindMarks(&SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"SubjectMarks": SubjectMarks})
}

func GetMyMarks(c *gin.Context) {
	var Student Models.Student
	err := Utils.IsPresent(c.Param("id"), &Student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student Not Found"})
		return
	}

	var SubjectMarks []Models.SubjectMarks
	Utils.MyMarks(c.Param("id"), &SubjectMarks)
	c.JSON(http.StatusOK, gin.H{"My SubjectMarks": SubjectMarks})
}

func FindSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "SubjectMarks Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SubjectMark": SubjectMarks})
}

func UpdateSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentMark(c.Param("id"), &SubjectMarks)
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

	var Student Models.Student
	newErr := Utils.IsPresent(strconv.Itoa(int(NewSubjectMarks.StudentId)), &Student)
	if newErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Student with id Not Found"})
		return
	}

	Utils.DoUpdateMark(&SubjectMarks, NewSubjectMarks)
	c.JSON(http.StatusOK, NewSubjectMarks)
}

func DeleteSubjectMarks(c *gin.Context) {
	var SubjectMarks Models.SubjectMarks
	err := Utils.IsPresentMark(c.Param("id"), &SubjectMarks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "SubjectMarks id to be deleted, record not found"})
		return
	}

	newErr := Utils.DoDeleteMark(&SubjectMarks)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the SubjectMarks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}
