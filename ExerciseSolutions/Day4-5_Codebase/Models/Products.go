package Models

import "errors"

type Product struct {
	ID          uint   `json:"id"`
	RetailerID  uint   `json:"retailerID" gorm:"foreign_key"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

var (
	ErrProductName = errors.New("customer id is invalid")
	ErrPrice       = errors.New("product id is invalid")
	ErrRetailerID  = errors.New("retailer id is invalid")
)

func (p *Product) ProductValidate() error {
	switch {
	case p.ID < 0:
		return ErrInvalidID
	case p.ProductName == "":
		return ErrProductName
	case p.Price <= 0:
		return ErrPrice
	case p.Quantity <= 0:
		return ErrQuantity
	case p.RetailerID <= 0:
		return ErrRetailerID

	default:
		return nil
	}
}

type Retailer struct {
	ID uint `json:"id" gorm:"primary_key"`
}

func (r *Retailer) RetailerValidate() error {
	switch {
	case r.ID < 0:
		return ErrRetailerID
	default:
		return nil
	}
}
