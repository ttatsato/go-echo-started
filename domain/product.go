package domain

type Product struct {
	GormModel
	ShopId int `json:"shopId"`
	Name string `json:"name"`
	Code string `json:"code"`
	Price int `json:"price"`
}
