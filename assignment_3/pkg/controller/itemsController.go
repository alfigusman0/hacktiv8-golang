package controller

import (
	"assignment_3/pkg/models"
	"assignment_3/pkg/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	service *service.ItemService
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (i *ItemController) Routes(r *gin.RouterGroup, IsAuth gin.HandlerFunc) {
	routeGroup := r.Group("/items")

	routeGroup.GET("", i.GetAllItem)
	routeGroup.POST("", IsAuth, i.CreateItem)
	routeGroup.GET("/:id", i.GetItemByID)
	routeGroup.PUT("/:id", IsAuth, i.UpdateItem)
	routeGroup.DELETE("/:id", IsAuth, i.DeleteItem)
}

func (i *ItemController) GetAllItem(c *gin.Context) {
	items, err := i.service.GetAllItems()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, items)
}

func (i *ItemController) GetItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	item, err := i.service.GetItem(uint(itemID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, item)
}

func (i *ItemController) CreateItem(c *gin.Context) {
	var req models.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	item, err := i.service.CreateItem(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, item)
}

func (i *ItemController) UpdateItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	var req models.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	item, err := i.service.UpdateItem(uint(itemID), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, item)
}

func (i *ItemController) DeleteItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	err = i.service.DeleteItem(uint(itemID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("item with id %d deleted", itemID)})
}
