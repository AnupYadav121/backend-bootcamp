package Service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	Utils "july8Files/DB_Utils"
	"july8Files/Models"
	"net/http"
	"strconv"
)

type ProductServiceInterface interface {
	SaveProduct(c *gin.Context)
	FindMyProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetAllTransactions(c *gin.Context)
	AuthRetailer(c *gin.Context)
	RemoveAuthRetailer(c *gin.Context)
	IsAuthRetailer(c *gin.Context)
}

type ProductService struct {
	db Utils.InterfaceDB
}

func NewProductService(db Utils.InterfaceDB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) SaveProduct(c *gin.Context) {
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
	errNew := ps.db.IsPresentR(strconv.Itoa(int(product.RetailerID)), &retailer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Sorry,Retailer is not Authenticated"})
		return
	}

	ps.db.DoCreate(&product)
	c.JSON(http.StatusOK, gin.H{"Product Saved": product})
}

func (ps *ProductService) FindMyProduct(c *gin.Context) {
	var product Models.Product
	err := ps.db.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product with id Found": product})
}

func (ps *ProductService) GetAllProducts(c *gin.Context) {
	var products []Models.Product
	ps.db.DoFind(&products)
	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func (ps *ProductService) UpdateProduct(c *gin.Context) {
	var productNew Models.Product
	err := c.ShouldBindJSON(&productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var ProductExist Models.Product
	errNew := ps.db.IsPresent(c.Param("id"), &ProductExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}

	if ProductExist.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not find your order with id to be updated"})
		return
	}

	ps.db.DoUpdate(&ProductExist, productNew)
	c.JSON(http.StatusOK, gin.H{"ProductNew": productNew})
}

func (ps *ProductService) DeleteProduct(c *gin.Context) {
	var product Models.Product
	err := ps.db.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	errNew := ps.db.DoDelete(&product)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product Deleted Successfully": &product})
}

func (ps *ProductService) GetAllTransactions(c *gin.Context) {
	var allOrders []Models.Order
	errNew := ps.db.FindAllOrders(&allOrders)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All Order updates": allOrders})
}

func (ps *ProductService) AuthRetailer(c *gin.Context) {
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

	ps.db.DoCreateR(&retailer)
	c.JSON(http.StatusOK, gin.H{"Retailer": retailer})
}

func (ps *ProductService) RemoveAuthRetailer(c *gin.Context) {
	var retailer Models.Retailer
	err := ps.db.IsPresentR(c.Param("id"), &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var products Models.Product
	newErr := ps.db.IsPresentRP(c.Param("id"), &products)
	if newErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can not remove authentication, One product is process for this retailer id"})
		return
	}

	errNew := ps.db.DoDeleteR(&retailer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Retailer Deleted": &retailer})
}

func (ps *ProductService) IsAuthRetailer(c *gin.Context) {
	var retailer Models.Retailer
	err := ps.db.IsPresentR(c.Param("id"), &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Retailer is not Authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Great !, Retailer is Authenticated"})
}
