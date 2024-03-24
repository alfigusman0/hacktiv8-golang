package service

import (
	"final_project/pkg/models"
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
		TmpTotal:     0,
		Potongan:     req.Potongan,
		Total:        0,
		DateCreated:  now,
		DateUpdated:  now,
	}
	if err := os.db.Create(&order).Error; err != nil {
		return nil, err
	}

	var items []models.Item
	tmptotal := 0.0
	total := 0.0

	for _, item := range req.Items {
		var itemModel models.Item
		itemModel.OrderID = order.OrderID
		itemModel.ProductID = item.ProductID
		itemModel.Harga = item.Harga
		itemModel.Jumlah = item.Jumlah
		itemModel.SubTotal = item.Harga * float64(item.Jumlah)
		tmptotal += itemModel.SubTotal
		items = append(items, itemModel)
	}

	if err := os.db.Create(&items).Error; err != nil {
		return nil, err
	}

	total = tmptotal - req.Potongan
	order.TmpTotal = tmptotal
	order.Total = total
	if err := os.db.Save(&order).Error; err != nil {
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
	tmptotal := 0.0
	total := 0.0

	for _, item := range req.Items {
		var itemModel models.Item
		itemModel.OrderID = order.OrderID
		itemModel.ProductID = item.ProductID
		itemModel.Harga = item.Harga
		itemModel.Jumlah = item.Jumlah
		itemModel.SubTotal = item.Harga * float64(item.Jumlah)
		tmptotal += itemModel.SubTotal
		items = append(items, itemModel)
	}

	if err := os.db.Save(&items).Error; err != nil {
		return nil, err
	}

	total = tmptotal - order.Potongan
	order.TmpTotal = tmptotal
	order.Total = total
	if err := os.db.Save(&order).Error; err != nil {
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
