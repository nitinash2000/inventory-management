package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderId     string    `json:"order_id" gorm:"primaryKey"`
	CustomerId  string    `json:"customer_id"`
	OrderedAt   time.Time `json:"ordered_at"`
	TotalAmount float64   `json:"total_amount"`
	NoOfItems   int       `json:"no_of_items"`
}

func (o *Order) BeforeSave(tx *gorm.DB) error {
	if o.CustomerId == "" {
		return errors.New("customer id is required")
	}

	return nil
}

type OrderItem struct {
	OrderItemId string `json:"order_item_id" gorm:"primaryKey"`
	OrderId     string `json:"order_id"`
	ArticleId   string `json:"article_id"`
	Quantity    int    `json:"quantity"`
}

func (oi *OrderItem) BeforeSave(tx *gorm.DB) error {
	if oi.ArticleId == "" {
		return errors.New("article id is required")
	}

	return nil
}
