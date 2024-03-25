package controller

import (
	"final_project/pkg/models"
	"final_project/pkg/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ItemController struct {
	service *service.ItemService
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (i *ItemController) Routes(r *gin.RouterGroup, IsAuth gin.HandlerFunc) {
	routeGroup := r.Group("/items")

	routeGroup.GET("", IsAuth, i.GetAllItem)
	routeGroup.POST("", IsAuth, i.CreateItem)
	routeGroup.GET("/:id", IsAuth, i.GetItemByID)
	routeGroup.PUT("/:id", IsAuth, i.UpdateItem)
	routeGroup.DELETE("/:id", IsAuth, i.DeleteItem)
}

func (i *ItemController) CreateItem(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	roles := userData["roles"].(string)
	idUser := uint(userData["id"].(float64))

	var req models.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	item, err := i.service.CreateItem(roles, idUser, req)
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
		"data":    item,
	})
}

func (i *ItemController) GetAllItem(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	items, err := i.service.GetAllItems(roles, idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "",
		"data":    items,
	})
}

func (i *ItemController) UpdateItem(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid item id",
		})
		return
	}
	var req models.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	item, err := i.service.UpdateItem(uint(id), roles, idUser, req)
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
		"data":    item,
	})
}

func (i *ItemController) DeleteItem(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid item id",
		})
		return
	}
	err = i.service.DeleteItem(uint(id), roles, idUser)
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
		"message": fmt.Sprintf("item with id %d deleted", id),
	})
}

func (i *ItemController) GetItemByID(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid item id",
		})
		return
	}

	item, err := i.service.GetItemByID(uint(id), roles, idUser)
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
		"data":    item,
	})
}
