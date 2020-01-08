package persistence

import (
	"app/domain"
	"app/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
)

// ShopRepositoryImpl Implements repository.ShopRepository
type ShopRepositoryImpl struct {
	Conn *gorm.DB
}

// ShopRepositoryWithRDB returns initialized ShopRepositoryImpl
func ShopRepositoryWithRDB(conn *gorm.DB) repository.ShopRepository {
	return &ShopRepositoryImpl{Conn: conn}
}

func (r *ShopRepositoryImpl) GetByShopId(shopId int) domain.Shop {
	Shop := domain.Shop{}
	if err := r.Conn.Where("id = ?", shopId).First(&Shop).Error; err != nil {
		return domain.Shop{}
	}
	return Shop
}

// Save to add Shop
func (r *ShopRepositoryImpl) Save(Shop *domain.Shop) error {
	if err := r.Conn.Save(&Shop).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *ShopRepositoryImpl) Create(Shop *domain.Shop) *domain.Shop {
	r.Conn.Create(&Shop)
	return Shop
}

// Remove to delete Shop by id
func (r *ShopRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	Shop := domain.Shop{}
	if err := tx.First(&Shop, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&Shop).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&Shop).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//// Update is update Shop
//func (r *ShopRepositoryImpl) Update(Shop *domain.Shop) error {
//	if err := r.Conn.Model(&Shop).UpdateColumns(domain.Shop{Title: Shop.Title, Slug: Shop.Slug, Content: Shop.Content, Status: Shop.Status, Topic: Shop.Topic}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}