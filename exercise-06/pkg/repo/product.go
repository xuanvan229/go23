package repo

import (
	"github.com/google/uuid"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
)

func GetProductList() ([]model.Product, error) {
	return rp.Product().GetProducts()
}

func CreateNewProduct(product model.Product) (*model.Product, error) {
	return rp.Product().CreateProduct(product)
}

func UpdateProduct(product model.Product) (*model.Product, error) {
	return rp.Product().UpdateProduct(product)
}

func DeleteProduct(id uuid.UUID) error {
	return rp.Product().DeleteProduct(id)
}
