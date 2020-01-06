package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
type Product struct {
	gorm.Model
	Name string `gorm:"default:'product_name'"`
	Code string
	Price int
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
