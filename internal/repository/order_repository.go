package repository

import "github.com/beuus39/order/internal/domain"

type Order interface {
	SaveOrder(order domain.Order) (isSuccess bool)
	FindAll() (orders []domain.Order, err error)
	FindById(orderId uint) (order domain.Order, err error)
}
