package models

import (
	"app/db"
)

type Order struct {
	GormModel
	UserId int `gorm:"default: 1" json:"userId"`
	ShopId int `gorm:"default: 1" json:"shopId"`
	Product Product `gorm:"foreignkey:ProductRefer" json:"product"`
	ProductRefer uint `json:"productRefer"`
	Memo string `json:"memo"`
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

func ReadOrder() []Order {
	var orders []Order
	mysqlConnection := db.ConnectMySql()
	mysqlConnection.Where("shop_id = ?", 133).Find(&orders)
	return orders
}