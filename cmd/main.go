package main

import (
	"fmt"
	"github.com/beuus39/order/internal/adapters/queue"
	"github.com/beuus39/order/pkg/nats"
	"log"
	"net/http"
	"os"

	grpcService "github.com/beuus39/order/internal/adapters/grpc"
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/config"
	"github.com/beuus39/order/internal/ports/rest"
	"github.com/beuus39/order/internal/repository"
	"github.com/beuus39/order/pkg/postgres"
	middleware "github.com/beuus39/product/pkg/logs/http"
	"github.com/gorilla/mux"
)

func main() {

	productConf := nats.NatsConfig{
		Url: "nats://localhost:4222",
	}
	productQueue := queue.NewProductSubscriber(productConf)
	conf := config.LoadEnv()

	fmt.Printf("Username = %s Host = %s DbName = %s Password = %s Port = %s",
		conf.DBUser, conf.Host, conf.DBName, conf.Password, conf.Port)
	cfg := postgres.Config{
		Username: "beu",
		Host:     conf.Host,
		DbName:   conf.DBName,
		Password: conf.Password,
		Port:     conf.Port,
	}
	connector := postgres.NewDriver(cfg)
	connector.Connection()
	db := connector.Get()

	orderRepository := repository.NewOrderRepository(db)

	productGrpcService, err := grpcService.NewProductGrpcClient(conf.GrpcClient)
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}
	orderApp := app.NewOrderImpl(productGrpcService, orderRepository)
	orderHttpHandler := rest.NewHttpOrderHandler(orderApp, productQueue)

	r := mux.NewRouter()

	r.Handle("/api/products", middleware.LogRequest(orderHttpHandler.GetProducts())).Methods("GET")
	r.Handle("/api/products/{id}", middleware.LogRequest(orderHttpHandler.GetProduct())).Methods("GET")

	r.Handle("/api/orders", middleware.LogRequest(orderHttpHandler.FindAllOrders())).Methods("GET")
	r.Handle("/api/orders/{id}", middleware.LogRequest(orderHttpHandler.FindOrderById())).Methods("GET")
	r.Handle("/api/orders", middleware.LogRequest(orderHttpHandler.SaveOrder())).Methods("POST")

	log.Fatal(http.ListenAndServe(":3004", r))
}
