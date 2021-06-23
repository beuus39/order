package dtos

type CreateOrderDto struct {
	ID uint `json:"id"`
	OrderStatus int `json:"order_status"`
}
