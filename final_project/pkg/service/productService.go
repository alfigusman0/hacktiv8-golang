package service

import (
	"final_project/pkg/models"
	"time"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) CreateProduct(req models.CreateProductRequest, idUser uint) (*models.Product, error) {
	//create product reference to user
	var user models.User
	if err := ps.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}
	product := models.Product{
		ProductName: req.ProductName,
		HargaBeli:   req.HargaBeli,
		HargaJual:   req.HargaJual,
		Stok:        req.Stok,
		CreatedBy:   user,
		DateCreated: time.Now(),
		UpdatedBy:   user,
		DateUpdated: time.Now(),
	}
	if err := ps.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) GetAllProducts(roles string, idUser uint) ([]models.Product, error) {
	var products []models.Product
	if roles == "SUPER ADMIN" {
		if err := ps.db.Find(&products).Error; err != nil {
			return nil, err
		}
	} else {
		if err := ps.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return nil, err
		}
	}

	// search for the user who created the product
	for i, product := range products {
		var user models.User
		if err := ps.db.First(&user, product.CreatedByID).Error; err != nil {
			return nil, err
		}
		products[i].CreatedBy = user
	}
	// search for the user who updated the product
	for i, product := range products {
		var user models.User
		if err := ps.db.First(&user, product.UpdatedByID).Error; err != nil {
			return nil, err
		}
		products[i].UpdatedBy = user
	}

	return products, nil
}

func (ps *ProductService) UpdateProduct(productID uint, roles string, idUser uint, req models.UpdateProductRequest) (*models.Product, error) {
	var product models.Product
	if roles == "SUPER ADMIN" {
		if err := ps.db.First(&product, productID).Error; err != nil {
			return nil, err
		}
	} else {
		if err := ps.db.Where("created_by = ?", idUser).First(&product, productID).Error; err != nil {
			return nil, err
		}
	}
	var user models.User
	if err := ps.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}
	product.ProductName = req.ProductName
	product.HargaBeli = req.HargaBeli
	product.HargaJual = req.HargaJual
	product.Stok = req.Stok
	product.UpdatedBy = user
	product.DateUpdated = time.Now()
	if err := ps.db.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) DeleteProduct(productID uint, roles string, idUser uint) error {
	var product models.Product
	if roles == "SUPER ADMIN" {
		if err := ps.db.First(&product, productID).Error; err != nil {
			return err
		}
	} else {
		if err := ps.db.Where("created_by = ?", idUser).First(&product, productID).Error; err != nil {
			return err
		}
	}

	if err := ps.db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetProductByID(productID uint, roles string, idUser uint) (*models.Product, error) {
	var product models.Product
	if roles == "SUPER ADMIN" {
		if err := ps.db.First(&product, productID).Error; err != nil {
			return nil, err
		}
	} else {
		if err := ps.db.Where("created_by = ?", idUser).First(&product, productID).Error; err != nil {
			return nil, err
		}
	}

	// search for the user who created the product
	var user models.User
	if err := ps.db.First(&user, product.CreatedByID).Error; err != nil {
		return nil, err
	}
	product.CreatedBy = user

	// search for the user who updated the product
	var user2 models.User
	if err := ps.db.First(&user2, product.UpdatedByID).Error; err != nil {
		return nil, err
	}
	product.UpdatedBy = user2

	return &product, nil
}
