package Utils

import (
	"Day4-5_Codebase/config"
	"Day4-5_Codebase/models"
)

type InterfaceDB interface {
	IsPresent(id string, product *models.Product) (er error)
	DoCreate(product *models.Product) (er error)
	DoFind(id string, products *[]models.Product) (er error)
	DoUpdate(Product *models.Product, newProduct *models.Product) (er error)
	DoDelete(Product *models.Product) (er error)
	DoCreateR(retailer *models.Retailer) (er error)
	DoDeleteR(retailer *models.Retailer) (er error)
	IsPresentR(id string, retailer *models.Retailer) (er error)
	IsPresentRP(id string, products *models.Product) (er error)

	DoCreateC(Customer *models.Customer) (er error)
	DoDeleteC(Customer *models.Customer) (er error)
	IsPresentC(id string, Customer *models.Customer) (er error)
	IsPresentO(id string, Order *models.Order) (er error)
	IsPresentCU(id string, OrderUpdate *models.Order) (er error)
	DoCreateO(Order *models.Order) (er error)
	DoCreateOU(Order *models.Order) (er error)
	FindAllOrders(id string) (Orders *[]models.Order, er error)
	IsCustomerOrder(id string, Orders *[]models.Order) (er error)
	DoUpdateO(newOrder models.Order) (er error)
}

type DB struct {
}

func GetDB() InterfaceDB {
	return &DB{}
}

func (db *DB) IsPresent(id string, product *models.Product) (er error) {
	return config.DB.Where("id = ?", id).First(product).Error
}

func (db *DB) DoCreate(product *models.Product) (er error) {
	return config.DB.Create(product).Error
}

func (db *DB) DoFind(id string, products *[]models.Product) (er error) {
	return config.DB.Where("retailer_id = ?", id).Find(products).Error
}

func (db *DB) DoUpdate(Product *models.Product, newProduct *models.Product) (er error) {
	return config.DB.Model(Product).Updates(newProduct).Error
}

func (db *DB) DoDelete(Product *models.Product) (er error) {
	return config.DB.Delete(Product).Error
}

func (db *DB) DoCreateR(retailer *models.Retailer) (er error) {
	return config.DB.Create(retailer).Error
}

func (db *DB) DoDeleteR(retailer *models.Retailer) (er error) {
	return config.DB.Delete(retailer).Error
}

func (db *DB) IsPresentR(id string, retailer *models.Retailer) (er error) {
	return config.DB.Where("id = ?", id).First(retailer).Error
}

func (db *DB) IsPresentRP(id string, products *models.Product) (er error) {
	return config.DB.Where("retailer_id = ?", id).First(products).Error
}

func (db *DB) DoCreateC(Customer *models.Customer) (er error) {
	return config.DB.Create(Customer).Error
}

func (db *DB) DoDeleteC(Customer *models.Customer) (er error) {
	return config.DB.Delete(Customer).Error
}

func (db *DB) IsPresentC(id string, Customer *models.Customer) (er error) {
	return config.DB.Where("id = ?", id).First(Customer).Error
}

func (db *DB) IsPresentO(id string, Order *models.Order) (er error) {
	return config.DB.Where("id = ?", id).First(Order).Error
}

func (db *DB) IsPresentCU(id string, OrderUpdate *models.Order) (er error) {
	return config.DB.Where("customer_id = ?", id).First(OrderUpdate).Error
}

func (db *DB) DoCreateO(Order *models.Order) (er error) {
	return config.DB.Save(Order).Error
}

func (db *DB) DoCreateOU(Order *models.Order) (er error) {
	return config.DB.Save(Order).Error
}

func (db *DB) FindAllOrders(id string) (Orders *[]models.Order, er error) {
	var myProducts []models.Product
	config.DB.Where("retailer_id = ?", id).Find(&myProducts)

	var orders []models.Order
	for i := 0; i < len(myProducts); i++ {
		var tmpOrder []models.Order
		err := config.DB.Where("product_id = ?", myProducts[i].ID).Find(&tmpOrder).Error
		for j := 0; j < len(tmpOrder); j++ {
			orders = append(orders, tmpOrder[j])
		}
		if err != nil {
			var tmpOrders []models.Order
			return &tmpOrders, err
		}
	}
	return &orders, nil
}

func (db *DB) IsCustomerOrder(id string, Orders *[]models.Order) (er error) {
	return config.DB.Where("customer_id = ?", id).Find(Orders).Error
}

func (db *DB) DoUpdateO(newOrder models.Order) (er error) {
	return config.DB.Save(newOrder).Error
}
