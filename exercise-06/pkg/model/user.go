package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Cart     []Cart    `json:"cart" gorm:"foreignKey:UserID"`
}

type UserRepo interface {
	CreateUser(user User) (*User, error)
	CheckUser(email string, password string) (*User, error)
}

type pgUser struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &pgUser{
		DB: db,
	}
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) GenerateId() {
	user.ID = uuid.New()
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (p pgUser) CreateUser(user User) (*User, error) {
	user.GenerateId()

	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	return &user, p.DB.Create(&user).Error
}

func (p pgUser) CheckUser(email string, password string) (*User, error) {
	user := User{}
	err := p.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = user.CheckPassword(password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
