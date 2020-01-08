package persistence

import (
	"app/domain"
	"app/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
)

// OrderRepositoryImpl Implements repository.OrderRepository
type OrderRepositoryImpl struct {
	Conn *gorm.DB
}

// OrderRepositoryWithRDB returns initialized OrderRepositoryImpl
func OrderRepositoryWithRDB(conn *gorm.DB) repository.OrderRepository {
	return &OrderRepositoryImpl{Conn: conn}
}

// Get Order by id return domain.Order
func (r *OrderRepositoryImpl) Get(id int) (*domain.Order, error) {
	Order := &domain.Order{}
	if err := r.Conn.Preload("Topic").First(&Order, id).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

func (r *OrderRepositoryImpl) GetByShopId(shopId int) ([]domain.Order, error) {
	Order := []domain.Order{}
	if err := r.Conn.Preload("Product").Where("shop_id = ?", shopId).Order("created_at desc").Find(&Order).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

func (r *OrderRepositoryImpl) GetByCustomerIdSortByCreatedAt(customerId int) ([]domain.Order, error) {
	Order := []domain.Order{}
	if err := r.Conn.Preload("Product").Where("user_id = ?", customerId).Order("created_at desc").Find(&Order).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

// GetAll Order return all domain.Order
func (r *OrderRepositoryImpl) GetAll() ([]domain.Order, error) {
	Order := []domain.Order{}
	if err := r.Conn.Preload("Topic").Find(&Order).Error; err != nil {
		return nil, err
	}

	return Order, nil
}

// Save to add Order
func (r *OrderRepositoryImpl) Save(Order *domain.Order) error {
	if err := r.Conn.Save(&Order).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Remove to delete Order by id
func (r *OrderRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	Order := domain.Order{}
	if err := tx.First(&Order, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	Order.Status = "deleted"
	if err := tx.Save(&Order).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&Order).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//// Update is update Order
//func (r *OrderRepositoryImpl) Update(Order *domain.Order) error {
//	if err := r.Conn.Model(&Order).UpdateColumns(domain.Order{Title: Order.Title, Slug: Order.Slug, Content: Order.Content, Status: Order.Status, Topic: Order.Topic}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

// GetAll Order return all domain.Order
func (r *OrderRepositoryImpl) GetAllByStatus(status string) ([]domain.Order, error) {
	if status == "deleted" {
		Order := []domain.Order{}
		if err := r.Conn.Unscoped().Where("status = ?", status).Preload("Topic").Find(&Order).Error; err != nil {
			return nil, err
		}

		return Order, nil
	}

	Order := []domain.Order{}
	if err := r.Conn.Where("status = ?", status).Preload("Topic").Find(&Order).Error; err != nil {
		return nil, err
	}

	return Order, nil
}