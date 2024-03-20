package controller

import (
	"assignment_3/pkg/models"
	"assignment_3/pkg/service"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UsersService
}

func NewUserController(service *service.UsersService) *UserController {
	return &UserController{service}
}

func (u *UserController) Routes(r *gin.RouterGroup, IsAuth gin.HandlerFunc) {
	routeGroup := r.Group("/users")

	routeGroup.GET("", IsAuth, u.GetAllUsers)
	routeGroup.POST("", IsAuth, u.CreateUser)
	routeGroup.GET("/:id", IsAuth, u.GetUserByID)
	routeGroup.PUT("/:id", IsAuth, u.UpdateUser)
	routeGroup.DELETE("/:id", IsAuth, u.DeleteUser)

	//signin and signout route
	r.POST("/sign-in", u.SignIn)
	r.POST("/sign-out", u.SignOut)
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
	getHeader := c.GetHeader("Authorization")
	if getHeader == "" {
		c.JSON(400, gin.H{"error": "token not found"})
		return
	}
	split := strings.Split(getHeader, "Bearer ")
	errInvalidToken := errors.New("invalid token")
	if len(split) != 2 {
		c.AbortWithStatusJSON(401, gin.H{
			"message": errInvalidToken.Error(),
		})
		return
	}
	getToken := split[1]
	err := u.service.SignOut(getToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "signout success"})
}
