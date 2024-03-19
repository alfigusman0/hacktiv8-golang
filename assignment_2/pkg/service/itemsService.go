package service

import (
	"assignment_2/pkg/models"
	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db}
}

func (is *ItemService) GetAllItems() ([]models.GetAllItemRequest, error) {
	var items []models.GetAllItemRequest
	if err := is.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (is *ItemService) GetItem(itemID uint) (*models.Item, error) {
	var item models.Item
	if err := is.db.First(&item, itemID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (is *ItemService) CreateItem(req models.CreateItemRequest) (*models.Item, error) {
	item := models.Item{
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}
	if err := is.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (is *ItemService) UpdateItem(itemID uint, req models.UpdateItemRequest) (*models.Item, error) {
	item, err := is.GetItem(itemID)
	if err != nil {
		return nil, err
	}
	item.ItemCode = req.ItemCode
	item.Description = req.Description
	item.Quantity = req.Quantity
	if err := is.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (is *ItemService) DeleteItem(itemID uint) error {
	item, err := is.GetItem(itemID)
	if err != nil {
		return err
	}
	if err := is.db.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
