package main

import (
	"fmt"
	grpcService "github.com/beuus39/order/internal/adapters/grpc"
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/ports/rest"
	middleware "github.com/beuus39/product/pkg/logs/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	productGrpcService, err := grpcService.NewProductGrpcClient("localhost:3002")
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}
	orderApp := app.NewOrderImpl(productGrpcService)
	orderHttpHandler := rest.NewHttpOrderHandler(orderApp)

	r := mux.NewRouter()

	r.Handle("/api/products", middleware.LogRequest(orderHttpHandler.GetProducts())).Methods("GET")
	r.Handle("/api/products/{id}", middleware.LogRequest(orderHttpHandler.GetProduct())).Methods("GET")

	log.Fatal(http.ListenAndServe(":3004", r))
}
