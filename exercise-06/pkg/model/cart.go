package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UserID    uuid.UUID `json:"user_id"`
}

type DeleteCart struct {
	ProductID uuid.UUID `json:"product_id"`
}

type CartRepo interface {
	GetCarts(userID uuid.UUID) ([]Cart, error)
	AddToCart(carts []Cart) (*[]Cart, error)
	CheckOut(userID uuid.UUID) error
	RemoveFromCart(ids []uuid.UUID, userID uuid.UUID) error
	GetCartsDetail(userID uuid.UUID) ([]CheckOutCart, error)
	RemoveAllCart(userID uuid.UUID) error
}

type pgCart struct {
	DB *gorm.DB
}

type CheckOutCart struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Total     int       `json:"total"`
}

func (p pgCart) GetCarts(userID uuid.UUID) ([]Cart, error) {
	rs := []Cart{}
	return rs, p.DB.Where("user_id = ?", userID).Find(&rs).Error
}

func (p pgCart) GetCartsDetail(userID uuid.UUID) ([]CheckOutCart, error) {
	rows, err := p.DB.Table("carts").Where("user_id = ?", userID).Select("carts.product_id, carts.quantity, products.name, products.price").Joins("left join products on products.id = carts.product_id").Rows()

	if err != nil {
		return nil, err
	}

	rs := []CheckOutCart{}
	for rows.Next() {
		p.DB.ScanRows(rows, &rs)
	}

	return rs, nil

}

func (p pgCart) AddToCart(cart []Cart) (*[]Cart, error) {
	return &cart, p.DB.Create(&cart).Error
}

func (p pgCart) CheckOut(userID uuid.UUID) error {
	return p.DB.Where("user_id = ?", userID).Delete(&Cart{}).Error
}

func (p pgCart) RemoveFromCart(ids []uuid.UUID, userID uuid.UUID) error {
	return p.DB.Where("product_id IN ? AND user_id = ?", ids, userID).Delete(&Cart{}).Error
}

func (p pgCart) RemoveAllCart(userID uuid.UUID) error {
	return p.DB.Where("user_id = ?", userID).Delete(&Cart{}).Error
}

func NewCartRepo(db *gorm.DB) CartRepo {
	return &pgCart{
		DB: db,
	}
}
