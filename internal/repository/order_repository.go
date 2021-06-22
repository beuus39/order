package repository

import "github.com/beuus39/order/internal/domain"

type Order interface {
	FindAllOrders(userId uint,
		page, pageSize int) (orders []domain.Order, totalOrderCount int, err error)
	FindOrderById(orderId uint) (order domain.Order, err error)
}
