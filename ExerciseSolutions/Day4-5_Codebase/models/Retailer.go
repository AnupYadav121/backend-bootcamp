package models

import "errors"

type Retailer struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var (
	ErrNameNew     = errors.New("retailer name is invalid")
	ErrPasswordNew = errors.New("retailer password provided is invalid")
)

func (r *Retailer) RetailerValidate() error {
	switch {
	case r.ID < 0:
		return ErrCustomerID
	case r.Name == "":
		return ErrNameNew
	case r.Password == "":
		return ErrPasswordNew
	default:
		return nil
	}
}
