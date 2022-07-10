// Routes/Routes.go
package Routes

import (
	"github.com/gin-gonic/gin"
	"practice-bootcamp/Controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")
	{
		grp1.POST("Book", Controllers.CreateBook)
		grp1.GET("Book/:id", Controllers.FindBook)
		grp1.GET("Books", Controllers.GetBooks)
		grp1.PUT("Book/:id", Controllers.UpdateBook)
		grp1.DELETE("Book/:id", Controllers.DeleteBook)
	}
	return r
}
