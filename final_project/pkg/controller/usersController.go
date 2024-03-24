package controller

import (
	"errors"
	"final_project/pkg/models"
	"final_project/pkg/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["role"] != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"status":  "error",
			"message": "Access Denied!",
		})
		return
	}
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
	}
	c.JSON(200, users)
}

func (u *UserController) CreateUser(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["role"] != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"status":  "error",
			"message": "Access Denied!",
		})
		return
	}
	var user models.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	createdUser, err := u.service.CreateUser(user)
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
		"data":    createdUser,
	})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["UserID"] != id && userData["role"] == "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"status":  "error",
			"message": "Access Denied!",
		})
		return
	}

	user, err := u.service.GetUserByID(uint(id))
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
		"data":    user,
	})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["UserID"] != id && userData["role"] == "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"status":  "error",
			"message": "Access Denied!",
		})
		return
	}
	var user models.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	updatedUser, err := u.service.UpdateUser(uint(id), user)
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
		"data":    updatedUser,
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	duser, _ := c.Get("user")
	userData := duser.(jwt.MapClaims)
	if userData["role"] != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"status":  "error",
			"message": "Access Denied!",
		})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	err = u.service.DeleteUser(uint(id))
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
		"message": "user deleted",
	})
}

func (u *UserController) SignIn(c *gin.Context) {
	var signInRequest models.SignInRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	token, err := u.service.SignIn(signInRequest, c.GetHeader("User-Agent"), c.ClientIP())
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
		"token":   token,
	})
}

func (u *UserController) SignOut(c *gin.Context) {
	getHeader := c.GetHeader("Authorization")
	if getHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"status":  "error",
			"message": "token not found.",
		})
		return
	}
	split := strings.Split(getHeader, "Bearer ")
	errInvalidToken := errors.New("invalid token")
	if len(split) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"status":  "error",
			"message": errInvalidToken.Error(),
		})
	}
	getToken := split[1]
	err := u.service.SignOut(getToken)
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
		"message": "signout success",
	})
}
