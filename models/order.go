package models

import (
	"app/db"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserId int `gorm:"default: 1"`
	ShopId int `gorm:"default: 1"`
	Product Product `gorm:"foreignkey:ProductRefer"`
	ProductRefer uint
	Memo string
}

func CreateNewOrder (insertRecord *Order) error {
	mysqlConnection := db.ConnectMySql()
	tx := mysqlConnection.Begin()

	if err := tx.Create(insertRecord).Error; err != nil {
		tx.Rollback()
		return err
	}

	// nilを返すとコミットされます
	return tx.Commit().Error
}