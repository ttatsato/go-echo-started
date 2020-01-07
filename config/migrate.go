package config

import (
	"app/domain"
	"github.com/jinzhu/gorm"
	"log"
)

func Migrate() (*gorm.DB, error) {
	conn, err := ConnectMySql()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.AutoMigrate(&domain.Product{}, &domain.Order{})
	log.Println("Migration has been processed")
	return conn, nil
}
