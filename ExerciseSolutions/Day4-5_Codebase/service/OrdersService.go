package service

import (
	"Day4-5_Codebase/db_utils"
	"Day4-5_Codebase/dto"
	"Day4-5_Codebase/models"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type OrderServiceInterface interface {
	AuthCustomer(customer *models.Customer) (c *models.Customer, er error)
	RemoveAuthCustomer(c *gin.Context, customer *models.Customer) (er error)
	IsAuthCustomer(c *gin.Context, customer *models.Customer) (cst *models.Customer, er error)
	OrderCreation(newOrder *models.Order) (o *models.Order, er error)
	SetStatus(c *gin.Context, orderStatus *models.Order) (order *models.Order, er error)
	FindUpdate(c *gin.Context, orderExist *models.Order) (o *models.Order, er error)
	GetMyOrders(c *gin.Context, customerExist *models.Customer) (orders *[]models.Order, er error)
	MultipleOrderCreation(newOrder *dto.Order) (er error)
}

type OrderService struct {
	db Utils.InterfaceDB
}

func NewCustomerService(db Utils.InterfaceDB) *OrderService {
	return &OrderService{db}
}

func (cs *OrderService) AuthCustomer(customer *models.Customer) (c *models.Customer, er error) {
	errNw := customer.CustomerValidate()
	if errNw != nil {
		var tmpCustomer models.Customer
		return &tmpCustomer, errors.New("provided customer body is invalid")
	}
	err := cs.db.DoCreateC(customer)
	if err != nil {
		var tmpCustomer models.Customer
		return &tmpCustomer, err
	}
	return customer, nil
}

func (cs *OrderService) RemoveAuthCustomer(c *gin.Context, customer *models.Customer) (er error) {
	err := cs.db.IsPresentC(c.Param("customerID"), customer)
	if err != nil {
		return errors.New("customer id not found")
	}

	var orderUpdates models.Order
	newErr := cs.db.IsPresentCU(c.Param("customerID"), &orderUpdates)
	if newErr == nil {
		return errors.New("can not remove authentication, customer is associated with a order")
	}

	errNew := cs.db.DoDeleteC(customer)
	if errNew != nil {
		return errNew
	}
	return nil
}

func (cs *OrderService) IsAuthCustomer(c *gin.Context, customer *models.Customer) (cst *models.Customer, er error) {
	err := cs.db.IsPresentC(c.Param("customerID"), customer)

	if err != nil {
		var tmpCustomer models.Customer
		return &tmpCustomer, errors.New("customer id is not authenticated")
	}
	return customer, nil
}

func (cs *OrderService) OrderCreation(newOrder *models.Order) (o *models.Order, er error) {
	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		var order models.Order
		return &order, errors.New("provided order body is invalid")
	}

	var customer models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		var order models.Order
		return &order, errors.New("customer id given is not authenticated")
	}

	var product models.Product
	newErr := cs.db.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		var order models.Order
		return &order, errors.New("product id given is not authenticated")
	}

	erNew := cs.db.DoCreateO(newOrder)
	if errNew != nil {
		var order models.Order
		return &order, erNew
	}
	return newOrder, nil
}

func (cs *OrderService) MultipleOrderCreation(newOrder *dto.Order) (er error) {
	nwErr := newOrder.OrderValidate()
	if nwErr != nil {
		return errors.New("provided order body is invalid")
	}

	var customer models.Customer
	errNew := cs.db.IsPresentC(strconv.Itoa(int(newOrder.CustomerID)), &customer)
	if errNew != nil {
		return errors.New("customer id given is not authenticated")
	}

	for i := 0; i < len(newOrder.ProductID); i++ {
		var product models.Product
		newErr := cs.db.IsPresent(strconv.Itoa(newOrder.ProductID[i]), &product)
		if newErr != nil {
			return errors.New("one of the product id given does not exist")
		}
		orderNew := models.Order{ProductID: newOrder.ProductID[i], CustomerID: newOrder.CustomerID, Quantity: newOrder.Quantity[i]}
		erNew := cs.db.DoCreateO(&orderNew)
		if errNew != nil {
			return erNew
		}
	}
	return nil
}

func (cs *OrderService) SetStatus(c *gin.Context, orderStatus *models.Order) (order *models.Order, er error) {
	var orderExist models.Order
	errNew := cs.db.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		var tmpOrder models.Order
		return &tmpOrder, errors.New("order id given does not exist")
	}

	orderUpdated := models.Order{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}

	erNew := cs.db.DoCreateOU(&orderUpdated)
	if errNew != nil {
		var tmpOrder models.Order
		return &tmpOrder, erNew
	}
	return &orderUpdated, nil
}

func (cs *OrderService) FindUpdate(c *gin.Context, orderExist *models.Order) (o *models.Order, er error) {
	err := cs.db.IsPresentO(c.Param("id"), orderExist)
	if err != nil {
		var order models.Order
		return &order, errors.New("order id given does not exist")
	}
	return orderExist, nil
}

func (cs *OrderService) GetMyOrders(c *gin.Context, customerExist *models.Customer) (orders *[]models.Order, er error) {
	err := cs.db.IsPresentC(c.Param("customerID"), customerExist)
	if err != nil {
		var tmpOrders []models.Order
		return &tmpOrders, errors.New("sorry ! you are not authenticated ,you can not have any orders")
	}

	var myOrders []models.Order
	errNew := cs.db.IsCustomerOrder(c.Param("customerID"), &myOrders)
	if errNew != nil {
		var tmpOrders []models.Order
		return &tmpOrders, errors.New("you do not have any orders")
	}
	return &myOrders, nil
}
