package models

type Item struct {
	ItemID      uint   `json:"item_id" gorm:"column:item_id;primaryKey;autoIncrement"`
	ItemCode    string `json:"item_code" gorm:"column:item_code"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
	OrderID     *uint  `json:"order_id" gorm:"column:order_id"`
}

type GetAllItemRequest struct {
	ItemID      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type CreateItemRequest struct {
	ItemCode    string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	OrderID     *uint  `json:"order_id"`
}

type UpdateItemRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     *uint  `json:"order_id"`
}
