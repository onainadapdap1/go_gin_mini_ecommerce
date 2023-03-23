package handler

import (
	"fmt"
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
	SignInUser(*gin.Context)
	GetAllUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
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

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := h.repo.DeleteUser(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot delete data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
	})
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

	user.Password = ""
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

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) SignInUser(c *gin.Context) {
	// binding request input
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get user from db
	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No such user found",
		})
		return
	}

	// compare password
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token, err := GenerateToken(dbUser.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "couldn't generate token",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully signIn",
			"token": token,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error": "Password not matched",
	})
	
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	fmt.Println(c.Get("userID"))
	users, err := h.repo.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	usersWithNoPass := []models.User{}
	for _, user := range users {
		user.Password = ""
		usersWithNoPass = append(usersWithNoPass, user)

	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": usersWithNoPass,
	})
}


func (h *userHandler) UpdateUser(c *gin.Context) {
	// get param
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// binding request input
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	user, err = h.repo.UpdateUser(uint(idInt), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating user data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"new_data": user,
	})
}