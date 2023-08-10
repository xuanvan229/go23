package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/go23/exercise-06/pkg/handler"
	"github.com/xuanvan229/go23/exercise-06/pkg/middleware"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	products := r.Group("/products")
	{
		products.GET("/", handler.GetProducts)
		products.POST("/", handler.CreateProduct)
		products.PUT("/:id", handler.UpdateProduct)
		products.DELETE("/:id", handler.DeleteProduct)
	}

	cart := r.Group("/cart")
	cart.Use(middleware.Auth())

	{
		cart.GET("/", handler.GetCarts)
		cart.POST("/add", handler.AddToCart)
		cart.DELETE("/remove", handler.DeleteCart)
		cart.POST("/checkout", handler.CheckOut)
	}

	return r
}
