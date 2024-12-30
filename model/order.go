package model

type Order struct {
	orderID uint64 		   `json:"order_id"`
	customerID uuid.UUID   `json:"customer_id"`
	lineItems []lineItem   `json:"line_items"`
	CreatedAt *time.Time   `json:"created_at`
	ShippedAt *time.Time   `json:"shipped_at"`
	CompletedAt *time.Time `json:"complete d_at"`
}

type LineItem struct {
	itemID uuid.UUID `json:"item_id"`
	Quantity uint	 `json:"quantity"`
	Price uint		 `json:"price"`
}

