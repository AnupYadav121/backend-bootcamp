package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/db_utils"
	"july8Files/dto"
	"july8Files/models"
	"net/http"
	"strconv"
)

type OrderServiceInterface interface {
	AuthCustomer(c *gin.Context, customer *models.Customer)
	RemoveAuthCustomer(c *gin.Context, customer *models.Customer)
	IsAuthCustomer(c *gin.Context, customer *models.Customer)
	OrderCreation(c *gin.Context, newOrder *models.Order)
	SetStatus(c *gin.Context, orderStatus *models.Order)
	FindUpdate(c *gin.Context, orderExist *models.Order)
	GetMyOrders(c *gin.Context, customerExist *models.Customer)
	MultipleOrderCreation(c *gin.Context, newOrder *dto.Order)
}

type OrderService struct {
	db Utils.InterfaceDB
}

func NewCustomerService(db Utils.InterfaceDB) *OrderService {
	return &OrderService{db}
}

func (cs *OrderService) AuthCustomer(c *gin.Context, customer *models.Customer) {
	errNw := customer.CustomerValidate()
	if errNw != nil {
		c.JSON(http.StatusBadRequest, errNw.Error())
		return
	}
	cs.db.DoCreateC(customer)
}

func (cs *OrderService) RemoveAuthCustomer(c *gin.Context, customer *models.Customer) {
	err := cs.db.IsPresentC(c.Param("id"), customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var orderUpdates models.Order
	newErr := cs.db.IsPresentCU(c.Param("id"), &orderUpdates)
	if newErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can not remove authentication, One order is process for this customer id"})
		return
	}

	errNew := cs.db.DoDeleteC(customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
}

func (cs *OrderService) IsAuthCustomer(c *gin.Context, customer *models.Customer) {
	err := cs.db.IsPresentC(c.Param("id"), customer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": "Customer is not Authenticated"})
		return
	}
}

func (cs *OrderService) OrderCreation(c *gin.Context, newOrder *models.Order) {
	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": nwErr})
		return
	}

	var customer models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	var product models.Product
	newErr := cs.db.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given product id does not exist"})
		return
	}

	cs.db.DoCreateO(newOrder)
}

func (cs *OrderService) MultipleOrderCreation(c *gin.Context, newOrder *dto.Order) {
	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": nwErr})
		return
	}

	var customer models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given customer id is not authenticated"})
		return
	}

	for i := 0; i < len(newOrder.ProductID); i++ {
		var product models.Product
		newErr := cs.db.IsPresent(strconv.Itoa(int(newOrder.ProductID[i])), &product)
		if newErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Given one of the product id does not exist"})
			return
		}
		orderNew := models.Order{ProductID: newOrder.ProductID[i], CustomerID: newOrder.CustomerID, Quantity: newOrder.Quantity[i]}
		cs.db.DoCreateO(&orderNew)
	}
}

func (cs *OrderService) SetStatus(c *gin.Context, orderStatus *models.Order) {
	var orderExist models.Order
	errNew := cs.db.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given Order id to be updated do not exist"})
		return
	}
	orderUpdated := models.Order{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}
	cs.db.DoCreateOU(&orderUpdated)
	c.JSON(http.StatusOK, gin.H{"Order": orderUpdated})
}

func (cs *OrderService) FindUpdate(c *gin.Context, orderExist *models.Order) {
	err := cs.db.IsPresentO(c.Param("id"), orderExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Either order with this id does not exist or order status is not updated by retailer yet."})
		return
	}

}

func (cs *OrderService) GetMyOrders(c *gin.Context, customerExist *models.Customer) {
	err := cs.db.IsPresentC(c.Param("id"), customerExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Sorry ! you are not authenticated ,you can not have any orders"})
		return
	}

	var myOrders []models.Order
	errNew := cs.db.IsCustomerOrder(c.Param("id"), &myOrders)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"My orders update": myOrders})
}
