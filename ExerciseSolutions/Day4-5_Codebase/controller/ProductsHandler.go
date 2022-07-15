package Controllers

import (
	"Day4-5_Codebase/models"
	"Day4-5_Codebase/mutex"
	"Day4-5_Codebase/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	} else {
		savedProduct, errNew := ph.ps.SaveProduct(&product)
		if errNew != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Product Saved": savedProduct})
		}
	}
}

func (ph ProductHandle) FindProduct(c *gin.Context) {
	var product models.Product
	_, err := ph.ps.FindMyProduct(c, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Product with id Found": product})
		return
	}
}

func (ph ProductHandle) GetProducts(c *gin.Context) {
	var products []models.Product
	listProduct, err := ph.ps.GetAllProducts(c, &products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Products": listProduct})
}

func (ph ProductHandle) UpdateProduct(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("product_id" + c.Param("id")); isAvailable == true {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "product id is already being updated, wait,you can try this on another id"})
		time.Sleep(2 * time.Second)
		return
	}
	defer mutex.Mutex.UnLock("product_id" + c.Param("id"))

	var productNew models.Product
	err := c.ShouldBindJSON(&productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedProduct, errNew := ph.ps.UpdateProduct(c, &productNew)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ProductNew": updatedProduct})
}

func (ph ProductHandle) DeleteProduct(c *gin.Context) {
	var product models.Product
	err := ph.ps.DeleteProduct(c, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if product.RetailerID != 0 {
		c.JSON(http.StatusOK, gin.H{"Product Deleted Successfully": product})
	}
}

func (ph ProductHandle) GetAllTransactions(c *gin.Context) {
	allOrders, err := ph.ps.GetAllTransactions(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"All transacted orders are": allOrders})
}

func (ph *ProductHandle) AuthRetailer(c *gin.Context) {
	var retailer models.Retailer
	err := c.ShouldBindJSON(&retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	newRetailer, errNew := ph.ps.AuthRetailer(c, &retailer)
	if errNew != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"Retailer added with": newRetailer})
}

func (ph *ProductHandle) RemoveAuthRetailer(c *gin.Context) {
	if isAvailable := mutex.Mutex.Lock("retailer_id" + c.Param("id")); isAvailable == false {
		c.JSON(http.StatusPreconditionFailed, gin.H{"Error": "retailer id is being deleted, wait,you can try this on another id"})
		return
	}
	if mutex.Mutex.Lock("retailer_id"+c.Param("id")) == true {
		time.Sleep(5 * time.Second)
		defer mutex.Mutex.UnLock("retailer_id" + c.Param("id"))
	}

	var retailer models.Retailer

	retailerRemoved, err := ph.ps.RemoveAuthRetailer(c, &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Retailer removed is": retailerRemoved})
}

func (ph *ProductHandle) IsRetailerAuthenticated(c *gin.Context) (retails *models.Retailer, er error) {
	var retailer models.Retailer
	resRetailer, err := ph.ps.IsAuthRetailer(c, &retailer)

	if err != nil {
		var rtls models.Retailer
		return &rtls, err
	} else {
		return resRetailer, nil
	}
}
