package service

import (
	"final_project/pkg/models"
	"gorm.io/gorm"
	"time"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db}
}

func (is *ItemService) CreateItem(roles string, idUser uint, req models.CreateItemRequest) (*models.Item, error) {
	//create item reference to user
	var user models.User
	if err := is.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}

	// if roles super admin, can add product to item by all user
	// if roles admin, can add product to item by admin
	var item models.Item
	if user.Roles == "SUPER ADMIN" {
		item = models.Item{
			OrderID:     req.OrderID,
			ProductID:   req.ProductID,
			Harga:       req.Harga,
			Jumlah:      req.Jumlah,
			SubTotal:    req.Harga * req.Jumlah,
			CreatedByID: idUser,
			UpdatedByID: idUser,
		}

		if err := is.db.Create(&item).Error; err != nil {
			return nil, err
		}
	} else {
		var products []models.Product
		if err := is.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return nil, err
		}

		for _, product := range products {
			item = models.Item{
				OrderID:     req.OrderID,
				ProductID:   product.ProductID,
				Harga:       req.Harga,
				Jumlah:      req.Jumlah,
				SubTotal:    req.Harga * req.Jumlah,
				CreatedByID: idUser,
				UpdatedByID: idUser,
			}

			if err := is.db.Create(&item).Error; err != nil {
				return nil, err
			}
		}
	}

	var order models.Order
	if err := is.db.First(&order, req.OrderID).Error; err != nil {
		return nil, err
	}
	tmptotal := 0
	total := 0
	items, err := is.GetAllItems(roles, idUser)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		tmptotal += item.SubTotal
	}
	total = tmptotal - order.Potongan
	order.TmpTotal = tmptotal
	order.Total = total
	if err := is.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (is *ItemService) GetAllItems(roles string, idUser uint) ([]models.Item, error) {
	var items []models.Item
	if roles == "SUPER ADMIN" {
		if err := is.db.Find(&items).Error; err != nil {
			return nil, err
		}
	} else {
		var products []models.Product
		if err := is.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return nil, err
		}

		var item []models.Item
		for _, product := range products {
			if err := is.db.Where("product_id = ?", product.ProductID).Find(&item).Error; err != nil {
				return nil, err
			}
		}
	}

	return items, nil
}

func (is *ItemService) UpdateItem(id uint, roles string, idUser uint, req models.UpdateItemRequest) (*models.Item, error) {
	var item models.Item
	if roles == "SUPER ADMIN" {
		if err := is.db.First(&item, id).Error; err != nil {
			return nil, err
		}
	} else {
		var products []models.Product
		if err := is.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return nil, err
		}

		var item []models.Item
		for _, product := range products {
			if err := is.db.Where("product_id = ?", product.ProductID).Find(&item).Error; err != nil {
				return nil, err
			}
		}
	}

	var user models.User
	if err := is.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}

	item.OrderID = req.OrderID
	item.ProductID = req.ProductID
	item.Harga = req.Harga
	item.Jumlah = req.Jumlah
	item.SubTotal = req.Harga * req.Jumlah
	item.UpdatedByID = idUser
	item.DateUpdated = time.Now()

	if err := is.db.Save(&item).Error; err != nil {
		return nil, err
	}

	var order models.Order
	if err := is.db.First(&order, item.OrderID).Error; err != nil {
		return nil, err
	}
	tmptotal := 0
	total := 0
	items, err := is.GetAllItems(roles, idUser)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		tmptotal += item.SubTotal
	}
	total = tmptotal - order.Potongan

	order.TmpTotal = tmptotal
	order.Total = total
	if err := is.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (is *ItemService) DeleteItem(id uint, roles string, idUser uint) error {
	var item models.Item
	if roles == "SUPER ADMIN" {
		if err := is.db.First(&item, id).Error; err != nil {
			return err
		}
	} else {
		var products []models.Product
		if err := is.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return err
		}

		var item []models.Item
		for _, product := range products {
			if err := is.db.Where("product_id = ?", product.ProductID).Find(&item).Error; err != nil {
				return err
			}
		}
	}

	if err := is.db.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}

func (is *ItemService) GetItemByID(id uint, roles string, idUser uint) (*models.Item, error) {
	var item models.Item
	if roles == "SUPER ADMIN" {
		if err := is.db.First(&item, id).Error; err != nil {
			return nil, err
		}
	} else {
		var products []models.Product
		if err := is.db.Where("created_by = ?", idUser).Find(&products).Error; err != nil {
			return nil, err
		}

		var item []models.Item
		for _, product := range products {
			if err := is.db.Where("product_id = ?", product.ProductID).Find(&item).Error; err != nil {
				return nil, err
			}
		}
	}

	return &item, nil
}
