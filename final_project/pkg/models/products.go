package models

import "time"

type Product struct {
	ProductID   uint      `json:"id" gorm:"column:id_product;primaryKey;autoIncrement"`
	ProductName string    `json:"product" gorm:"column:product;"`
	HargaBeli   float64   `json:"harga_beli" gorm:"column:harga_beli"`
	HargaJual   float64   `json:"harga_jual" gorm:"column:harga_jual"`
	Stok        int       `json:"stok" gorm:"column:stok"`
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedByID;references:UserID"`
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	UpdatedBy   User      `json:"updated_by" gorm:"foreignKey:UpdatedByID;references"`
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllProductRequest struct {
	ProductID   uint    `json:"id"`
	ProductName string  `json:"product"`
	HargaBeli   float64 `json:"harga_beli"`
	HargaJual   float64 `json:"harga_jual"`
	Stok        int     `json:"stok"`
	CreatedBy   *User   `json:"created_by"`
}

type CreateProductRequest struct {
	ProductName string  `json:"product" binding:"required"`
	HargaBeli   float64 `json:"harga_beli" binding:"required"`
	HargaJual   float64 `json:"harga_jual" binding:"required"`
	Stok        int     `json:"stok" binding:"required"`
	CreatedByID uint    `json:"created_by" binding:"required"`
	UpdatedByID uint    `json:"updated_by" binding:"required"`
}

type UpdateProductRequest struct {
	ProductID   uint    `json:"id" binding:"required"`
	ProductName string  `json:"product" binding:"required"`
	HargaBeli   float64 `json:"harga_beli" binding:"required"`
	HargaJual   float64 `json:"harga_jual" binding:"required"`
	Stok        int     `json:"stok" binding:"required"`
	UpdatedByID uint    `json:"updated_by" binding:"required"`
}
