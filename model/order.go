package model

type Order struct {
	ID          int     `json:"id"`
	ProductIDs  string  `json:"product_ids"`
	OrderAmount float32 `json:"order_amount"`
}

func NewOrder() *Order {
	return &Order{}
}
