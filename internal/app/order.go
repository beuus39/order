package app

import (
	grpcService "github.com/beuus39/order/internal/adapters/grpc"
	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/product/pkg/grpc"
	"github.com/pkg/errors"
)
type orderImpl struct {
	productGrpcService  grpcService.ProductGrpcClient
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
func NewOrderImpl(productGrpcService grpcService.ProductGrpcClient) OrderApp {
	return &orderImpl{
		productGrpcService:    productGrpcService,
	}
}
