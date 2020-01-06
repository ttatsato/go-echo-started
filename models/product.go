package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
type Product struct {
	GormModel
	Name string `gorm:"default:'product_name'" json:"name"`
	Code string `json:"code"`
	Price int `json:"price"`
}

func (p *Product) BeforeSave() (err error) {
	fmt.Println("Model Product BeforeSave")
	return
}

func (p *Product) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Println("Model Product AfterCreate")
	fmt.Println("レコードを追加しました")
	return
}
