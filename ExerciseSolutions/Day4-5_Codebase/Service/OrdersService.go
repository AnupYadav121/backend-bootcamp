package Service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/DB_Utils"
	"july8Files/Dto"
	"july8Files/Models"
	"net/http"
	"strconv"
)

type OrderServiceInterface interface {
	AuthCustomer(c *gin.Context)
	RemoveAuthCustomer(c *gin.Context)
	IsAuthCustomer(c *gin.Context)
	OrderCreation(c *gin.Context)
	SetStatus(c *gin.Context)
	FindUpdate(c *gin.Context)
	GetMyOrders(c *gin.Context)
	MultipleOrderCreation(c *gin.Context)
}

type OrderService struct {
	db Utils.InterfaceDB
}

func NewCustomerService(db Utils.InterfaceDB) *OrderService {
	return &OrderService{db}
}

func (cs *OrderService) AuthCustomer(c *gin.Context) {
	var customer Models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	errNw := customer.CustomerValidate()
	if errNw != nil {
		c.JSON(http.StatusBadRequest, errNw.Error())
		return
	}

	cs.db.DoCreateC(&customer)
	c.JSON(http.StatusOK, gin.H{"Customer": customer})
}

func (cs *OrderService) RemoveAuthCustomer(c *gin.Context) {
	var customer Models.Customer
	err := cs.db.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var orderUpdates Models.Order
	newErr := cs.db.IsPresentCU(c.Param("id"), &orderUpdates)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can not remove authentication, One order is process for this customer id"})
		return
	}

	errNew := cs.db.DoDeleteC(&customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func (cs *OrderService) IsAuthCustomer(c *gin.Context) {
	var customer Models.Customer
	err := cs.db.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": "Customer is not Authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Customer is Authenticated"})
}

func (cs *OrderService) OrderCreation(c *gin.Context) {
	var newOrder Models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": nwErr})
		return
	}

	var customer Models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	var product Models.Product
	newErr := cs.db.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given product id does not exist"})
		return
	}

	cs.db.DoCreateO(&newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func (cs *OrderService) MultipleOrderCreation(c *gin.Context) {
	var newOrder Dto.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": nwErr})
		return
	}

	var customer Models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	for i := 0; i < len(newOrder.ProductID); i++ {
		var product Models.Product
		newErr := cs.db.IsPresent(strconv.Itoa(int(newOrder.ProductID[i])), &product)
		if newErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Given one of the product id does not exist"})
			return
		}
		orderNew := Models.Order{ProductID: newOrder.ProductID[i], CustomerID: newOrder.CustomerID, Quantity: newOrder.Quantity[i]}
		cs.db.DoCreateO(&orderNew)
	}

	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func (cs *OrderService) SetStatus(c *gin.Context) {
	var orderStatus Models.Order
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

	var orderExist Models.Order
	errNew := cs.db.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given Order id to be updated do not exist"})
		return
	}
	orderUpdated := Models.Order{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}

	cs.db.DoCreateOU(&orderUpdated)
	c.JSON(http.StatusOK, gin.H{"Order": orderUpdated})
}

func (cs *OrderService) FindUpdate(c *gin.Context) {
	var orderExist Models.Order
	err := cs.db.IsPresentO(c.Param("id"), &orderExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Either order with this id does not exist or order status is not updated by retailer yet."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order updates": orderExist})
}

func (cs *OrderService) GetMyOrders(c *gin.Context) {
	var customerExist Models.Customer
	err := cs.db.IsPresentC(c.Param("id"), &customerExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Sorry ! you are not authenticated ,you can not have any orders"})
		return
	}

	var myOrders []Models.Order
	errNew := cs.db.IsCustomerOrder(c.Param("id"), &myOrders)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"My orders update": myOrders})
}
