package grpc

import "github.com/beuus39/product/pkg/grpc"

type ProductGrpcClient interface {
	FindByID(id int) <-chan grpc.ServiceResult
	FindByCategory(categoryID int) <-chan grpc.ServiceResult
	FindAll() <-chan grpc.ServiceResult
}
