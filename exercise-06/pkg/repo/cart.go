package repo

import (
	"github.com/google/uuid"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
)

func GetCarts(userID uuid.UUID) ([]model.Cart, error) {
	return rp.Cart().GetCarts(userID)
}

func AddToCart(carts []model.Cart) (*[]model.Cart, error) {
	return rp.Cart().AddToCart(carts)
}

func CheckOut(userID uuid.UUID) error {
	return rp.Cart().CheckOut(userID)
}

func RemoveFromCart(ids []uuid.UUID, userID uuid.UUID) error {
	return rp.Cart().RemoveFromCart(ids, userID)
}

func GetCartDetails(userID uuid.UUID) ([]model.CheckOutCart, error) {
	return rp.Cart().GetCartsDetail(userID)
}

func RemoveAllCart(userID uuid.UUID) error {
	return rp.Cart().RemoveAllCart(userID)
}
