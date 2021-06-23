package domain

type Order struct {
	OrderStatus int `gorm:"column:order_status;default:0"`
	ID uint `gorm:"column:id;primaryKey;autoIncrement"`
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
