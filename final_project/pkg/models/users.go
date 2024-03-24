package models

import "time"

type User struct {
	UserID          uint      `json:"id" gorm:"column:id_user;primaryKey;autoIncrement"`
	Nama            string    `json:"nama" gorm:"column:nama"`
	Username        string    `json:"username" gorm:"column:username;unique;"`
	Password        string    `json:"password" gorm:"column:password"`
	ConfirmPassword string    `json:"confirm_password" gorm:"-"` // not in database
	Roles           string    `json:"roles" gorm:"column:roles"`
	DateCreated     time.Time `json:"date_created" gorm:"column:date_created"`
	DateUpdated     time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type GetAllUserRequest struct {
	UserID   uint   `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Roles    string `json:"roles"`
}

type CreateUserRequest struct {
	Nama            string `json:"nama" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Roles           string `json:"roles" binding:"required"`
}

type UpdateUserRequest struct {
	Nama            string  `json:"nama"`
	Username        *string `json:"username"`
	Password        *string `json:"password"`
	ConfirmPassword *string `json:"confirm_password"`
	Roles           *string `json:"roles"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
