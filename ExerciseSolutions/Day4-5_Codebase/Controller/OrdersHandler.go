package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	Utils "july8Files/DB_Utils"
	"july8Files/Service"
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
	db Utils.InterfaceDB
}

func NewCustomer(db Utils.InterfaceDB) *CustomerHandle {
	return &CustomerHandle{db}
}

func (ch *CustomerHandle) CreateCustomer(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.AuthCustomer(c)
}

func (ch *CustomerHandle) DeleteCustomer(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.RemoveAuthCustomer(c)
}

func (ch *CustomerHandle) IsCustomerAuthenticated(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.IsAuthCustomer(c)
}

func (ch *CustomerHandle) CreateOrder(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.OrderCreation(c)
}

func (ch *CustomerHandle) CreateMultipleOrder(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.MultipleOrderCreation(c)
}

func (ch *CustomerHandle) SetOrderStatus(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.SetStatus(c)
}

func (ch *CustomerHandle) FindOrderUpdates(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.FindUpdate(c)
}

func (ch *CustomerHandle) GetMyOrders(c *gin.Context) {
	os := Service.NewCustomerService(ch.db)
	os.GetMyOrders(c)
}
