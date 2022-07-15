package service

import (
	Utils "Day4-5_Codebase/db_utils"
	"Day4-5_Codebase/models"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type ProductServiceInterface interface {
	SaveProduct(product *models.Product) (p *models.Product, err error)
	FindMyProduct(c *gin.Context, product *models.Product) (p *models.Product, err error)
	GetAllProducts(c *gin.Context, products *[]models.Product) (prods *[]models.Product, er error)
	UpdateProduct(c *gin.Context, productNew *models.Product) (productUpdated *models.Product, er error)
	DeleteProduct(c *gin.Context, product *models.Product) (er error)
	GetAllTransactions(c *gin.Context) (orders *[]models.Order, er error)
	AuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error)
	RemoveAuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error)
	IsAuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error)
}

type ProductService struct {
	db Utils.InterfaceDB
}

func NewProductService(db Utils.InterfaceDB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) SaveProduct(product *models.Product) (p *models.Product, err error) {
	nwErr := product.ProductValidate()
	if nwErr != nil {
		var product models.Product
		return &product, errors.New("provided product body is invalid")
	}

	var retailer models.Retailer
	errNew := ps.db.IsPresentR(strconv.Itoa(product.RetailerID), &retailer)
	if errNew != nil {
		var product models.Product
		return &product, errors.New("retailer with id is not authenticated")
	}

	erNew := ps.db.DoCreate(product)
	if erNew != nil {
		var product models.Product
		return &product, erNew
	}
	return product, nil
}

func (ps *ProductService) FindMyProduct(c *gin.Context, product *models.Product) (p *models.Product, err error) {
	errNew := ps.db.IsPresent(c.Param("id"), product)
	if errNew != nil {
		var tmpProduct models.Product
		return &tmpProduct, errors.New("product with given id could not be found")
	}
	return product, nil
}

func (ps *ProductService) GetAllProducts(c *gin.Context, products *[]models.Product) (prods *[]models.Product, er error) {
	erNew := ps.db.DoFind(c.Param("retailerID"), products)
	if erNew != nil {
		var product []models.Product
		return &product, erNew
	}
	return products, nil
}

func (ps *ProductService) UpdateProduct(c *gin.Context, productNew *models.Product) (productUpdated *models.Product, er error) {
	var ProductExist models.Product
	errNew := ps.db.IsPresent(c.Param("id"), &ProductExist)
	if errNew != nil {
		var product models.Product
		return &product, errors.New("product with given id does not exist")
	}

	if ProductExist.Price == 0 {
		var product models.Product
		return &product, errors.New("could not find your order with id to be updated")
	}

	erNew := ps.db.DoUpdate(&ProductExist, productNew)
	if erNew != nil {
		var product models.Product
		return &product, erNew
	}
	return productNew, nil
}

func (ps *ProductService) DeleteProduct(c *gin.Context, product *models.Product) (er error) {
	err := ps.db.IsPresent(c.Param("id"), product)
	if err != nil {
		return errors.New("product with given id not found")
	}
	errNew := ps.db.DoDelete(product)
	if errNew != nil {
		return errNew
	}
	return nil
}

func (ps *ProductService) GetAllTransactions(c *gin.Context) (orders *[]models.Order, er error) {
	orders, errNew := ps.db.FindAllOrders(c.Param("retailerID"))
	if errNew != nil {
		var ords []models.Order
		return &ords, errNew
	}
	return orders, nil
}

func (ps *ProductService) AuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error) {
	errNw := retailer.RetailerValidate()
	if errNw != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errNw.Error()})
		var tmp models.Retailer
		return &tmp, errors.New("provided retailer body is invalid")
	}
	erNew := ps.db.DoCreateR(retailer)
	if erNew != nil {
		var rtl models.Retailer
		return &rtl, erNew
	}
	return retailer, nil
}

func (ps *ProductService) RemoveAuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error) {
	err := ps.db.IsPresentR(c.Param("retailerID"), retailer)
	if err != nil {
		var tmpRetailer models.Retailer
		return &tmpRetailer, errors.New("retailer with id not found")
	}

	var products models.Product
	newErr := ps.db.IsPresentRP(c.Param("retailerID"), &products)
	if newErr == nil {
		var tmpRetailer models.Retailer
		return &tmpRetailer, errors.New("retailer is associated with a product , can not delete it")
	}

	errNew := ps.db.DoDeleteR(retailer)
	if errNew != nil {
		var tmpRetailer models.Retailer
		return &tmpRetailer, errNew
	}
	return retailer, nil
}

func (ps *ProductService) IsAuthRetailer(c *gin.Context, retailer *models.Retailer) (r *models.Retailer, er error) {
	err := ps.db.IsPresentR(c.Param("retailerID"), retailer)
	if err != nil {
		var tmpRetailer models.Retailer
		return &tmpRetailer, err
	}
	return retailer, nil
}
