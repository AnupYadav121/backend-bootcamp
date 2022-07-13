package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	Utils "july8Files/DB_Utils"
	"july8Files/Service"
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
	db Utils.InterfaceDB
}

func NewProduct(db Utils.InterfaceDB) *ProductHandle {
	return &ProductHandle{db}
}

func (ph ProductHandle) CreateProduct(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.SaveProduct(c)
}

func (ph ProductHandle) FindProduct(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.FindMyProduct(c)
}

func (ph ProductHandle) GetProducts(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.GetAllProducts(c)
}

func (ph ProductHandle) UpdateProduct(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.UpdateProduct(c)
}

func (ph ProductHandle) DeleteProduct(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.DeleteProduct(c)
}

func (ph ProductHandle) GetAllTransactions(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.GetAllTransactions(c)
}

func (ph *ProductHandle) AuthRetailer(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.AuthRetailer(c)
}

func (ph *ProductHandle) RemoveAuthRetailer(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.RemoveAuthRetailer(c)
}

func (ph *ProductHandle) IsRetailerAuthenticated(c *gin.Context) {
	ps := Service.NewProductService(ph.db)
	ps.IsAuthRetailer(c)
}
