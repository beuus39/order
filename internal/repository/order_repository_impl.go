package repository

import (
	"github.com/beuus39/order/internal/domain"
	"github.com/jinzhu/gorm"
)

type Connector struct {
	db *gorm.DB
}

func (c *Connector) FindAll() (orders []domain.Order, err error) {
	err = c.db.Find(&orders).Error
	if err != nil  {
		return nil, err
	}
	return orders, nil
}

func (c *Connector) FindById(orderId uint) (order domain.Order, err error) {
	err = c.db.Model(&domain.Order{}).First(&order, orderId).Error
	return order, err
}

func (c *Connector) SaveOrder(order domain.Order) (isSuccess bool) {
	savedOrder := c.db.Create(&order)

	if savedOrder.Error != nil {
		return false
	}
	return true
}

func NewOrderRepository(db *gorm.DB) Order {
	return &Connector{
		db: db,
	}
}