package service

import (
	"final_project/pkg/models"
	"gorm.io/gorm"
	"time"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db}
}

func (os *OrderService) CreateOrder(idUser uint, req models.CreateOrderRequest) (*models.Order, error) {
	//create order reference to user
	var user models.User
	if err := os.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}

	order := models.Order{
		CustomerName: req.CustomerName,
		TmpTotal:     0,
		Potongan:     req.Potongan,
		Total:        0,
		CreatedBy:    user,
		DateCreated:  time.Now(),
		UpdatedBy:    user,
		DateUpdated:  time.Now(),
	}
	if err := os.db.Create(&order).Error; err != nil {
		return nil, err
	}

	var items []models.Item
	tmptotal := 0
	total := 0

	// if roles super admin, can add product to item by all user
	// if roles admin, can add product to item by admin
	if user.Roles == "SUPER ADMIN" {
		for _, item := range req.Items {
			var product models.Product
			if err := os.db.First(&product, item.ProductID).Error; err != nil {
				return nil, err
			}
			var itemModel models.Item
			itemModel.OrderID = order.OrderID
			itemModel.ProductID = item.ProductID
			itemModel.Harga = product.HargaJual
			itemModel.Jumlah = item.Jumlah
			itemModel.SubTotal = product.HargaJual * item.Jumlah
			tmptotal += itemModel.SubTotal
			itemModel.CreatedBy = user
			itemModel.DateCreated = time.Now()
			itemModel.UpdatedBy = user
			itemModel.DateUpdated = time.Now()
			items = append(items, itemModel)
		}
	} else {
		for _, item := range req.Items {
			var product models.Product
			// search for the product that created by the user
			if err := os.db.Where("created_by = ? and id_product = ?", idUser, item.ProductID).First(&product, item.ProductID).Error; err != nil {
				return nil, err
			}
			var itemModel models.Item
			itemModel.OrderID = order.OrderID
			itemModel.ProductID = item.ProductID
			itemModel.Harga = product.HargaJual
			itemModel.Jumlah = item.Jumlah
			itemModel.SubTotal = product.HargaJual * item.Jumlah
			tmptotal += itemModel.SubTotal
			itemModel.CreatedBy = user
			itemModel.DateCreated = time.Now()
			itemModel.UpdatedBy = user
			itemModel.DateUpdated = time.Now()
			items = append(items, itemModel)
		}
	}

	if err := os.db.Save(&items).Error; err != nil {
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

func (os *OrderService) GetAllOrders(roles string, idUser uint) ([]models.Order, error) {
	var orders []models.Order
	if roles == "SUPER ADMIN" {
		if err := os.db.Preload("Items").Find(&orders).Error; err != nil {
			return nil, err
		}
	} else {
		if err := os.db.Preload("Items").Where("created_by = ?", idUser).Find(&orders).Error; err != nil {
			return nil, err
		}
	}

	// search for the user who created the order
	for i, order := range orders {
		var user models.User
		if err := os.db.First(&user, order.CreatedByID).Error; err != nil {
			return nil, err
		}
		orders[i].CreatedBy = user
	}
	// search for the user who updated the order
	for i, order := range orders {
		var user models.User
		if err := os.db.First(&user, order.UpdatedByID).Error; err != nil {
			return nil, err
		}
		orders[i].UpdatedBy = user
	}

	return orders, nil
}

func (os *OrderService) UpdateOrder(orderID uint, roles string, idUser uint, req models.UpdateOrderRequest) (*models.Order, error) {
	var order models.Order
	if roles == "SUPER ADMIN" {
		if err := os.db.First(&order, orderID).Error; err != nil {
			return nil, err
		}
	} else {
		if err := os.db.Where("created_by = ?", idUser).First(&order, orderID).Error; err != nil {
			return nil, err
		}
	}

	var user models.User
	if err := os.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}
	tmptotal := 0
	total := 0

	if roles == "SUPER ADMIN" {
		for _, item := range req.Items {
			var itemModel models.Item
			if err := os.db.First(&itemModel, item.ItemID).Error; err != nil {
				return nil, err
			}

			var product models.Product
			if err := os.db.First(&product, item.ProductID).Error; err != nil {
				return nil, err
			}

			itemModel.ItemID = item.ItemID
			itemModel.OrderID = order.OrderID
			itemModel.ProductID = item.ProductID
			itemModel.Harga = product.HargaJual
			itemModel.Jumlah = item.Jumlah
			itemModel.SubTotal = product.HargaJual * item.Jumlah
			tmptotal += itemModel.SubTotal
			itemModel.UpdatedBy = user
			itemModel.DateUpdated = time.Now()
			if err := os.db.Save(&itemModel).Error; err != nil {
				return nil, err
			}
		}
	} else {
		for _, item := range req.Items {
			var itemModel models.Item
			if err := os.db.Where("created_by = ?", idUser).First(&itemModel, item.ItemID).Error; err != nil {
				return nil, err
			}

			var product models.Product
			if err := os.db.Where("created_by = ? and id_product = ?", idUser, item.ProductID).First(&product, item.ProductID).Error; err != nil {
				return nil, err
			}

			itemModel.ItemID = item.ItemID
			itemModel.OrderID = order.OrderID
			itemModel.ProductID = item.ProductID
			itemModel.Harga = product.HargaJual
			itemModel.Jumlah = item.Jumlah
			itemModel.SubTotal = product.HargaJual * item.Jumlah
			tmptotal += itemModel.SubTotal
			itemModel.UpdatedBy = user
			itemModel.DateUpdated = time.Now()
			if err := os.db.Save(&itemModel).Error; err != nil {
				return nil, err
			}
		}
	}

	total = tmptotal - order.Potongan
	order.TmpTotal = tmptotal
	order.Total = total
	order.UpdatedBy = user
	order.DateUpdated = time.Now()
	if err := os.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (os *OrderService) DeleteOrder(orderID uint, roles string, idUser uint) error {
	var order models.Order
	if roles == "SUPER ADMIN" {
		if err := os.db.First(&order, orderID).Error; err != nil {
			return err
		}
	} else {
		if err := os.db.Where("created_by = ?", idUser).First(&order, orderID).Error; err != nil {
			return err
		}
	}

	if err := os.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

func (os *OrderService) GetOrderByID(orderID uint, roles string, idUser uint) (*models.Order, error) {
	var order models.Order
	if roles == "SUPER ADMIN" {
		if err := os.db.Preload("Items").First(&order, orderID).Error; err != nil {
			return nil, err
		}
	} else {
		if err := os.db.Preload("Items").Where("created_by = ?", idUser).First(&order, orderID).Error; err != nil {
			return nil, err
		}
	}

	// search for the user who created the order
	var user models.User
	if err := os.db.First(&user, order.CreatedByID).Error; err != nil {
		return nil, err
	}
	order.CreatedBy = user

	// search for the user who updated the order
	if err := os.db.First(&user, order.UpdatedByID).Error; err != nil {
		return nil, err
	}
	order.UpdatedBy = user

	return &order, nil
}
