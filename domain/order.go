package domain
type OrderSet struct {
	OrderSet []Order
}
type Order struct {
	GormModel
	UserId int `json:"userId"`
	ShopId int `json:"shopId"`
	Product Product `gorm:"foreignkey:ProductRefer" json:"product"`
	ProductRefer uint `json:"productRefer"`
	Memo string `json:"memo"`
	Status string `json: "status"`
}
