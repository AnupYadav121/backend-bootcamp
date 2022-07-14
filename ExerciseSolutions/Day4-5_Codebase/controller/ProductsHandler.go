package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/models"
	"july8Files/mutex"
	"july8Files/service"
	"net/http"
	"time"
)

type ProductsInterface interface {
	CreateProduct(c *gin.Context)
	FindProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetAllTransactions(c *gin.Context)
	AuthRetailer(c *gin.Context)
	RemoveAuthRetailer(c *gin.Context)
	IsRetailerAuthenticated(c *gin.Context)
}

type ProductHandle struct {
	ps service.ProductServiceInterface
}

func NewProduct(ps service.ProductServiceInterface) *ProductHandle {
	return &ProductHandle{ps}
}

func (ph ProductHandle) CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ph.ps.SaveProduct(c, &product)
	c.JSON(http.StatusOK, gin.H{"Product Saved": product})
}

func (ph ProductHandle) FindProduct(c *gin.Context) {
	var product models.Product
	ph.ps.FindMyProduct(c, &product)
	c.JSON(http.StatusOK, gin.H{"Product with id Found": product})
}

func (ph ProductHandle) GetProducts(c *gin.Context) {
	var products []models.Product
	ph.ps.GetAllProducts(c, &products)
	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func (ph ProductHandle) UpdateProduct(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("product_id" + c.Param("id")); isAvailable == false {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "product id is already being updated, wait"})
		return
	}
	time.Sleep(3 * time.Second)
	defer mutex.Mutex.UnLock("product_id" + c.Param("id"))

	var productNew models.Product
	err := c.ShouldBindJSON(&productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ph.ps.UpdateProduct(c, &productNew)
	c.JSON(http.StatusOK, gin.H{"ProductNew": productNew})
}

func (ph ProductHandle) DeleteProduct(c *gin.Context) {
	var product models.Product
	ph.ps.DeleteProduct(c, &product)
	c.JSON(http.StatusOK, gin.H{"Product Deleted Successfully": product})
}

func (ph ProductHandle) GetAllTransactions(c *gin.Context) {
	var allOrders []models.Order
	ph.ps.GetAllTransactions(c, &allOrders)
	c.JSON(http.StatusOK, gin.H{"All Order updates": allOrders})
}

func (ph *ProductHandle) AuthRetailer(c *gin.Context) {
	var retailer models.Retailer
	err := c.ShouldBindJSON(&retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ph.ps.AuthRetailer(c, &retailer)
	c.JSON(http.StatusOK, gin.H{"Retailer": retailer})
}

func (ph *ProductHandle) RemoveAuthRetailer(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("retailer_id" + c.Param("id")); isAvailable == false {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "retailer id is being deleted, wait"})
		return
	}
	time.Sleep(2 * time.Second)
	defer mutex.Mutex.UnLock("retailer_id" + c.Param("id"))

	var retailer models.Retailer

	ph.ps.RemoveAuthRetailer(c, &retailer)
}

func (ph *ProductHandle) IsRetailerAuthenticated(c *gin.Context) {
	var retailer models.Retailer
	ph.ps.IsAuthRetailer(c, retailer)
	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Retailer is Authenticated"})
}
