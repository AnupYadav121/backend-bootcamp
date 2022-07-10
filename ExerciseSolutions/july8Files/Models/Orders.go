package Models

type Order struct {
	ID         uint `json:"id"`
	CustomerID uint `json:"customer_id"`
	ProductID  uint ` json:"product_id"`
	Quantity   int  `json:"quantity"`
}

type OrderUpdated struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	ProductID  uint   ` json:"product_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}

type Customer struct {
	CustomerId int     `json:"customer_id"`
	Orders     []Order `json:"orders"`
}
