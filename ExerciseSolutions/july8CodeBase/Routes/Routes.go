package Routes

import (
	"github.com/gin-gonic/gin"
	"july8Files/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")
	{
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.FindProduct)
		grp1.GET("products", Controllers.GetProducts)
		grp1.PUT("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)

		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.FindOrderUpdates)

		grp1.POST("user", Controllers.CreateCustomer)
		grp1.DELETE("user/:id", Controllers.DeleteCustomer)
	}
	return r
}
