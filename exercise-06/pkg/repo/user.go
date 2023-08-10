package repo

import (
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
)

func Register(user model.User) (*model.User, error) {
	return rp.User().CreateUser(user)
}

func Login(email string, password string) (*model.User, error) {
	return rp.User().CheckUser(email, password)
}
