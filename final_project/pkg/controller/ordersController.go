package controller

import (
	"final_project/pkg/models"
	"final_project/pkg/service"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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

	routeGroup.GET("", IsAuth, o.GetAllOrder)
	routeGroup.POST("", IsAuth, o.CreateOrder)
	routeGroup.GET("/:id", IsAuth, o.GetOrderByID)
	routeGroup.PUT("/:id", IsAuth, o.UpdateOrder)
	routeGroup.DELETE("/:id", IsAuth, o.DeleteOrder)
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))

	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	order, err := o.service.CreateOrder(idUser, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "",
		"data":    order,
	})
}

func (o *OrderController) GetAllOrder(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	orders, err := o.service.GetAllOrders(roles, idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "success",
			"message": "no data found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "",
		"data":    orders,
	})
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid order id",
		})
		return
	}
	var req models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	order, err := o.service.UpdateOrder(uint(orderID), idUser, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "",
		"data":    order,
	})
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid order id",
		})
		return
	}

	err = o.service.DeleteOrder(uint(orderID), roles, idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": fmt.Sprintf("order with id %d deleted", orderID),
	})
}

func (o *OrderController) GetOrderByID(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid product id",
		})
		return
	}

	order, err := o.service.GetOrderByID(uint(orderID), roles, idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "",
		"data":    order,
	})
}
