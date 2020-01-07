package persistence

import (
	"app/domain"
	"app/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
)

// ProductRepositoryImpl Implements repository.ProductRepository
type ProductRepositoryImpl struct {
	Conn *gorm.DB
}

// ProductRepositoryWithRDB returns initialized ProductRepositoryImpl
func ProductRepositoryWithRDB(conn *gorm.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{Conn: conn}
}

// Get Product by id return domain.Product
func (r *ProductRepositoryImpl) Get(id int) (*domain.Product, error) {
	Product := &domain.Product{}
	if err := r.Conn.Preload("Topic").First(&Product, id).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (r *ProductRepositoryImpl) GetByShopId(shopId int) ([]domain.Product, error) {
	Product := []domain.Product{}
	if err := r.Conn.Where("shop_id = ?", shopId).Find(&Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (r *ProductRepositoryImpl) GetByCustomerId(customerId int) ([]domain.Product, error) {
	Product := []domain.Product{}
	if err := r.Conn.Where("user_id = ?", customerId).Find(&Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

// GetAll Product return all domain.Product
func (r *ProductRepositoryImpl) GetAll() ([]domain.Product, error) {
	Product := []domain.Product{}
	if err := r.Conn.Preload("Topic").Find(&Product).Error; err != nil {
		return nil, err
	}

	return Product, nil
}

// Save to add Product
func (r *ProductRepositoryImpl) Save(Product *domain.Product) error {
	log.Println("Save()")
	if err := r.Conn.Save(&Product).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) Create(Product *domain.Product) *domain.Product {
	log.Println("Create()")
	r.Conn.Create(&Product)
	log.Println(Product)
	return Product
}

// Remove to delete Product by id
func (r *ProductRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	Product := domain.Product{}
	if err := tx.First(&Product, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&Product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&Product).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//// Update is update Product
//func (r *ProductRepositoryImpl) Update(Product *domain.Product) error {
//	if err := r.Conn.Model(&Product).UpdateColumns(domain.Product{Title: Product.Title, Slug: Product.Slug, Content: Product.Content, Status: Product.Status, Topic: Product.Topic}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

// GetAll Product return all domain.Product
func (r *ProductRepositoryImpl) GetAllByStatus(status string) ([]domain.Product, error) {
	if status == "deleted" {
		Product := []domain.Product{}
		if err := r.Conn.Unscoped().Where("status = ?", status).Preload("Topic").Find(&Product).Error; err != nil {
			return nil, err
		}

		return Product, nil
	}

	Product := []domain.Product{}
	if err := r.Conn.Where("status = ?", status).Preload("Topic").Find(&Product).Error; err != nil {
		return nil, err
	}

	return Product, nil
}