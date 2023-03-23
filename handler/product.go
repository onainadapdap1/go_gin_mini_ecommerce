package handler

import (
	"go_gin_mini_ecommerce/models"
	"go_gin_mini_ecommerce/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	AddProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{repo: repository.NewProductRepository()}
}

func (h *productHandler) GetProduct(c *gin.Context) {
	// get id product
	prodStr := c.Param("product")
	prodId, err := strconv.Atoi(prodStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get product data
	product, err := h.repo.GetProduct(prodId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// set response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func(h *productHandler) AddProduct(c *gin.Context) {
	// bind json data
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call repo to save product
	product, err := h.repo.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": product,
	})
}