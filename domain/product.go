package domain

type Product struct {
	GormModel
	Name string `gorm:"default:'product_name'" json:"name"`
	Code string `json:"code"`
	Price int `json:"price"`
}
