package repository

import "app/domain"

type ProductRepository interface {
	Get(id int) (*domain.Product, error)
	GetByShopId(shopId int) ([]domain.Product, error)
	GetByCustomerId(CustomerId int) ([]domain.Product, error)
	GetAll() ([]domain.Product, error)
	Save(*domain.Product) error
	Create(*domain.Product) *domain.Product
	Remove(id int) error
	//Update(*domain.Product) error
}
