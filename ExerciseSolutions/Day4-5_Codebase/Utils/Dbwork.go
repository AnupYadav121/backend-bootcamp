package Utils

import (
	"july8Files/Config"
	"july8Files/Models"
)

func IsPresent(id string, product *Models.Product) error {
	return Config.DB.Where("id = ?", id).First(product).Error
}

func DoCreate(product *Models.Product) {
	Config.DB.Create(product)
}

func DoFind(products *[]Models.Product) {
	Config.DB.Find(products)
}

func DoUpdate(Product *Models.Product, newProduct Models.UpdatedProduct) {
	Config.DB.Model(Product).Updates(newProduct)
}

func DoDelete(Product *Models.Product) error {
	return Config.DB.Delete(Product).Error
}

func DoCreateC(Customer *Models.Customer) {
	Config.DB.Create(Customer)
}

func DoDeleteC(Customer *Models.Customer) error {
	return Config.DB.Delete(Customer).Error
}

func IsPresentC(id string, Customer *Models.Customer) error {
	return Config.DB.Where("id = ?", id).First(Customer).Error
}

func IsPresentO(id string, Order *Models.Order) error {
	return Config.DB.Where("id = ?", id).First(Order).Error
}

func IsPresentOU(id string, OrderUpdate *Models.OrderUpdated) error {
	return Config.DB.Where("id = ?", id).First(OrderUpdate).Error
}

func IsPresentCU(id string, OrderUpdate *Models.OrderUpdated) error {
	return Config.DB.Where("customer_id = ?", id).First(OrderUpdate).Error
}

func DoCreateO(Order *Models.Order) {
	Config.DB.Create(Order)
}

func DoCreateOU(Order *Models.OrderUpdated) {
	Config.DB.Create(Order)
}

func FindAllOrders(products *[]Models.OrderUpdated) error {
	return Config.DB.Find(products).Error
}

func IsCustomerOrder(id string, Orders *[]Models.Order) error {
	return Config.DB.Where("customer_id = ?", id).First(Orders).Error
}
