package postgres

import "github.com/jinzhu/gorm"

type Driver interface {
	Connection() *gorm.DB
	Remove(db *gorm.DB) error
	Get() *gorm.DB
}