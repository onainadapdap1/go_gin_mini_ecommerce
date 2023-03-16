package handler

import (
	"go_gin_mini_ecommerce/models"
	"go_gin_mini_ecommerce/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// interface to user entity
type UserHandler interface {
	GetUser(*gin.Context) 
	AddUser(*gin.Context)
}

// depend on UserRepository
type userHandler struct {
	// variable ini menampung objek dari struct userRepository
	repo repository.UserRepository // var repo tipe interface UserRepository
	// repo = userRepository{db: DB()}
}

// return new UserHandler
func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hashPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hashPass)
}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get user
	user, err := h.repo.GetUser(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": user,
	})
}

func (h *userHandler) AddUser(c *gin.Context) {
	// binding input request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input is wrong",
		})
		return
	}

	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "this one is wrong",
		})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": user,
	})
}