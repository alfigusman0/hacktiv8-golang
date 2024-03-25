package models

import "time"

type Order struct {
	OrderID      uint      `json:"id" gorm:"column:id_order;primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name"`
	Items        []Item    `json:"items"`
	TmpTotal     int       `json:"tmp_total" gorm:"column:tmp_total"`
	Potongan     int       `json:"potongan" gorm:"column:potongan"`
	Total        int       `json:"total" gorm:"column:total"`
	CreatedByID  uint      `json:"created_by_id" gorm:"column:created_by"`   // Foreign key ID
	CreatedBy    User      `json:"created_by" gorm:"foreignKey:CreatedByID"` // Foreign key relationship
	DateCreated  time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedByID  uint      `json:"updated_by_id" gorm:"column:updated_by"`   // Foreign key ID
	UpdatedBy    User      `json:"updated_by" gorm:"foreignKey:UpdatedByID"` // Foreign key relationship
	DateUpdated  time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllOrderRequest struct {
	OrderID      uint   `json:"id"`
	CustomerName string `json:"customer_name"`
	Items        *Item  `json:"items"`
	TmpTotal     int    `json:"tmp_total"`
	Potongan     int    `json:"potongan"`
	Total        int    `json:"total"`
	CreatedBy    *User  `json:"created_by"`
}

type CreateOrderRequest struct {
	CustomerName string              `json:"customer_name" binding:"required"`
	Items        []CreateItemRequest `json:"items" binding:"required"`
	Potongan     int                 `json:"potongan" binding:"required"`
}

type UpdateOrderRequest struct {
	CustomerName string              `json:"customer_name"`
	Items        []UpdateItemRequest `json:"items"`
	Potongan     int                 `json:"potongan"`
}
