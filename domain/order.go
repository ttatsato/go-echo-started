package domain
type OrderSet struct {
	OrderSet []Order
}
type Order struct {
	GormModel
	UserId int `json:"userId"`
	ShopId int `json:"shopId"`
	Code string `json:"code"`
	Product Product `gorm:"foreignkey:ProductRefer" json:"product"`
	ProductRefer uint `json:"productRefer"`
	Memo string `json:"memo"`
	OrderStatus OrderStatusMaster `gorm:"foreignkey:OrderStatusRefer" json:"orderStatus"`
	OrderStatusRefer uint `gorm:"defualt:1" json:"orderStatusRefer"`
}
