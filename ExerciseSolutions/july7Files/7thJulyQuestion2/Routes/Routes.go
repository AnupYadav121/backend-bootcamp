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
	}
	return r
}
