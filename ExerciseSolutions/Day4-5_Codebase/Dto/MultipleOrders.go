package Dto

import "errors"

type Order struct {
	ID         uint   `json:"id"`
	CustomerID int    `json:"customer_id" gorm:"foreign_key"`
	ProductID  []int  ` json:"product_id"`
	Quantity   []int  `json:"quantity"`
	Status     string `json:"status"`
}

var (
	ErrInvalidID  = errors.New("invalid ID")
	ErrCustomerID = errors.New("customer id is invalid")
	ErrProductID  = errors.New("product id is invalid")
	ErrQuantity   = errors.New("quantity value is invalid")
)

func (o *Order) OrderValidate() error {
	switch {
	case o.ID < 0:
		return ErrInvalidID
	case len(o.ProductID) == 0:
		return ErrProductID
	case o.CustomerID <= 0:
		return ErrCustomerID
	case len(o.Quantity) == 0:
		return ErrQuantity
	default:
		return nil
	}
}
