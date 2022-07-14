package Utils

import (
	"Day4-5_Codebase/config"
	"Day4-5_Codebase/models"
)

type InterfaceDB interface {
	IsPresent(id string, product *models.Product) error
	DoCreate(product *models.Product)
	DoFind(products *[]models.Product)
	DoUpdate(Product *models.Product, newProduct *models.Product)
	DoDelete(Product *models.Product) error
	DoCreateR(retailer *models.Retailer)
	DoDeleteR(retailer *models.Retailer) error
	IsPresentR(id string, retailer *models.Retailer) error
	IsPresentRP(id string, products *models.Product) error

	DoCreateC(Customer *models.Customer)
	DoDeleteC(Customer *models.Customer) error
	IsPresentC(id string, Customer *models.Customer) error
	IsPresentO(id string, Order *models.Order) error
	IsPresentCU(id string, OrderUpdate *models.Order) error
	DoCreateO(Order *models.Order)
	DoCreateOU(Order *models.Order)
	FindAllOrders(products *[]models.Order) error
	IsCustomerOrder(id string, Orders *[]models.Order) error
	DoUpdateO(newOrder models.Order)
}

type DB struct {
}

func GetDB() InterfaceDB {
	return &DB{}
}

func (db *DB) IsPresent(id string, product *models.Product) error {
	return config.DB.Where("id = ?", id).First(product).Error
}

func (db *DB) DoCreate(product *models.Product) {
	config.DB.Create(product)
}

func (db *DB) DoFind(products *[]models.Product) {
	config.DB.Find(products)
}

func (db *DB) DoUpdate(Product *models.Product, newProduct *models.Product) {
	config.DB.Model(Product).Updates(newProduct)
}

func (db *DB) DoDelete(Product *models.Product) error {
	return config.DB.Delete(Product).Error
}

func (db *DB) DoCreateR(retailer *models.Retailer) {
	config.DB.Create(retailer)
}

func (db *DB) DoDeleteR(retailer *models.Retailer) error {
	return config.DB.Delete(retailer).Error
}

func (db *DB) IsPresentR(id string, retailer *models.Retailer) error {
	return config.DB.Where("id = ?", id).First(retailer).Error
}

func (db *DB) IsPresentRP(id string, products *models.Product) error {
	return config.DB.Where("retailer_id = ?", id).First(products).Error
}

func (db *DB) DoCreateC(Customer *models.Customer) {
	config.DB.Create(Customer)
}

func (db *DB) DoDeleteC(Customer *models.Customer) error {
	return config.DB.Delete(Customer).Error
}

func (db *DB) IsPresentC(id string, Customer *models.Customer) error {
	return config.DB.Where("id = ?", id).First(Customer).Error
}

func (db *DB) IsPresentO(id string, Order *models.Order) error {
	return config.DB.Where("id = ?", id).First(Order).Error
}

func (db *DB) IsPresentCU(id string, OrderUpdate *models.Order) error {
	return config.DB.Where("customer_id = ?", id).First(OrderUpdate).Error
}

func (db *DB) DoCreateO(Order *models.Order) {
	config.DB.Save(Order)
}

func (db *DB) DoCreateOU(Order *models.Order) {
	config.DB.Save(Order)
}

func (db *DB) FindAllOrders(products *[]models.Order) error {
	return config.DB.Find(products).Error
}

func (db *DB) IsCustomerOrder(id string, Orders *[]models.Order) error {
	return config.DB.Where("customer_id = ?", id).Find(Orders).Error
}

func (db *DB) DoUpdateO(newOrder models.Order) {
	config.DB.Save(newOrder)
}
