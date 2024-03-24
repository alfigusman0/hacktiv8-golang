package service

import (
	"final_project/pkg/models"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := ps.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) CreateProduct(req models.CreateProductRequest) (*models.Product, error) {
	product := models.Product{
		ProductName: req.ProductName,
		HargaBeli:   req.HargaBeli,
		HargaJual:   req.HargaJual,
		Stok:        req.Stok,
	}
	if err := ps.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) UpdateProduct(productID uint, req models.UpdateProductRequest) (*models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	product.ProductName = req.ProductName
	product.HargaBeli = req.HargaBeli
	product.HargaJual = req.HargaJual
	product.Stok = req.Stok
	if err := ps.db.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) DeleteProduct(productID uint) error {
	var product models.Product
	if err := ps.db.First(&product, productID).Error; err != nil {
		return err
	}
	if err := ps.db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
