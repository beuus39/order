package rest

import (
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/common/json"
	"github.com/beuus39/order/internal/domain"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// HttpOrderHandler model
type HttpOrderHandler struct {
	orderApp app.OrderApp
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpOrderHandler(orderApp app.OrderApp) *HttpOrderHandler {
	return &HttpOrderHandler{
		orderApp: orderApp,
	}
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
