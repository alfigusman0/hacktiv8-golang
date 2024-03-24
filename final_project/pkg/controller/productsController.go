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

	routeGroup.GET("", pc.GetAllProducts)
	routeGroup.POST("", pc.CreateProduct)
	routeGroup.GET("/:id", pc.GetProduct)
	routeGroup.PUT("/:id", pc.UpdateProduct)
	routeGroup.DELETE("/:id", pc.DeleteProduct)
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["role"] != "SUPER ADMIN" {
		products, err := pc.service.GetAllProducts()
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
			"data":    products,
		})
	} else {
		products, err := pc.service.GetAllProductsByCreatedBy(uint(userData["id"].(float64)))
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
			"data":    products,
		})
	}
}

func (pc *ProductController) GetProduct(c *gin.Context) {
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
	if userData["role"] != "SUPER ADMIN" {
		product, err := pc.service.GetProductByID(uint(id))
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
	} else {
		product, err := pc.service.GetProductByIDAndCreatedBy(uint(id), uint(userData["id"].(float64)))
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
	product, err := pc.service.CreateProduct(req)
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
	if userData["role"] != "SUPER ADMIN" {
		product, err := pc.service.UpdateProduct(uint(id), req)
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
	} else {
		product, err := pc.service.UpdateProductByCreatedBy(uint(id), uint(userData["id"].(float64)), req)
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
	if userData["role"] != "SUPER ADMIN" {
		err = pc.service.DeleteProduct(uint(id))
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
	} else {
		err = pc.service.DeleteProductByCreatedBy(uint(id), uint(userData["id"].(float64)))
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
}
