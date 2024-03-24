package routers

import (
	"database/sql"
	"final_project/pkg/controller"
	"final_project/pkg/middleware"
	"final_project/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *sql.DB, gorm *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// API
	api := r.Group("/api")

	orderService := service.NewOrderService(gorm)
	orderController := controller.NewOrderController(orderService)
	orderController.Routes(api, middleware.IsAuth(gorm))

	itemService := service.NewItemService(gorm)
	itemController := controller.NewItemController(itemService)
	itemController.Routes(api, middleware.IsAuth(gorm))

	productService := service.NewProductService(gorm)
	productController := controller.NewProductController(productService)
	productController.Routes(api, middleware.IsAuth(gorm))

	usersService := service.NewUsersService(gorm)
	userController := controller.NewUserController(usersService)
	userController.Routes(api, middleware.IsAuth(gorm))

	return r
}
