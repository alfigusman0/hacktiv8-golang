package models

import (
	"time"
)

type Jwt struct {
	ID          uint      `json:"id" gorm:"column:ids_jwt;primaryKey;autoIncrement"`
	Header      string    `json:"header" gorm:"column:header"`
	IpAddr      string    `json:"ip_addr" gorm:"column:ip_address"`
	Token       string    `json:"token" gorm:"column:token"`
	ExpireAt    time.Time `json:"expire_at" gorm:"column:expire_at"`
	Expired     string    `json:"expired" gorm:"column:expired"`
	Keterangan  string    `json:"keterangan" gorm:"column:keterangan"`
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
}

type JwtRequest struct {
	Token string `json:"token" binding:"required"`
}

type JwtCreateRequest struct {
	Header     string    `json:"header" binding:"required"`
	IpAddr     string    `json:"ip_addr" binding:"required"`
	Token      string    `json:"token" binding:"required"`
	ExpireAt   time.Time `json:"expire_at" binding:"required"`
	Expired    string    `json:"expired" binding:"required"`
	Keterangan string    `json:"keterangan" binding:"required"`
}
