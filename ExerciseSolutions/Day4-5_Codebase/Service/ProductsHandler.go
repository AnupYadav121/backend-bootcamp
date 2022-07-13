package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/DB_Utils"
	"july8Files/Models"
	"net/http"
	"strconv"
)

type ProductsInterface interface {
	CreateProduct(c *gin.Context)
	FindProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetAllTransactions(c *gin.Context)
	SetOrderStatus(c *gin.Context)
	AuthRetailer(c *gin.Context)
	RemoveAuthRetailer(c *gin.Context)
	IsRetailerAuthenticated(c *gin.Context)
}

type ProductHandle struct {
	db Utils.InterfaceDB
}

func NewProduct(db Utils.InterfaceDB) *ProductHandle {
	return &ProductHandle{db}
}

func (ph ProductHandle) CreateProduct(c *gin.Context) {
	var product Models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	nwErr := product.ProductValidate()
	if nwErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": nwErr.Error()})
		return
	}

	var retailer Models.Retailer
	errNew := ph.db.IsPresentR(strconv.Itoa(int(product.RetailerID)), &retailer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Sorry,Retailer is not Authenticated"})
		return
	}

	ph.db.DoCreate(&product)
	c.JSON(http.StatusOK, gin.H{"Product Saved": product})
}

func (ph ProductHandle) FindProduct(c *gin.Context) {
	var product Models.Product
	err := ph.db.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product with id Found": product})
}

func (ph ProductHandle) GetProducts(c *gin.Context) {
	var products []Models.Product
	ph.db.DoFind(&products)
	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func (ph ProductHandle) UpdateProduct(c *gin.Context) {
	var productNew Models.Product
	err := c.ShouldBindJSON(&productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var ProductExist Models.Product
	errNew := ph.db.IsPresent(c.Param("id"), &ProductExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}

	if ProductExist.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not find your order with id to be updated"})
		return
	}

	ph.db.DoUpdate(&ProductExist, productNew)
	c.JSON(http.StatusOK, gin.H{"ProductNew": productNew})
}

func (ph ProductHandle) DeleteProduct(c *gin.Context) {
	var product Models.Product
	err := ph.db.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	errNew := ph.db.DoDelete(&product)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product Deleted Successfully": &product})
}

func (ph ProductHandle) GetAllTransactions(c *gin.Context) {
	var allOrders []Models.Order
	errNew := ph.db.FindAllOrders(&allOrders)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All Order updates": allOrders})
}

func (ph ProductHandle) SetOrderStatus(c *gin.Context) {
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
	errNew := ph.db.IsPresentO(c.Param("id"), &orderExist)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given Order id to be updated do not exist"})
		return
	}
	orderUpdated := Models.Order{ID: orderExist.ID, CustomerID: orderExist.CustomerID, ProductID: orderExist.ProductID, Quantity: orderExist.Quantity, Status: orderStatus.Status}

	ph.db.DoCreateOU(&orderUpdated)
	c.JSON(http.StatusOK, gin.H{"Order": orderUpdated})
}

func (ph *ProductHandle) AuthRetailer(c *gin.Context) {
	var retailer Models.Retailer
	err := c.ShouldBindJSON(&retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	errNw := retailer.RetailerValidate()
	if errNw != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNw.Error()})
		return
	}

	ph.db.DoCreateR(&retailer)
	c.JSON(http.StatusOK, gin.H{"Retailer": retailer})
}

func (ph *ProductHandle) RemoveAuthRetailer(c *gin.Context) {
	var retailer Models.Retailer
	err := ph.db.IsPresentR(c.Param("id"), &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var products Models.Product
	newErr := ph.db.IsPresentRP(c.Param("id"), &products)
	if newErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can not remove authentication, One product is process for this retailer id"})
		return
	}

	errNew := ph.db.DoDeleteR(&retailer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Retailer Deleted": &retailer})
}

func (ph *ProductHandle) IsRetailerAuthenticated(c *gin.Context) {
	var retailer Models.Retailer
	err := ph.db.IsPresentR(c.Param("id"), &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Retailer is not Authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Retailer is Authenticated"})
}
