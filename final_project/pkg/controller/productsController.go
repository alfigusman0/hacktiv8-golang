package controller

import (
	"final_project/pkg/models"
	"final_project/pkg/service"
	"github.com/gin-gonic/gin"
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

	routeGroup.GET("", pc.GetAllProducts)
	routeGroup.POST("", pc.CreateProduct)
	routeGroup.GET("/:id", pc.GetProduct)
	routeGroup.PUT("/:id", pc.UpdateProduct)
	routeGroup.DELETE("/:id", pc.DeleteProduct)
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	product, err := pc.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.service.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.service.UpdateProduct(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	err = pc.service.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
