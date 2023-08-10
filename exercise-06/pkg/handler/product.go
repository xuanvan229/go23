package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"github.com/xuanvan229/go23/exercise-06/pkg/repo"
)

func GetProducts(c *gin.Context) {
	products, _ := repo.GetProductList()
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var product model.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
	}

	product.ID = uuid.New()

	newProduct, err := repo.CreateNewProduct(product)
	c.JSON(200, newProduct)
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	err := c.BindJSON(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))

	product.ID = id

	updatedProduct, err := repo.UpdateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}
	c.JSON(200, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	parseID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid id",
		})
	}

	err = repo.DeleteProduct(parseID)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid id",
		})
	}
	c.JSON(200, gin.H{
		"message": "Delete success",
	})
}
