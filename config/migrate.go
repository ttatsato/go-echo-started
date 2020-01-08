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

	conn.AutoMigrate(&domain.Product{}, &domain.Order{}, &domain.Shop{}, &domain.OrderStatusMaster{})
	conn.Create( &domain.Shop{Name: "テストショップ 本店", Domain: "test-shop-honten"})
	conn.Save( &domain.OrderStatusMaster{Id: 1, Name: "未提供"})
	conn.Save( &domain.OrderStatusMaster{Id: 2, Name: "提供済み"})
	conn.Save( &domain.OrderStatusMaster{Id: 3, Name: "キャンセル申請中"})
	conn.Save( &domain.OrderStatusMaster{Id: 4, Name: "返金処理済み"})
	log.Println("Migration has been processed")
	return conn, nil
}
