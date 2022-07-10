package Models

type Order struct {
	ID         uint `json:"id"`
	CustomerID int  `json:"customer_id"`
	ProductID  int  ` json:"product_id"`
	Quantity   int  `json:"quantity"`
}

type OrderUpdated struct {
	ID         uint   `json:"id"`
	CustomerID int    `json:"customer_id"`
	ProductID  int    ` json:"product_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}

type Customer struct {
	ID uint `json:"id"`
}
