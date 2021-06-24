package main

import (
	"github.com/beuus39/order/internal/app"
	"github.com/beuus39/order/internal/domain"
	"github.com/beuus39/testify/assert"
	"github.com/beuus39/testify/mock"
	"testing"
)

type MockOrderRepository struct {
	mock.Mock
}

func (r *MockOrderRepository) SaveOrder(order domain.Order) (isSuccess bool) {
	ret := r.Called(order)
	return ret.Get(0).(bool)
}

func (r *MockOrderRepository) FindAll() (orders []domain.Order, err error) {
	ret := r.Called()
	return ret.Get(0).([]domain.Order), ret.Error(1)
}

func (r *MockOrderRepository) FindById(orderId uint) (order domain.Order, err error) {
	ret := r.Called(orderId)
	return ret.Get(0).(domain.Order), ret.Error(1)
}

func TestFindOrderById(t *testing.T) {
	mockOrder := domain.Order{OrderStatus: 1, ID: 0}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.On("FindById", uint(0)).Return(mockOrder, nil)

	orderApp := app.NewOrderImpl(nil, mockOrderRepository)
	res, err := orderApp.FindOrderById(0)
	assert.Nil(t, err, "Error should be nill")
	assert.Equal(t, mockOrder, res, "service should return the same one as repository")
}

func TestFindAllOrders(t *testing.T) {
	mockOrders := []domain.Order{
		{OrderStatus: 1, ID: 0},
		{OrderStatus: 2, ID: 1},
		{OrderStatus: 3, ID: 2},
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.On("FindAll").Return(mockOrders, nil)

	orderApp := app.NewOrderImpl(nil, mockOrderRepository)
	res, err := orderApp.FindAllOrders()
	assert.Nil(t, err, "Error should be nill")
	assert.Equal(t, mockOrders, res, "service should return the same one as repository")
}

func TestSaveOrders(t *testing.T) {
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.On("SaveOrder", domain.Order{
		OrderStatus: 3,
		ID: uint(5),
	}).Return(true, nil)

	orderApp := app.NewOrderImpl(nil, mockOrderRepository)
	res := orderApp.SaveOrder(domain.Order{
		OrderStatus: 3,
		ID: uint(5),
	})
	assert.Equal(t, true, res, "service should return the same one as repository")
}