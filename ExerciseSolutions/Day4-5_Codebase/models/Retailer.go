package models

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
