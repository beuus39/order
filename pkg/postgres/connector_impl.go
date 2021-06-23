package postgres

import (
	"fmt"
	"github.com/beuus39/order/internal/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"path"
)

type Config struct {
	Username string
	Password string
	DbName string
	Host string
	Port string
}

var DB *gorm.DB
func (c Config) Connection() *gorm.DB {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable ",
		c.Host, c.Port, c.Username, c.DbName, c.Password)
	db, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Println("Connect err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	DB = db
	DB.AutoMigrate(&domain.Order{})
	return DB
}

func (c Config) Remove(db *gorm.DB) error {
	db.Close()
	err := os.Remove(path.Join(".", "app.db"))
	return err
}

func (c Config) Get() *gorm.DB {
	return DB
}

func NewDriver(conf Config) Driver {
	return &Config{
		Username: conf.Username,
		Password: conf.Password,
		DbName: conf.DbName,
		Host: conf.Host,
		Port: conf.Port,
	}
}