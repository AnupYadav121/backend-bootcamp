package service

import (
	Utils "Day4-5_Codebase/db_utils"
	"Day4-5_Codebase/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type ProductServiceInterface interface {
	SaveProduct(c *gin.Context, product *models.Product)
	FindMyProduct(c *gin.Context, product *models.Product)
	GetAllProducts(c *gin.Context, products *[]models.Product)
	UpdateProduct(c *gin.Context, productNew *models.Product)
	DeleteProduct(c *gin.Context, product *models.Product)
	GetAllTransactions(c *gin.Context, allOrders *[]models.Order)
	AuthRetailer(c *gin.Context, retailer *models.Retailer)
	RemoveAuthRetailer(c *gin.Context, retailer *models.Retailer)
	IsAuthRetailer(c *gin.Context, retailer models.Retailer)
}

var prdSrv ProductService

type ProductService struct {
	db Utils.InterfaceDB
}

func NewProductService(db Utils.InterfaceDB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) SaveProduct(c *gin.Context, product *models.Product) (p *models.Product, err error) {

	nwErr := product.ProductValidate()
	if nwErr != nil {
		var tmpProduct models.Product
		return &tmpProduct, nwErr
	}
	fmt.Println(ps.db)
	var retailer models.Retailer
	errNew := ps.db.IsPresentR(strconv.Itoa(int(product.RetailerID)), &retailer)
	if errNew != nil {
		var tmpProduct models.Product
		return &tmpProduct, errNew
	}
	ps.db.DoCreate(product)
	return product, nil
}

func (ps *ProductService) FindMyProduct(c *gin.Context, product *models.Product) {
	err := ps.db.IsPresent(c.Param("id"), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
}

func (ps *ProductService) GetAllProducts(c *gin.Context, products *[]models.Product) {
	ps.db.DoFind(products)
}

func (ps *ProductService) UpdateProduct(c *gin.Context, productNew *models.Product) {
	var ProductExist models.Product
	errNew := ps.db.IsPresent(c.Param("id"), &ProductExist)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}

	if ProductExist.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not find your order with id to be updated"})
		return
	}

	ps.db.DoUpdate(&ProductExist, productNew)
}

func (ps *ProductService) DeleteProduct(c *gin.Context, product *models.Product) {
	err := ps.db.IsPresent(c.Param("id"), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	errNew := ps.db.DoDelete(product)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
}

func (ps *ProductService) GetAllTransactions(c *gin.Context, allOrders *[]models.Order) {
	errNew := ps.db.FindAllOrders(allOrders)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
}

func (ps *ProductService) AuthRetailer(c *gin.Context, retailer *models.Retailer) {
	errNw := retailer.RetailerValidate()
	if errNw != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNw.Error()})
		return
	}
	ps.db.DoCreateR(retailer)
}

func (ps *ProductService) RemoveAuthRetailer(c *gin.Context, retailer *models.Retailer) {

	err := ps.db.IsPresentR(c.Param("id"), retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var products models.Product
	newErr := ps.db.IsPresentRP(c.Param("id"), &products)
	if newErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can not remove authentication, One product is process for this retailer id"})
		return
	}

	errNew := ps.db.DoDeleteR(retailer)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Retailer Deleted": retailer})
}

func (ps *ProductService) IsAuthRetailer(c *gin.Context, retailer models.Retailer) {
	err := ps.db.IsPresentR(c.Param("id"), &retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Retailer is not Authenticated"})
		return
	}
}
