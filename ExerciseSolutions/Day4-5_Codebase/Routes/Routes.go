package Routes

import "C"
import (
	"github.com/gin-gonic/gin"
	Controller "july8Files/Controller"
	Utils "july8Files/DB_Utils"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")

	withCustomer := Controller.NewCustomer(Utils.GetDB())
	withProduct := Controller.NewProduct(Utils.GetDB())
	{
		grp1.POST("product", withProduct.CreateProduct)
		grp1.GET("product/:id", withProduct.FindProduct)
		grp1.GET("products", withProduct.GetProducts)
		grp1.PUT("product/:id", withProduct.UpdateProduct)
		grp1.DELETE("product/:id", withProduct.DeleteProduct)
		grp1.GET("orders", withProduct.GetAllTransactions)

		grp1.POST("retailer", withProduct.AuthRetailer)
		grp1.POST("retailer/:id", withProduct.RemoveAuthRetailer)
		grp1.GET("retailer/:id", withProduct.IsRetailerAuthenticated)

		grp1.POST("order", withCustomer.CreateOrder)
		grp1.POST("orders", withCustomer.CreateMultipleOrder)
		grp1.GET("order/:id", withCustomer.FindOrderUpdates)
		grp1.POST("order/:id", withCustomer.SetOrderStatus)
		grp1.GET("orders/:id", withCustomer.GetMyOrders)
		grp1.POST("user", withCustomer.CreateCustomer)
		grp1.DELETE("user/:id", withCustomer.DeleteCustomer)
		grp1.GET("user/:id", withCustomer.IsCustomerAuthenticated)
	}
	return r
}
