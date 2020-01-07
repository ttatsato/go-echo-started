package domain

type Product struct {
	GormModel
	ShopId int `json:"shopId"`
	Name string `gorm:"default:'product_name'" json:"name"`
	Code string `json:"code"`
	Price int `json:"price"`
}
