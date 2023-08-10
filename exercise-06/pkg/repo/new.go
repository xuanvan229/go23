package repo

import (
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"gorm.io/gorm"
)

type Repo interface {
	DB() *gorm.DB
	AutoMigrate(models ...interface{}) error
	Product() model.ProductRepo
	User() model.UserRepo
	Cart() model.CartRepo
}
