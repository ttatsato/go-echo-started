package repository

import "app/domain"

type OrderRepository interface {
	Get(id int) (*domain.Order, error)
	GetByShopId(shopId int) ([]domain.Order, error)
	GetByCustomerId(CustomerId int) ([]domain.Order, error)
	GetAll() ([]domain.Order, error)
	Save(*domain.Order) error
	Remove(id int) error
	//Update(*domain.Order) error
}
