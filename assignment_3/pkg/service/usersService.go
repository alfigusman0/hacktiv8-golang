package service

import (
	"assignment_3/pkg/helpers"
	"assignment_3/pkg/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UsersService struct {
	db *gorm.DB
}

func NewUsersService(db *gorm.DB) *UsersService {
	return &UsersService{db}
}

func (us *UsersService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := us.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UsersService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := us.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UsersService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	var count int64
	now := time.Now()
	// validation for unique username
	us.db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf("username %s already exists", req.Username)
	}
	// check same password and confirm password
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password not match")
	}
	// hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Nama:        req.Nama,
		Username:    req.Username,
		Password:    hashedPassword,
		DateCreated: now,
		DateUpdated: now,
	}
	if err := us.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UsersService) UpdateUser(userID uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := us.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	// check same password and confirm password
	if req.Password != nil && req.ConfirmPassword != nil {
		if *req.Password != *req.ConfirmPassword {
			return nil, fmt.Errorf("password and confirm password not match")
		}
	}
	// hash password
	if req.Password != nil {
		hashedPassword, err := helpers.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}
	// update user
	user.Nama = req.Nama
	user.Username = *req.Username
	user.DateUpdated = time.Now()
	if err := us.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UsersService) DeleteUser(userID uint) error {
	user, err := us.GetUserByID(userID)
	if err != nil {
		return err
	}
	if err := us.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// sign-in function and generate token jwt
func (us *UsersService) SignIn(req models.SignInRequest, Header string, IpAddr string) (string, error) {
	var user models.User
	now := time.Now()

	if err := us.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return "", fmt.Errorf("username %s not found", req.Username)
	}
	if !helpers.ComparePassword(user.Password, req.Password) {
		return "", fmt.Errorf("invalid password")
	}
	token, err := helpers.GenerateToken(user.UserID, user.Nama, user.Username)
	if err != nil {
		return "", err
	}

	jwt := models.Jwt{
		Header:      Header,
		IpAddr:      IpAddr,
		Token:       token,
		ExpireAt:    now.Add(time.Hour * 24),
		Expired:     "TIDAK",
		Keterangan:  "LOGIN",
		DateCreated: now,
		DateUpdated: now,
	}
	if err := us.db.Create(&jwt).Error; err != nil {
		return "", err
	}

	return token, nil
}

// sign-out function
func (us *UsersService) SignOut(token string) error {
	var jwt models.Jwt
	now := time.Now()

	if err := us.db.Where("token = ? AND expired = ?", token, "TIDAK").First(&jwt).Error; err != nil {
		return fmt.Errorf("token not found")
	}
	jwt.Expired = "YA"
	jwt.Keterangan = "LOGOUT"
	jwt.DateUpdated = now
	if err := us.db.Save(&jwt).Error; err != nil {
		return err
	}
	return nil
}
