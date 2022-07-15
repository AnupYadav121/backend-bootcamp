package models

import "errors"

type Customer struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var (
	ErrName     = errors.New("customer name is invalid")
	ErrPassword = errors.New("password provided is invalid")
)

func (c *Customer) CustomerValidate() error {
	switch {
	case c.ID < 0:
		return ErrCustomerID
	case c.Name == "":
		return ErrName
	case c.Password == "":
		return ErrPassword
	default:
		return nil
	}
}
