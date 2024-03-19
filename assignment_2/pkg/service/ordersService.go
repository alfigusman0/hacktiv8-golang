package service

import (
	"assignment_2/pkg/models"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db}
}

func (os *OrderService) GetAllOrders() ([]models.GetAllOrderRequest, error) {
	var orders []models.GetAllOrderRequest
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

func (os *OrderService) CreateOrder(req models.CreateOrderRequest) (*models.CreateOrderRequest, error) {
	order := models.CreateOrderRequest{
		CustomerName: req.CustomerName,
		Items:        req.Items,
	}
	if err := os.db.Create(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (os *OrderService) UpdateOrder(orderID uint, req models.UpdateOrderRequest) (*models.UpdateOrderRequest, error) {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	order.CustomerName = req.CustomerName
	order.Items = req.Items
	if err := os.db.Save(&order).Error; err != nil {
		return nil, err
	}
	return &req, nil
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
