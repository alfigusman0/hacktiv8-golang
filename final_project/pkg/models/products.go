package models

import "time"

type Product struct {
	ProductID   uint      `json:"id" gorm:"column:id_product;primaryKey;autoIncrement"`
	ProductName string    `json:"product" gorm:"column:product;"`
	HargaBeli   int       `json:"harga_beli" gorm:"column:harga_beli"`
	HargaJual   int       `json:"harga_jual" gorm:"column:harga_jual"`
	Stok        int       `json:"stok" gorm:"column:stok"`
	CreatedByID uint      `json:"created_by_id" gorm:"column:created_by"`   // Foreign key ID
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID"` // Foreign key relationship
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedByID uint      `json:"updated_by_id" gorm:"column:updated_by"`   // Foreign key ID
	UpdatedBy   User      `json:"updated_by" gorm:"foreignKey:UpdatedByID"` // Foreign key relationship
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllProductRequest struct {
	ProductID   uint   `json:"id"`
	ProductName string `json:"product"`
	HargaBeli   int    `json:"harga_beli"`
	HargaJual   int    `json:"harga_jual"`
	Stok        int    `json:"stok"`
	CreatedBy   *User  `json:"created_by"`
}

type CreateProductRequest struct {
	ProductName string `json:"product" binding:"required"`
	HargaBeli   int    `json:"harga_beli"`
	HargaJual   int    `json:"harga_jual"`
	Stok        int    `json:"stok"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type UpdateProductRequest struct {
	ProductName string `json:"product" binding:"required"`
	HargaBeli   int    `json:"harga_beli"`
	HargaJual   int    `json:"harga_jual"`
	Stok        int    `json:"stok"`
	UpdatedByID uint   `json:"updated_by"`
}
