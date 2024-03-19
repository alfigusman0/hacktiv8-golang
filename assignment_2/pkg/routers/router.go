package routers

import (
	"assignment_2/pkg/controller"
	"assignment_2/pkg/service"
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
	orderController.Routes(api)

	itemService := service.NewItemService(gorm)
	itemController := controller.NewItemController(itemService)
	itemController.Routes(api)

	return r
}
