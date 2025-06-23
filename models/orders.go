package models

import "time"

type Order struct {
	OrderId     string    `json:"order_id" gorm:"primaryKey"`
	CustomerId  string    `json:"customer_id"`
	OrderedAt   time.Time `json:"ordered_at"`
	TotalAmount float64   `json:"total_amount"`
	NoOfItems   int       `json:"no_of_items"`
}

type OrderItem struct {
	OrderItemId string `json:"order_item_id" gorm:"primaryKey"`
	OrderId     string `json:"order_id"`
	ArticleId   string `json:"article_id"`
	Quantity    string `json:"quantity"`
}
