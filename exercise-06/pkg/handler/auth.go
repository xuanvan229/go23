package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"github.com/xuanvan229/go23/exercise-06/pkg/repo"
)

type AuthHandler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var authHandler AuthHandler

	if err := c.ShouldBindJSON(&authHandler); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	user, err := repo.Login(authHandler.Email, authHandler.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	token, err := model.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
	return

}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	registerUser, err := repo.Register(user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	c.JSON(200, registerUser)

}
