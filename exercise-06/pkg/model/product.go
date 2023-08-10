package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
	Sku   string    `json:"sku"`
	Cart  []Cart    `json:"cart" gorm:"foreignKey:ProductID"`
}

type ProductRepo interface {
	GetProducts() ([]Product, error)
	CreateProduct(product Product) (*Product, error)
	UpdateProduct(product Product) (*Product, error)
	DeleteProduct(id uuid.UUID) error
}

type pgProduct struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &pgProduct{
		DB: db,
	}
}

func (p pgProduct) GetProducts() ([]Product, error) {
	rs := []Product{}
	return rs, p.DB.Find(&rs).Error
}

func (p pgProduct) CreateProduct(product Product) (*Product, error) {
	return &product, p.DB.Create(&product).Error
}

func (p pgProduct) UpdateProduct(product Product) (*Product, error) {
	return &product, p.DB.Where("id = ?", product.ID).Updates(&product).Error
}

func (p pgProduct) DeleteProduct(id uuid.UUID) error {
	return p.DB.Where("id = ?", id).Delete(&Product{}).Error
}
