package config

import (
	"github.com/jinzhu/gorm"
	"log"
)
/**
 * gormを使って、MySqlと接続する
 * @see http://gorm.io/ja_JP/docs
 */
func ConnectMySql() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@(localhost)/order_sass?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("データベースへの接続に失敗しました")
		panic("データベースへの接続に失敗しました")
	}
	return db, err
}
