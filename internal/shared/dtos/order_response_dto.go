package dtos

type FindOrdersResponse struct {
	ID uint `json:"id"`
	OrderStatus int `json:"order_status"`
}
