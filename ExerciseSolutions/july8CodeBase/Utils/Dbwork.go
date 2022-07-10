package Utils

import (
	"july8Files/Config"
	"july8Files/Models"
)

func IsPresent(id string, Product *Models.Product) error {
	return Config.DB.Where("id = ?", id).First(Product).Error
}

func DoCreate(Product *Models.Product) {
	Config.DB.Create(Product)
}

func DoFind(Product *[]Models.Product) {
	Config.DB.Find(Product)
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

func IsPresentO(id string, Order *Models.OrderUpdated) error {
	return Config.DB.Where("id = ?", id).First(Order).Error
}

func DoCreateO(Order *Models.Order) {
	Config.DB.Create(Order)
}
