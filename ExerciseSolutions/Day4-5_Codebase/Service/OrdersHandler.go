package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/DB_Utils"
	"july8Files/Models"
	"net/http"
	"strconv"
)

type CustomerInterface interface {
	CreateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
	IsCustomerAuthenticated(c *gin.Context)
	CreateOrder(c *gin.Context)
	SetOrderStatus(c *gin.Context)
	FindOrderUpdates(c *gin.Context)
	GetMyOrders(c *gin.Context)
}

type CustomerHandle struct {
	db Utils.InterfaceDB
}

func NewCustomer(db Utils.InterfaceDB) *CustomerHandle {
	return &CustomerHandle{db}
}

func (ch *CustomerHandle) CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ch.db.DoCreateC(&customer)
	c.JSON(http.StatusOK, gin.H{"Customer": customer})
}

func (ch *CustomerHandle) DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	err := ch.db.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	var orderUpdates Models.OrderUpdated
	newErr := ch.db.IsPresentCU(c.Param("id"), &orderUpdates)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Can not remove authentication, One order is process for this customer id"})
		return
	}

	errNew := ch.db.DoDeleteC(&customer)
	if errNew != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func (ch *CustomerHandle) IsCustomerAuthenticated(c *gin.Context) {
	var customer Models.Customer
	err := ch.db.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Customer is not Authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Customer is Authenticated"})
}

func (ch *CustomerHandle) CreateOrder(c *gin.Context) {
	var newOrder Models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var customer Models.Customer
	errNew := ch.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	var product Models.Product
	newErr := ch.db.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given product id does not exist"})
		return
	}

	ch.db.DoCreateO(&newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func (ch *CustomerHandle) SetOrderStatus(c *gin.Context) {
	var orderStatus Models.OrderUpdated
	err := c.ShouldBindJSON(&orderStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var orderExist Models.Order
	errNew := ch.db.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Given Order id to be updated do not exist"})
		return
	}
	orderUpdated := Models.OrderUpdated{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}

	ch.db.DoCreateOU(&orderUpdated)
	c.JSON(http.StatusOK, gin.H{"Order": orderUpdated})
}

func (ch *CustomerHandle) FindOrderUpdates(c *gin.Context) {
	var orderExist Models.OrderUpdated
	err := ch.db.IsPresentOU(c.Param("id"), &orderExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Either order with this id does not exist or order status is not updated by retailer yet."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order updates": orderExist})
}

func (ch *CustomerHandle) GetMyOrders(c *gin.Context) {
	var customerExist Models.Customer
	err := ch.db.IsPresentC(c.Param("id"), &customerExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Sorry ! you are not authenticated ,you can not have any orders"})
		return
	}

	var myOrders []Models.Order
	errNew := ch.db.IsCustomerOrder(c.Param("id"), &myOrders)
	if errNew != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"My orders update": myOrders})
}
