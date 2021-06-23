package app

import (
	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/product/pkg/grpc"
)
type OrderApp interface {
	FindProductByID(id int) <-chan grpc.ServiceResult
	FindProductAll() <-chan grpc.ServiceResult

	FindOrderById(orderId uint) (order domain.Order, err error)
	FindAllOrders() (orders []domain.Order, err error)
	SaveOrder(order domain.Order) (isSuccess bool)
}
