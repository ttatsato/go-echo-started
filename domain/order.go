package domain

type Order struct {
	GormModel
	UserId int `gorm:"default: 1" json:"userId"`
	ShopId int `gorm:"default: 1" json:"shopId"`
	Product Product `gorm:"foreignkey:ProductRefer" json:"product"`
	ProductRefer uint `json:"productRefer"`
	Memo string `json:"memo"`
	Status string `json: "status"`
}
