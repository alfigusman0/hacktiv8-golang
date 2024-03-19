package models

import "time"

type Order struct {
	OrderID      uint      `json:"id" gorm:"column:order_id;primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"column:ordered_at"`
	Items        *Item     `json:"items" gorm:"foreignKey:OrderID;references:OrderID"`
}

type GetAllOrderRequest struct {
	OrderID      uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        *Item     `json:"items"`
}

type CreateOrderRequest struct {
	CustomerName string              `json:"customer_name" binding:"required"`
	Items        []CreateItemRequest `json:"items" binding:"required"`
}

type UpdateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	Items        []Item `json:"items"`
}
