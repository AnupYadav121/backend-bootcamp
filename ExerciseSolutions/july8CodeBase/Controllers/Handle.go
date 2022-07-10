package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"july8Files/Models"
	"july8Files/Utils"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var product Models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Utils.DoCreate(&product)
	c.JSON(http.StatusOK, gin.H{"Products": product})
}

func FindProduct(c *gin.Context) {
	var product Models.Product
	err := Utils.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Products": product})
}

func GetProducts(c *gin.Context) {
	var products []Models.Product
	Utils.DoFind(&products)
	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func UpdateProduct(c *gin.Context) {
	var productNew Models.UpdatedProduct
	err := c.ShouldBindJSON(&productNew).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err()})
		return
	}

	var ProductExist Models.Product
	errNew := Utils.IsPresent(c.Param("id"), &ProductExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errNew.Error()})
	}

	Utils.DoUpdate(&ProductExist, productNew)
	c.JSON(http.StatusOK, gin.H{"ProductNew": productNew})
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	err := Utils.IsPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	errNew := Utils.DoDelete(&product)
	if errNew != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Products": &product})
}

func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	Utils.DoCreateC(&customer)
	c.JSON(http.StatusOK, gin.H{"Customer": customer})
}

func DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	err := Utils.IsPresentC(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	errNew := Utils.DoDeleteC(&customer)
	if errNew != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": errNew.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer Deleted": &customer})
}

func CreateOrder(c *gin.Context) {
	var newOrder Models.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var product Models.Product
	newErr := Utils.IsPresent(strconv.Itoa(int(newOrder.ProductID)), &product)
	if newErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": newErr.Error()})
		return
	}

	Utils.DoCreateO(&newOrder)
	c.JSON(http.StatusOK, gin.H{"Order": newOrder})
}

func FindOrderUpdates(c *gin.Context) {
	var orderExist Models.OrderUpdated
	err := Utils.IsPresentO(c.Param("id"), &orderExist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Order": orderExist})
}
