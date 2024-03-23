package models

import "time"

type Item struct {
	ItemID      uint      `json:"id_item" gorm:"column:id_item;primaryKey;autoIncrement"`
	OrderID     uint      `json:"id_order" gorm:"column:id_order"`
	ProductID   uint      `json:"id_product" gorm:"column:id_product"`
	Harga       float64   `json:"harga" gorm:"column:harga"`
	Jumlah      int       `json:"jumlah" gorm:"column:jumlah"`
	SubTotal    float64   `json:"sub_total" gorm:"column:sub_total"`
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID;references:UserID"`
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedBy   User      `json:"updated_by" gorm:"foreignKey:UpdatedByID;references:UserID"`
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllItemRequest struct {
	ItemID    uint    `json:"id_item"`
	OrderID   uint    `json:"id_order"`
	ProductID uint    `json:"id_product"`
	Harga     float64 `json:"harga"`
	Jumlah    int     `json:"jumlah"`
	SubTotal  float64 `json:"sub_total"`
	CreatedBy *User   `json:"created_by"`
}

type CreateItemRequest struct {
	OrderID     uint    `json:"id_order" bidnding:"required"`
	ProductID   uint    `json:"id_product" bidnding:"required"`
	Harga       float64 `json:"harga" bidnding:"required"`
	Jumlah      int     `json:"jumlah" bidnding:"required"`
	SubTotal    float64 `json:"sub_total" bidnding:"required"`
	CreatedByID uint    `json:"created_by" bidnding:"required"`
	UpdatedByID uint    `json:"updated_by" bidnding:"required"`
}

type UpdateItemRequest struct {
	OrderID     uint    `json:"id_order"`
	ProductID   uint    `json:"id_product"`
	Harga       float64 `json:"harga"`
	Jumlah      int     `json:"jumlah"`
	SubTotal    float64 `json:"sub_total"`
	UpdatedByID uint    `json:"updated_by" bidnding:"required"`
}
