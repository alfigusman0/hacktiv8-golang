package controller

import (
	"final_project/pkg/models"
	"final_project/pkg/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (o *OrderController) Routes(r *gin.RouterGroup, IsAuth gin.HandlerFunc) {
	routeGroup := r.Group("/orders")

	routeGroup.GET("", o.GetAllOrder)
	routeGroup.POST("", IsAuth, o.CreateOrder)
	routeGroup.GET("/:id", o.GetOrderByID)
	routeGroup.PUT("/:id", IsAuth, o.UpdateOrder)
	routeGroup.DELETE("/:id", IsAuth, o.DeleteOrder)
}

func (o *OrderController) GetAllOrder(c *gin.Context) {
	orders, err := o.service.GetAllOrders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, orders)
}

func (o *OrderController) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid order id"})
		return
	}
	order, err := o.service.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	order, err := o.service.CreateOrder(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, order)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid order id"})
		return
	}
	var req models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	order, err := o.service.UpdateOrder(uint(orderID), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid order id"})
		return
	}
	err = o.service.DeleteOrder(uint(orderID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("order with id %d deleted", orderID)})
}
