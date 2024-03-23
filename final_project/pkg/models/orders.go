package models

import "time"

type Order struct {
	OrderID      uint      `json:"id" gorm:"column:id_order;primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name"`
	Items        []Item    `json:"items"`
	TmpTotal     float64   `json:"tmp_total" gorm:"column:tmp_total"`
	Potongan     float64   `json:"potongan" gorm:"column:potongan"`
	Total        float64   `json:"total" gorm:"column:total"`
	CreatedBy    User      `json:"created_by" gorm:"foreignKey:CreatedByID;references:UserID"`
	DateCreated  time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedBy    User      `json:"updated_by" gorm:"foreignKey:UpdatedByID;references"`
	DateUpdated  time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllOrderRequest struct {
	OrderID      uint    `json:"id"`
	CustomerName string  `json:"customer_name"`
	Items        *Item   `json:"items"`
	TmpTotal     float64 `json:"tmp_total"`
	Potongan     float64 `json:"potongan"`
	Total        float64 `json:"total"`
	CreatedBy    *User   `json:"created_by"`
}

type CreateOrderRequest struct {
	CustomerName string              `json:"customer_name" binding:"required"`
	Items        []CreateItemRequest `json:"items" binding:"required"`
	Potongan     float64             `json:"potongan" binding:"required"`
	CreatedByID  uint                `json:"created_by" binding:"required"`
	UpdatedByID  uint                `json:"updated_by" binding:"required"`
}

type UpdateOrderRequest struct {
	CustomerName string              `json:"customer_name"`
	Items        []UpdateItemRequest `json:"items"`
	Potongan     float64             `json:"potongan"`
	UpdatedByID  uint                `json:"updated_by" binding:"required"`
}
