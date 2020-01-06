package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
 * gormを使って、MySqlと接続する
 * @see http://gorm.io/ja_JP/docs
 */
func ConnectMySql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/order_sass?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました")
		panic("データベースへの接続に失敗しました")
	}
	return db
}