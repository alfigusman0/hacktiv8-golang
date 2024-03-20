package routers

import (
	"assignment_3/pkg/controller"
	"assignment_3/pkg/middleware"
	"assignment_3/pkg/service"
	"database/sql"

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

	usersService := service.NewUsersService(gorm)
	userController := controller.NewUserController(usersService)
	userController.Routes(api, middleware.IsAuth(gorm))

	return r
}
