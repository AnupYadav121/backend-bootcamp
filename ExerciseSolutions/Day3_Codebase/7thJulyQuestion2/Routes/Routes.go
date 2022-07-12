// Routes/Routes.go
package Routes

import "C"
import (
	"7thJulyQuestion2/Service"
	"7thJulyQuestion2/Utils"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")

	marksEndpointContext := Service.NewMarks()
	studentEndpointContext := Service.NewStudent(Utils.GetDB())

	{
		grp1.POST("Student", studentEndpointContext.CreateStudent)
		grp1.GET("Student/:id", studentEndpointContext.FindStudent)
		grp1.GET("MyInfo/:id", studentEndpointContext.GetStudentInfo)
		grp1.GET("Students", studentEndpointContext.GetStudents)
		grp1.PUT("Student/:id", studentEndpointContext.UpdateStudent)
		grp1.DELETE("Student/:id", studentEndpointContext.DeleteStudent)

		grp1.POST("SubjectMarks", marksEndpointContext.SaveMarks)
		grp1.GET("SubjectMarks/:id", marksEndpointContext.FindMarks)
		grp1.GET("MyMarks/:id", marksEndpointContext.GetMyMarks)
		grp1.GET("SubjectMarks", marksEndpointContext.GetMyMarks)
		grp1.PUT("SubjectMarks/:id", marksEndpointContext.UpdateMarks)
		grp1.DELETE("SubjectMarks/:id", marksEndpointContext.DeleteMarks)
	}
	return r
}
