package Utils

import (
	"july8Files/Config"
	"july8Files/Models"
)

type InterfaceDB interface {
	IsPresent(id string, product *Models.Product) error
	DoCreate(product *Models.Product)
	DoFind(products *[]Models.Product)
	DoUpdate(Product *Models.Product, newProduct Models.UpdatedProduct)
	DoDelete(Product *Models.Product) error
	DoCreateC(Customer *Models.Customer)
	DoDeleteC(Customer *Models.Customer) error
	IsPresentC(id string, Customer *Models.Customer) error
	IsPresentO(id string, Order *Models.Order) error
	IsPresentOU(id string, OrderUpdate *Models.OrderUpdated) error
	IsPresentCU(id string, OrderUpdate *Models.OrderUpdated) error
	DoCreateO(Order *Models.Order)
	DoCreateOU(Order *Models.OrderUpdated)
	FindAllOrders(products *[]Models.OrderUpdated) error
	IsCustomerOrder(id string, Orders *[]Models.Order) error
}

type DB struct {
}

func GetDB() InterfaceDB {
	return &DB{}
}

func (db *DB) IsPresent(id string, product *Models.Product) error {
	return Config.DB.Where("id = ?", id).First(product).Error
}

func (db *DB) DoCreate(product *Models.Product) {
	Config.DB.Create(product)
}

func (db *DB) DoFind(products *[]Models.Product) {
	Config.DB.Find(products)
}

func (db *DB) DoUpdate(Product *Models.Product, newProduct Models.UpdatedProduct) {
	Config.DB.Model(Product).Updates(newProduct)
}

func (db *DB) DoDelete(Product *Models.Product) error {
	return Config.DB.Delete(Product).Error
}

func (db *DB) DoCreateC(Customer *Models.Customer) {
	Config.DB.Create(Customer)
}

func (db *DB) DoDeleteC(Customer *Models.Customer) error {
	return Config.DB.Delete(Customer).Error
}

func (db *DB) IsPresentC(id string, Customer *Models.Customer) error {
	return Config.DB.Where("id = ?", id).First(Customer).Error
}

func (db *DB) IsPresentO(id string, Order *Models.Order) error {
	return Config.DB.Where("id = ?", id).First(Order).Error
}

func (db *DB) IsPresentOU(id string, OrderUpdate *Models.OrderUpdated) error {
	return Config.DB.Where("id = ?", id).First(OrderUpdate).Error
}

func (db *DB) IsPresentCU(id string, OrderUpdate *Models.OrderUpdated) error {
	return Config.DB.Where("customer_id = ?", id).First(OrderUpdate).Error
}

func (db *DB) DoCreateO(Order *Models.Order) {
	Config.DB.Create(Order)
}

func (db *DB) DoCreateOU(Order *Models.OrderUpdated) {
	Config.DB.Create(Order)
}

func (db *DB) FindAllOrders(products *[]Models.OrderUpdated) error {
	return Config.DB.Find(products).Error
}

func (db *DB) IsCustomerOrder(id string, Orders *[]Models.Order) error {
	return Config.DB.Where("customer_id = ?", id).First(Orders).Error
}
