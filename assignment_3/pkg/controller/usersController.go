package controller

import (
	"assignment_3/pkg/models"
	"assignment_3/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UsersService
}

func NewUserController(service *service.UsersService) *UserController {
	return &UserController{service}
}

func (u *UserController) Routes(r *gin.RouterGroup) {
	routeGroup := r.Group("/users")

	routeGroup.GET("", u.GetAllUsers)
	routeGroup.POST("", u.CreateUser)
	routeGroup.GET("/:id", u.GetUserByID)
	routeGroup.PUT("/:id", u.UpdateUser)
	routeGroup.DELETE("/:id", u.DeleteUser)

	//signin and signout route
	r.POST("/signin", u.SignIn)
	r.POST("/signout", u.SignOut)
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := u.service.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, createdUser)
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := u.service.UpdateUser(uint(id), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedUser)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = u.service.DeleteUser(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted"})
}

func (u *UserController) SignIn(c *gin.Context) {
	var signInRequest models.SignInRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := u.service.SignIn(signInRequest, c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (u *UserController) SignOut(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{"error": "token not found"})
		return
	}
	err := u.service.SignOut(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "signout success"})
}
