package domain

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderStatus int `gorm:"default:0"`
	TrackingNumber string
	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`

	AddressId uint

	UserId uint `gorm:"default:null"`
	OrderItemCount int `gorm:"-"`
}

func (order *Order) GetOrderStatusAsString() string  {
	switch order.OrderStatus {
	case 0:
		return "PROCESSED"
	case 1:
		return "DELIVERED"
	case 2:
		return "SHIPPED"
	default:
		return "UNKNOWN"
	}
}
