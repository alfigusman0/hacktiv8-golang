package models

import "time"

type Order struct {
	OrderID      uint      `json:"order_id" gorm:"primary_key,auto_increment"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []*Item   `json:"items" gorm:"foreignKey:OrderID, references:OrderID"`
}

type GetAllOrderRequest struct {
	OrderID      uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []*Item   `json:"items" gorm:"foreignKey:OrderID, references:OrderID"`
}

type CreateOrderRequest struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	Items        []*Item `json:"items" binding:"required"`
}

type UpdateOrderRequest struct {
	CustomerName string  `json:"customer_name"`
	Items        []*Item `json:"items"`
}
