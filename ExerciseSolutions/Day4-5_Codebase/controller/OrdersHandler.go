package Controllers

import (
	"Day4-5_Codebase/dto"
	"Day4-5_Codebase/models"
	"Day4-5_Codebase/mutex"
	"Day4-5_Codebase/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type CustomerInterface interface {
	AuthCustomer(c *gin.Context)
	RemoveAuthCustomer(c *gin.Context)
	IsCustomerAuthenticated(c *gin.Context) (cust *models.Customer, err error)
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

func (ch *CustomerHandle) AuthCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	savedCustomer, errNew := ch.os.AuthCustomer(&customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer": savedCustomer})
}

func (ch *CustomerHandle) RemoveAuthCustomer(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("customer_id" + c.Param("customerID")); isAvailable == false {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "customer id is being deleted, wait,you can try this on another id"})
		time.Sleep(2 * time.Second)
		return
	}
	defer mutex.Mutex.UnLock("customer_id" + c.Param("customerID"))

	var customer models.Customer
	err := ch.os.RemoveAuthCustomer(c, &customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func (ch *CustomerHandle) IsCustomerAuthenticated(c *gin.Context) (cust *models.Customer, err error) {
	var customer models.Customer
	resCustomer, resErr := ch.os.IsAuthCustomer(c, &customer)
	if resErr != nil {
		var tmpCustomer models.Customer
		return &tmpCustomer, resErr
	}
	return resCustomer, nil
}

func (ch *CustomerHandle) CreateOrder(c *gin.Context) {
	var newOrder models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	savedOrder, errNew := ch.os.OrderCreation(&newOrder)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order": savedOrder})
}

func (ch *CustomerHandle) CreateMultipleOrder(c *gin.Context) {
	var newOrder dto.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	errNew := ch.os.MultipleOrderCreation(&newOrder)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
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
	updatedOrder, errNew := ch.os.SetStatus(c, &orderStatus)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"New Order": updatedOrder})
}

func (ch *CustomerHandle) FindOrderUpdates(c *gin.Context) {
	var orderExist models.Order
	updatedOrder, errNew := ch.os.FindUpdate(c, &orderExist)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order updated": updatedOrder})
}

func (ch *CustomerHandle) GetMyOrders(c *gin.Context) {
	var customerExist models.Customer
	orders, err := ch.os.GetMyOrders(c, &customerExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"my orders": orders})
}
