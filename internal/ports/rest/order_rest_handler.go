package rest

import (
	encodingJson "encoding/json"
	"fmt"
	"github.com/beuus39/order/internal/adapters/queue"
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/common/json"
	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/order/internal/shared/dtos"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// HttpOrderHandler model
type HttpOrderHandler struct {
	orderApp   app.OrderApp
	subscriber queue.ProductSubscriber
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpOrderHandler(orderApp app.OrderApp, subscriber queue.ProductSubscriber) *HttpOrderHandler {
	return &HttpOrderHandler{
		orderApp: orderApp,
		subscriber: subscriber,
	}
}

func (h *HttpOrderHandler) FindOrderById() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			json.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		h.subscriber.SubscriberProduct("product")

		paths := mux.Vars(req)
		orderId, _ := strconv.Atoi(paths["id"])
		order, err := h.orderApp.FindOrderById(uint(orderId))

		if err != nil {
			json.JsonResponse(res, "Cannot Get Order", http.StatusInternalServerError)
			return
		}

		orderDto := &dtos.FindOrdersResponse{ID: order.ID, OrderStatus: order.OrderStatus}
		json.JsonResponse(res, orderDto, http.StatusOK)
	})
}
func (h *HttpOrderHandler) FindAllOrders() http.Handler  {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			json.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		orders, err := h.orderApp.FindAllOrders()
		if err != nil {
			json.JsonResponse(res, "Cannot Get Orders", http.StatusInternalServerError)
			return
		}

		ordersDto := orders
		json.JsonResponse(res, ordersDto, http.StatusOK)
	})

}
func (h *HttpOrderHandler) SaveOrder() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			json.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		var orderDto dtos.CreateOrderDto
		err := encodingJson.NewDecoder(req.Body).Decode(&orderDto)
		if err != nil {
			json.JsonResponse(res, "Could not receive body", http.StatusBadRequest)
			return
		}

		order := domain.Order{
			OrderStatus: orderDto.OrderStatus,
		}

		isSuccess := h.orderApp.SaveOrder(order)
		json.JsonResponse(res, isSuccess, http.StatusOK)
	})
}

// GetProduct http handler function, for get product by ID
func (h *HttpOrderHandler) GetProduct() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			json.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("MemberId")

		fmt.Println(memberID)

		paths := mux.Vars(req)
		productIDStr := paths["id"]
		productID, _ := strconv.Atoi(productIDStr)

		productResult := <-h.orderApp.FindProductByID(productID)

		if productResult.Error != nil {
			log.Printf("Error get Product = %s", productResult.Error.Error())
			json.JsonResponse(res, "Cannot Get Product", http.StatusInternalServerError)
			return
		}

		product, ok := productResult.Result.(domain.Product)

		if !ok {
			json.JsonResponse(res, "Result is not product", http.StatusInternalServerError)
			return
		}

		json.JsonResponse(res, product, http.StatusOK)

	})
}

// GetProducts http handler function, for get all products
func (h *HttpOrderHandler) GetProducts() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			json.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		productResult := <-h.orderApp.FindProductAll()

		if productResult.Error != nil {
			log.Printf("Error get Product = %s", productResult.Error.Error())
			json.JsonResponse(res, "Cannot Get Products", http.StatusInternalServerError)
			return
		}

		products, ok := productResult.Result.(domain.Products)

		if !ok {
			json.JsonResponse(res, "Result is not products", http.StatusInternalServerError)
			return
		}

		json.JsonResponse(res, products, http.StatusOK)

	})
}
