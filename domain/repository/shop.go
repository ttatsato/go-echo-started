package repository

import "app/domain"

type ShopRepository interface {
	GetByShopId(shopId int) ([]domain.Shop, error)
	Save(*domain.Shop) error
	Create(*domain.Shop) *domain.Shop
	Remove(id int) error
	//Update(*domain.Shop) error
}
