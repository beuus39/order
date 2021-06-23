package app

import (
	grpcService "github.com/beuus39/order/internal/adapters/grpc"
	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/order/internal/repository"
	"github.com/beuus39/product/pkg/grpc"
	"github.com/pkg/errors"
)
type orderImpl struct {
	productGrpcService  grpcService.ProductGrpcClient
	orderRepository repository.Order
}

func (u *orderImpl) FindOrderById(orderId uint) (order domain.Order, err error) {
	return u.orderRepository.FindById(orderId)
}

func (u *orderImpl) FindAllOrders() (orders []domain.Order, err error) {
	return u.orderRepository.FindAll()
}

func (u *orderImpl) SaveOrder(order domain.Order) (isSuccess bool) {
	return u.orderRepository.SaveOrder(order)
}

//FindProductByID
func (u *orderImpl) FindProductByID(id int) <-chan grpc.ServiceResult {
	output := make(chan grpc.ServiceResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindByID(id)

		if productResult.Error != nil {
			output <- grpc.ServiceResult{Error: productResult.Error}
			return
		}

		product, ok := productResult.Result.(domain.Product)

		if !ok {
			err := errors.New("Result is not Product")
			output <- grpc.ServiceResult{Error: err}
			return
		}

		output <- grpc.ServiceResult{Result: product}
	}()
	return output
}

//FindProductAll
func (u *orderImpl) FindProductAll() <-chan grpc.ServiceResult {
	output := make(chan grpc.ServiceResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindAll()

		if productResult.Error != nil {
			output <- grpc.ServiceResult{Error: productResult.Error}
			return
		}

		products, ok := productResult.Result.(domain.Products)

		if !ok {
			err := errors.New("Result is not Products")
			output <- grpc.ServiceResult{Error: err}
			return
		}

		output <- grpc.ServiceResult{Result: products}
	}()
	return output
}

//NewOrderUseCase
func NewOrderImpl(productGrpcService grpcService.ProductGrpcClient, orderRepository repository.Order) OrderApp {
	return &orderImpl{
		productGrpcService: productGrpcService,
		orderRepository: orderRepository,
	}
}
