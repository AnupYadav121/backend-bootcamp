package Models

type Transaction struct {
	CustomerId  uint   `json:"customer_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
