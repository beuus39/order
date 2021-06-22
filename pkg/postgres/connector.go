package postgres

import (
	"fmt"
	"github.com/beuus39/order/internal/config"
	"github.com/jinzhu/gorm"
	"os"
	"path"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func OpenDbConnection() *gorm.DB {
	conf, _ := config.NewConfig("C:/Users/ADMIN/go/src/BeU/BeUUS/order/internal/config/application.yml")

	//dialect := conf.Database.Dialect
	username := conf.Database.DBUser
	password := conf.Database.Password
	dbName := conf.Database.DBName
	host := conf.Database.Host

	var db *gorm.DB
	var err error

	databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)
	db, err = gorm.Open(databaseUrl)

	if err != nil {
		fmt.Println("DB err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	DB = db
	return DB
}

func RemoveDb(db *gorm.DB) error {
	db.Close()
	err := os.Remove(path.Join(".", "app.db"))
	return err
}

func GetDb() *gorm.DB {
	return DB
}
