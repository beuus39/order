package main

import (
	"github.com/beuus39/product/pkg/grpc"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/order/internal/ports/rest"
	"github.com/beuus39/testify/mock"
)

type MockOrderApp struct {
	mock.Mock
}

func (m MockOrderApp) FindProductByID(id int) <-chan grpc.ServiceResult {
	panic("implement me")
}

func (m MockOrderApp) FindProductAll() <-chan grpc.ServiceResult {
	panic("implement me")
}

func (m *MockOrderApp) FindOrderById(orderId uint) (order domain.Order, err error) {
	ret := m.Called(orderId)
	return ret.Get(0).(domain.Order), ret.Error(1)
}

func (m *MockOrderApp) FindAllOrders() (orders []domain.Order, err error) {
	ret := m.Called()
	return ret.Get(0).([]domain.Order), ret.Error(1)
}

func (m *MockOrderApp) SaveOrder(order domain.Order) (isSuccess bool) {
	ret := m.Called(order)
	return ret.Get(0).(bool)
}

func TestFindAllOrders(t *testing.T) {
	mockOrder := []domain.Order{
		{ OrderStatus: 1, ID: 0 },
		{ OrderStatus: 2, ID: 1 },
		{ OrderStatus: 3, ID: 3 },
	}
	mockOrderApp := new(MockOrderApp)
	mockOrderApp.On("FindAllOrders").Return(mockOrder, nil)

	orderHttpHandler := rest.NewHttpOrderHandler(mockOrderApp)
	r := mux.NewRouter()
	r.Handle("/api/orders", orderHttpHandler.FindAllOrders())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Method not allowed", func(t *testing.T) {
		apitest.
			New().
			Handler(r).
			Put("/api/orders").
			Expect(t).
			Status(http.StatusMethodNotAllowed).
			End()
	})


	t.Run("ok", func(t *testing.T) {
		apitest.
			New().
			Handler(r).
			Get("/api/orders").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
