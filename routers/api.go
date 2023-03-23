package routers

import (
	"go_gin_mini_ecommerce/handler"
	"go_gin_mini_ecommerce/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	// set objek UserHandler, untuk dapat 
	// memanggil method di UserHandler
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to our mini ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/:id", userHandler.GetUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/:product", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
	}
	return r.Run(address)
}