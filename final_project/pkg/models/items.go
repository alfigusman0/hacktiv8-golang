package models

import "time"

type Item struct {
	ItemID      uint      `json:"id_item" gorm:"column:id_item;primaryKey;autoIncrement"`
	OrderID     uint      `json:"id_order" gorm:"column:id_order"`
	ProductID   uint      `json:"id_product" gorm:"column:id_product"`
	Harga       int       `json:"harga" gorm:"column:harga"`
	Jumlah      int       `json:"jumlah" gorm:"column:jumlah"`
	SubTotal    int       `json:"sub_total" gorm:"column:sub_total"`
	CreatedByID uint      `json:"created_by_id" gorm:"column:created_by"`   // Foreign key ID
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID"` // Foreign key relationship
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedByID uint      `json:"updated_by_id" gorm:"column:updated_by"`   // Foreign key ID
	UpdatedBy   User      `json:"updated_by" gorm:"foreignKey:UpdatedByID"` // Foreign key relationship
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllItemRequest struct {
	ItemID    uint  `json:"id_item"`
	OrderID   uint  `json:"id_order"`
	ProductID uint  `json:"id_product"`
	Harga     int   `json:"harga"`
	Jumlah    int   `json:"jumlah"`
	SubTotal  int   `json:"sub_total"`
	CreatedBy *User `json:"created_by"`
}

type CreateItemRequest struct {
	OrderID   uint `json:"id_order"`
	ProductID uint `json:"id_product" bidnding:"required"`
	Harga     int  `json:"harga"`
	Jumlah    int  `json:"jumlah" bidnding:"required"`
	SubTotal  int  `json:"sub_total"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

type UpdateItemRequest struct {
	ItemID    uint `json:"id_item" bidnding:"required"`
	OrderID   uint `json:"id_order"`
	ProductID uint `json:"id_product" bidnding:"required"`
	Harga     int  `json:"harga"`
	Jumlah    int  `json:"jumlah" bidnding:"required"`
	SubTotal  int  `json:"sub_total"`
	UpdatedBy uint `json:"updated_by"`
}
