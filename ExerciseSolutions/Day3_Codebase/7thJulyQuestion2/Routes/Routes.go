// Routes/Routes.go
package Routes

import "C"
import (
	"7thJulyQuestion2/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")

	marksEndpointContext := Controllers.NewMarks()
	studentEndpointContext := Controllers.NewStudent()

	{
		grp1.POST("Student", studentEndpointContext.CreateStudent)
		grp1.GET("Student/:id", studentEndpointContext.FindStudent)
		grp1.GET("MyInfo/:id", studentEndpointContext.GetStudentInfo)
		grp1.GET("Students", studentEndpointContext.GetStudents)
		grp1.PUT("Student/:id", studentEndpointContext.UpdateStudent)
		grp1.DELETE("Student/:id", studentEndpointContext.DeleteStudent)

		grp1.POST("SubjectMarks", marksEndpointContext.CreateSubjectMarks)
		grp1.GET("SubjectMarks/:id", marksEndpointContext.FindSubjectMarks)
		grp1.GET("MyMarks/:id", marksEndpointContext.GetMyMarks)
		grp1.GET("SubjectMarks", marksEndpointContext.GetSubjectMarks)
		grp1.PUT("SubjectMarks/:id", marksEndpointContext.UpdateSubjectMarks)
		grp1.DELETE("SubjectMarks/:id", marksEndpointContext.DeleteSubjectMarks)
	}
	return r
}
