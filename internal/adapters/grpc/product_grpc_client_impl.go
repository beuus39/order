package grpc

import (
	"github.com/beuus39/order/internal/domain"
	pkgGrpc "github.com/beuus39/product/pkg/grpc"
	pb "github.com/beuus39/product/pkg/grpc/product"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"strconv"
	"time"
)
type productGrpcClientImpl struct {
	client pb.ProductServiceClient
}

//NewProductGrpcClient
func NewProductGrpcClient(host string) (*productGrpcClientImpl, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewProductServiceClient(conn)

	return &productGrpcClientImpl{
		client:      client,
	}, nil
}

//FindByID function, for find Product By Id using GRPC service client
func (c *productGrpcClientImpl) FindByID(id int) <-chan pkgGrpc.ServiceResult {
	output := make(chan pkgGrpc.ServiceResult)

	go func() {
		defer close(output)

		ctx := metadata.NewOutgoingContext(context.Background(), nil)
		arg := &pb.ProductQueryRequest{ID: int32(id)}
		res, err := c.client.FindByID(ctx, arg)

		if err != nil {
			output <- pkgGrpc.ServiceResult{Error: err}
			return
		}

		//stock from int32 to decimal
		stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

		if err != nil {
			output <- pkgGrpc.ServiceResult{Error: err}
			return
		}

		//price from float64 to decimal
		price := decimal.NewFromFloat(res.Price)

		product := domain.Product{
			ID:          int(res.ID),
			CategoryID:  int(res.CategoryID),
			Name:        res.Name,
			Description: res.Description,
			Image:       res.Image,
			Stock:       stock,
			Price:       price,
			Version:     1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		output <- pkgGrpc.ServiceResult{Result: product}
	}()

	return output
}

//FindByCategory function, for find Product By Category Id using GRPC service client
func (c *productGrpcClientImpl) FindByCategory(categoryID int) <-chan pkgGrpc.ServiceResult {
	output := make(chan pkgGrpc.ServiceResult)

	go func() {
		ctx := metadata.NewOutgoingContext(context.Background(), nil)
		arg := &pb.ProductQueryRequest{CategoryID: int32(categoryID)}
		resStream, err := c.client.FindByCategory(ctx, arg)

		if err != nil {
			output <- pkgGrpc.ServiceResult{Error: err}
			return
		}

		var products domain.Products

		for {
			res, err := resStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				output <- pkgGrpc.ServiceResult{Error: err}
				return
			}

			//stock from int32 to decimal
			stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

			if err != nil {
				output <- pkgGrpc.ServiceResult{Error: err}
				return
			}

			//price from float64 to decimal
			price := decimal.NewFromFloat(res.Price)

			product := domain.Product{
				ID:          int(res.ID),
				CategoryID:  int(res.CategoryID),
				Name:        res.Name,
				Description: res.Description,
				Image:       res.Image,
				Stock:       stock,
				Price:       price,
				Version:     1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			products = append(products, product)
		}

		output <- pkgGrpc.ServiceResult{Result: products}

	}()
	return output
}

//FindAll function, for find all Product Using GRPC service client
func (c *productGrpcClientImpl) FindAll() <-chan pkgGrpc.ServiceResult {
	output := make(chan pkgGrpc.ServiceResult)

	go func() {
		ctx := metadata.NewOutgoingContext(context.Background(), nil)
		arg := &pb.ProductQueryRequest{}
		resStream, err := c.client.FindAll(ctx, arg)

		if err != nil {
			output <- pkgGrpc.ServiceResult{Error: err}
			return
		}

		var products domain.Products

		for {
			res, err := resStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				output <- pkgGrpc.ServiceResult{Error: err}
				return
			}

			//stock from int32 to decimal
			stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

			if err != nil {
				output <- pkgGrpc.ServiceResult{Error: err}
				return
			}

			//price from float64 to decimal
			price := decimal.NewFromFloat(res.Price)

			product := domain.Product{
				ID:          int(res.ID),
				CategoryID:  int(res.CategoryID),
				Name:        res.Name,
				Description: res.Description,
				Image:       res.Image,
				Stock:       stock,
				Price:       price,
				Version:     1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			products = append(products, product)
		}

		output <- pkgGrpc.ServiceResult{Result: products}

	}()
	return output
}
