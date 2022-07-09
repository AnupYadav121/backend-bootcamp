// Routes/Routes.go
package Routes

import (
	"7thJulyQuestion2/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")
	{
		grp1.POST("Student", Controllers.CreateStudent)
		grp1.GET("Student/:id", Controllers.FindStudent)
		grp1.GET("Students", Controllers.GetStudents)
		grp1.PUT("Student/:id", Controllers.UpdateStudent)
		grp1.DELETE("Student/:id", Controllers.DeleteStudent)

		grp1.POST("SubjectMarks", Controllers.CreateSubjectMarks)
		grp1.GET("SubjectMarks/:id", Controllers.FindSubjectMarks)
		grp1.GET("SubjectMarks", Controllers.GetSubjectMarks)
		grp1.PUT("SubjectMarks/:id", Controllers.UpdateSubjectMarks)
		grp1.DELETE("SubjectMarks/:id", Controllers.DeleteSubjectMarks)
	}
	return r
}
