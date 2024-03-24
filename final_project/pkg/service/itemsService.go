package service

import (
	"final_project/pkg/models"
	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db}
}

func (is *ItemService) GetAllItems() ([]models.Item, error) {
	var items []models.Item
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
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Harga:     req.Harga,
		Jumlah:    req.Jumlah,
		SubTotal:  req.Harga * float64(req.Jumlah),
	}
	if err := is.db.Create(&item).Error; err != nil {
		return nil, err
	}

	var order models.Order
	if err := is.db.First(&order, item.OrderID).Error; err != nil {
		return nil, err
	}
	tmptotal := 0.0
	total := 0.0
	items, err := is.GetAllItems()
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

func (is *ItemService) UpdateItem(itemID uint, req models.UpdateItemRequest) (*models.Item, error) {
	item, err := is.GetItem(itemID)
	if err != nil {
		return nil, err
	}

	item.OrderID = req.OrderID
	item.ProductID = req.ProductID
	item.Harga = req.Harga
	item.Jumlah = req.Jumlah
	item.SubTotal = req.Harga * float64(req.Jumlah)

	if err := is.db.Save(&item).Error; err != nil {
		return nil, err
	}

	var order models.Order
	if err := is.db.First(&order, item.OrderID).Error; err != nil {
		return nil, err
	}
	tmptotal := 0.0
	total := 0.0
	items, err := is.GetAllItems()
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
