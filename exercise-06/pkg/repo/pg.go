package repo

import (
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var rp Repo

type pgRepo struct {
	db          *gorm.DB
	productRepo model.ProductRepo
	userRepo    model.UserRepo
	cartRepo    model.CartRepo
}

func (p pgRepo) DB() *gorm.DB {
	return p.db
}

func (p pgRepo) AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := p.db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (p pgRepo) Product() model.ProductRepo {
	return p.productRepo
}

func (p pgRepo) User() model.UserRepo {
	return p.userRepo
}

func (p pgRepo) Cart() model.CartRepo {
	return p.cartRepo
}

// NewPGRepo returns a new instance of a PGRepo.
func NewPGRepo(connectionStr string) Repo {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &pgRepo{
		db:          db,
		productRepo: model.NewProductRepo(db),
		userRepo:    model.NewUserRepo(db),
		cartRepo:    model.NewCartRepo(db),
	}
}

// SetupRepo set up a new instance of a Repo.
func SetupRepo() {
	rp = NewPGRepo("postgres://postgres:postgres@localhost:5435/postgres?sslmode=disable")
	rp.AutoMigrate(model.Product{}, model.User{}, model.Cart{})
}
