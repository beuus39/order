package app

import "github.com/beuus39/product/pkg/grpc"
type OrderApp interface {
	FindProductByID(id int) <-chan grpc.ServiceResult
	FindProductAll() <-chan grpc.ServiceResult
}
