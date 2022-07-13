package Models

type Customer struct {
	ID uint `json:"id" gorm:"primary_key"`
}

func (c *Customer) CustomerValidate() error {
	switch {
	case c.ID < 0:
		return ErrCustomerID
	default:
		return nil
	}
}
