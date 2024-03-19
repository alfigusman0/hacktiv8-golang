package service

import (
	"assignment_2/pkg/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db}
}

func (os *OrderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := os.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (os *OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := os.db.Preload("Items").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (os *OrderService) CreateOrder(req models.CreateOrderRequest) (*models.Order, error) {
	now := time.Now()
	order := models.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    now,
	}
	if err := os.db.Create(&order).Error; err != nil {
		return nil, err
	}

	var items []models.Item

	for _, item := range req.Items {
		var itemModel models.Item
		itemModel.ItemCode = item.ItemCode
		itemModel.Description = item.Description
		itemModel.Quantity = item.Quantity
		itemModel.OrderID = &order.OrderID
		items = append(items, itemModel)
	}

	if err := os.db.Create(&items).Error; err != nil {
		return nil, err
	}

	order.Items = items
	return &order, nil
}

func (os *OrderService) UpdateOrder(orderID uint, req models.UpdateOrderRequest) (*models.Order, error) {
	order, err := os.GetOrderByID(orderID)
	fmt.Printf("order: %+v\n", order)
	if err != nil {
		return nil, err
	}

	order.CustomerName = req.CustomerName

	if err := os.db.Save(&order).Error; err != nil {
		return nil, err
	}

	var items []models.Item

	for _, item := range req.Items {
		var itemModel models.Item
		itemModel.ItemID = item.ItemID
		itemModel.ItemCode = item.ItemCode
		itemModel.Description = item.Description
		itemModel.Quantity = item.Quantity
		itemModel.OrderID = &order.OrderID
		items = append(items, itemModel)
	}

	if err := os.db.Save(&items).Error; err != nil {
		return nil, err
	}

	order.Items = items

	return order, nil
}

func (os *OrderService) DeleteOrder(orderID uint) error {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if err := os.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}
