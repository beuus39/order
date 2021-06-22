package domain

import "github.com/jinzhu/gorm"

type OrderItem struct {
	gorm.Model
	Order Order
	OrderId uint `gorm:"not null"`

	ProductId uint `gorm:"not null"`

	Slug string `gorm:"not null"`
	ProductName string `gorm:"not null"`
	Price int `gorm:"not null"`
	Quantity int `gorm:"not null"`

	UserId uint `gorm:"default:null"`
}
