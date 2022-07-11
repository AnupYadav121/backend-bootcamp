package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/Models"
	"july8Files/Utils"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var product Models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Utils.DoCreate(&product)
	c.JSON(http.StatusOK, gin.H{"Product Saved": product})
}

func FindProduct(c *gin.Context) {
	var product Models.Product
	err := Utils.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product with id Found": product})
}

func GetProducts(c *gin.Context) {
	var products []Models.Product
	Utils.DoFind(&products)
	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func UpdateProduct(c *gin.Context) {
	var productNew Models.UpdatedProduct
	err := c.ShouldBindJSON(&productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var ProductExist Models.Product
	errNew := Utils.IsPresent(c.Param("id"), &ProductExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errNew.Error()})
		return
	}

	if ProductExist.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not find your order with id to be updated"})
		return
	}

	Utils.DoUpdate(&ProductExist, productNew)
	c.JSON(http.StatusOK, gin.H{"ProductNew": productNew})
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	err := Utils.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	errNew := Utils.DoDelete(&product)
	if errNew != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product Deleted Successfully": &product})
}

func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Utils.DoCreateC(&customer)
	c.JSON(http.StatusOK, gin.H{"Customer": customer})
}

func DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	err := Utils.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	var orderUpdates Models.OrderUpdated
	newErr := Utils.IsPresentCU(c.Param("id"), &orderUpdates)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Can not remove authentication, One order is process for this customer id"})
		return
	}

	errNew := Utils.DoDeleteC(&customer)
	if errNew != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func IsCustomerAuthenticated(c *gin.Context) {
	var customer Models.Customer
	err := Utils.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Customer is not Authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Customer is Authenticated"})
}

func CreateOrder(c *gin.Context) {
	var newOrder Models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var customer Models.Customer
	errNew := Utils.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	var product Models.Product
	newErr := Utils.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given product id does not exist"})
		return
	}

	Utils.DoCreateO(&newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func SetOrderStatus(c *gin.Context) {
	var orderStatus Models.OrderUpdated
	err := c.ShouldBindJSON(&orderStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var orderExist Models.Order
	errNew := Utils.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given Order id to be updated do not exist"})
		return
	}
	orderUpdated := Models.OrderUpdated{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}

	Utils.DoCreateOU(&orderUpdated)
	c.JSON(http.StatusOK, gin.H{"Order": orderUpdated})
}

func FindOrderUpdates(c *gin.Context) {
	var orderExist Models.OrderUpdated
	err := Utils.IsPresentOU(c.Param("id"), &orderExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Either order with this id does not exist or order status is not updated by retailer yet."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order updates": orderExist})
}

func GetMyOrders(c *gin.Context) {
	var customerExist Models.Customer
	err := Utils.IsPresentC(c.Param("id"), &customerExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Sorry ! you are not authenticated ,you can not have any orders"})
		return
	}

	var myOrders []Models.Order
	errNew := Utils.IsCustomerOrder(c.Param("id"), &myOrders)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"My orders update": myOrders})
}

func GetAllTransactions(c *gin.Context) {
	var allOrders []Models.OrderUpdated
	errNew := Utils.FindAllOrders(&allOrders)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errNew.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All Order updates": allOrders})
}
