package main

import (
	"fmt"
	grpcService "github.com/beuus39/order/internal/adapters/grpc"
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/config"
	"github.com/beuus39/order/internal/ports/rest"
	"github.com/beuus39/order/internal/repository"
	"github.com/beuus39/order/pkg/postgres"
	middleware "github.com/beuus39/product/pkg/logs/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	conf, _ := config.NewConfig("C:/Users/cong.nguyenthanh4/go/src/BeUUS/order/internal/config/application.yml")
	cfg := postgres.Config{
		Username: conf.Database.DBUser,
		Host: conf.Database.Host,
		DbName: conf.Database.DBName,
		Password: conf.Database.Password,
		Port: conf.Database.Port,
	}
	connector := postgres.NewDriver(cfg)
	connector.Connection()
	db := connector.Get()

	orderRepository := repository.NewOrderRepository(db)

	productGrpcService, err := grpcService.NewProductGrpcClient("localhost:3002")
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}
	orderApp := app.NewOrderImpl(productGrpcService, orderRepository)
	orderHttpHandler := rest.NewHttpOrderHandler(orderApp)

	r := mux.NewRouter()

	r.Handle("/api/products", middleware.LogRequest(orderHttpHandler.GetProducts())).Methods("GET")
	r.Handle("/api/products/{id}", middleware.LogRequest(orderHttpHandler.GetProduct())).Methods("GET")

	r.Handle("/api/orders", middleware.LogRequest(orderHttpHandler.FindAllOrders())).Methods("GET")
	r.Handle("/api/orders/{id}", middleware.LogRequest(orderHttpHandler.FindOrderById())).Methods("GET")
	r.Handle("/api/orders", middleware.LogRequest(orderHttpHandler.SaveOrder())).Methods("POST")

	log.Fatal(http.ListenAndServe(":3004", r))
}
