package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/dto"
	"july8Files/models"
	"july8Files/mutex"
	"july8Files/service"
	"net/http"
	"time"
)

type CustomerInterface interface {
	CreateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
	IsCustomerAuthenticated(c *gin.Context)
	CreateOrder(c *gin.Context)
	SetOrderStatus(c *gin.Context)
	FindOrderUpdates(c *gin.Context)
	GetMyOrders(c *gin.Context)
	CreateMultipleOrder(c *gin.Context)
}

type CustomerHandle struct {
	os service.OrderServiceInterface
}

func NewCustomer(os service.OrderServiceInterface) *CustomerHandle {
	return &CustomerHandle{os}
}

func (ch *CustomerHandle) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ch.os.AuthCustomer(c, &customer)
	c.JSON(http.StatusOK, gin.H{"Customer": customer})
}

func (ch *CustomerHandle) DeleteCustomer(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("customer_id" + c.Param("id")); isAvailable == false {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "customer id is being deleted, wait"})
		return
	}
	time.Sleep(2 * time.Second)
	defer mutex.Mutex.UnLock("customer_id" + c.Param("id"))

	var customer models.Customer
	ch.os.RemoveAuthCustomer(c, &customer)
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func (ch *CustomerHandle) IsCustomerAuthenticated(c *gin.Context) {
	var customer models.Customer
	ch.os.IsAuthCustomer(c, &customer)
	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Customer is Authenticated"})
}

func (ch *CustomerHandle) CreateOrder(c *gin.Context) {
	var newOrder models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ch.os.OrderCreation(c, &newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func (ch *CustomerHandle) CreateMultipleOrder(c *gin.Context) {
	var newOrder dto.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ch.os.MultipleOrderCreation(c, &newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func (ch *CustomerHandle) SetOrderStatus(c *gin.Context) {
	var orderStatus models.Order
	err := c.ShouldBindJSON(&orderStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	nwErr := orderStatus.Status
	if nwErr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "please provide order status"})
		return
	}
	ch.os.SetStatus(c, &orderStatus)
}

func (ch *CustomerHandle) FindOrderUpdates(c *gin.Context) {
	var orderExist models.Order
	ch.os.FindUpdate(c, &orderExist)
	c.JSON(http.StatusOK, gin.H{"Order updates": orderExist})
}

func (ch *CustomerHandle) GetMyOrders(c *gin.Context) {
	var customerExist models.Customer
	ch.os.GetMyOrders(c, &customerExist)
}
