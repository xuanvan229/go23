package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"github.com/xuanvan229/go23/exercise-06/pkg/repo"
)

func GetCarts(c *gin.Context) {
	userId, err := uuid.Parse(c.MustGet("userId").(string))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}
	carts, err := repo.GetCarts(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}
	c.JSON(200, carts)
	return
}

func AddToCart(c *gin.Context) {
	var carts []model.Cart
	userId, err := uuid.Parse(c.MustGet("userId").(string))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Invalid data",
		})
		return
	}
	err = c.BindJSON(&carts)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid data",
		})
		return
	}

	for i, cart := range carts {
		cart.ID = uuid.New()
		cart.UserID = userId
		carts[i] = cart
	}

	newCart, err := repo.AddToCart(carts)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid data",
		})
		return
	}

	if err != nil {
		c.JSON(402, gin.H{
			"message": "Invalid data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": newCart,
	})
	return
}

func DeleteCart(c *gin.Context) {

	var deleteProducts []model.DeleteCart

	err := c.BindJSON(&deleteProducts)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid data",
		})
		return
	}

	userId, err := uuid.Parse(c.MustGet("userId").(string))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Invalid data",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	ids := make([]uuid.UUID, len(deleteProducts))
	for i, deleteProduct := range deleteProducts {
		ids[i] = deleteProduct.ProductID
	}

	err = repo.RemoveFromCart(ids, userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
	return

}

func CheckOut(c *gin.Context) {
	userId, err := uuid.Parse(c.MustGet("userId").(string))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}
	carts, err := repo.GetCartDetails(userId)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	result := []model.CheckOutCart{}

	// combine same product to one and add to result
	for _, cart := range carts {
		if len(result) == 0 {
			result = append(result, cart)
			continue
		}
		for i, r := range result {
			if r.ProductID == cart.ProductID {
				result[i].Quantity += cart.Quantity
				break
			}
			if i == len(result)-1 {
				result = append(result, cart)
				break
			}
		}
	}

	for i, cart := range result {
		cart.Total = cart.Quantity * cart.Price
		result[i] = cart
	}

	err = repo.RemoveAllCart(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid data",
		})
		return
	}

	c.JSON(200, result)
}
