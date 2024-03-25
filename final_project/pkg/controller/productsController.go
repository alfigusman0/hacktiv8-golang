package controller

import (
	"final_project/pkg/models"
	"final_project/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service}
}

func (pc *ProductController) Routes(r *gin.RouterGroup, IsAuth gin.HandlerFunc) {
	routeGroup := r.Group("/products")

	routeGroup.GET("", IsAuth, pc.GetAllProducts)
	routeGroup.POST("", IsAuth, pc.CreateProduct)
	routeGroup.GET("/:id", IsAuth, pc.GetProductByID)
	routeGroup.PUT("/:id", IsAuth, pc.UpdateProduct)
	routeGroup.DELETE("/:id", IsAuth, pc.DeleteProduct)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	product, err := pc.service.CreateProduct(req, idUser)
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
		"message": "product created",
		"data":    product,
	})
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	products, err := pc.service.GetAllProducts(roles, idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if len(products) == 0 {
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
		"data":    products,
	})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid product id",
		})
		return
	}
	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)
	product, err := pc.service.UpdateProduct(uint(id), roles, idUser, req)
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
		"message": "product updated",
		"data":    product,
	})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid product id",
		})
		return
	}
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)
	err = pc.service.DeleteProduct(uint(id), roles, idUser)
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
		"message": "product deleted",
	})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "invalid product id",
		})
		return
	}

	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	idUser := uint(userData["id"].(float64))
	roles := userData["roles"].(string)

	product, err := pc.service.GetProductByID(uint(id), roles, idUser)
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
		"data":    product,
	})
}
