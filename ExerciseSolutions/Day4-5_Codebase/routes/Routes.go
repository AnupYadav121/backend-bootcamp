package routes

import (
	Controller "Day4-5_Codebase/controller"
	Utils "Day4-5_Codebase/db_utils"
	"Day4-5_Codebase/middleware"
	"Day4-5_Codebase/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/Home")

	withCustomer := Controller.NewCustomer(service.NewCustomerService(Utils.GetDB()))
	withProduct := Controller.NewProduct(service.NewProductService(Utils.GetDB()))
	{
		grp1.POST("product/:retailerID", middleware.BasicAuthRetailer, withProduct.CreateProduct)
		grp1.GET("product/:retailerID/:id", middleware.BasicAuthRetailer, withProduct.FindProduct)
		grp1.GET("products/:retailerID", middleware.BasicAuthRetailer, withProduct.GetProducts)
		grp1.PUT("product/:retailerID/:id", middleware.BasicAuthRetailer, withProduct.UpdateProduct)
		grp1.DELETE("product/:retailerID/:id", middleware.BasicAuthRetailer, withProduct.DeleteProduct)
		grp1.POST("retailer", withProduct.AuthRetailer)
		grp1.POST("retailer/:retailerID", middleware.BasicAuthRetailer, withProduct.RemoveAuthRetailer)

		grp1.POST("order/:customerID", middleware.BasicAuthCustomer, withCustomer.CreateOrder)
		grp1.POST("orders/:customerID", middleware.BasicAuthCustomer, withCustomer.CreateMultipleOrder)
		grp1.GET("order/:customerID/:id", middleware.BasicAuthCustomer, withCustomer.FindOrderUpdates)
		grp1.POST("order/:customerID/:id", middleware.BasicAuthCustomer, withCustomer.SetOrderStatus)
		grp1.GET("orders/:customerID", middleware.BasicAuthCustomer, withCustomer.GetMyOrders)
		grp1.GET("transactionOrders/:retailerID", middleware.BasicAuthRetailer, withProduct.GetAllTransactions)
		grp1.POST("user", withCustomer.AuthCustomer)
		grp1.DELETE("user/:customerID", middleware.BasicAuthCustomer, withCustomer.RemoveAuthCustomer)
	}
	return r
}
