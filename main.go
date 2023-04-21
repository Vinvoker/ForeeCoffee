package main

import (
	"foreecoffee/controllers"
	"log"
	"os"

	docs "foreecoffee/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("ROUTER_PORT")

	// EMAIL
	router.POST("/cron", controllers.AuthMiddleware("ADMIN"), controllers.ActivateCRON)
	// @BasePath /api/v1

	// PingExample godoc
	// @Summary send email
	// @Description sending email to investors
	// @Router /email [post]
	router.POST("/email", controllers.AuthMiddleware("ADMIN"), controllers.SendEmail)

	// ORDERS
	router.POST("/order", controllers.AuthMiddleware("CUSTOMER"), controllers.InsertOrder)

	// LOGIN
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.POST("/signup", controllers.Signup)

	// PRODUCTS
	productsRoutes := router.Group("/products")
	productsRoutes.GET("", controllers.AuthMiddleware("ADMIN"), controllers.GetAllProductsAndTheirBranches)
	productsRoutes.GET("/branch", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetAllProductsByBranch)
	productsRoutes.GET("/name", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductByNameAndBranch)
	productsRoutes.GET("/coffee", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductsCoffeeByBranch)
	productsRoutes.GET("/tea", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductsTeaByBranch)
	productsRoutes.GET("/yakult", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductsYakultByBranch)

	productsRoutes.POST("", controllers.AuthMiddleware("ADMIN"), controllers.InsertProduct)
	productsRoutes.PUT("/:id", controllers.AuthMiddleware("ADMIN"), controllers.UpdateProduct)
	productsRoutes.DELETE("/:id", controllers.AuthMiddleware("ADMIN"), controllers.DeleteProduct)

	// BRANCHES
	branchesRoutes := router.Group("/branches")
	branchesRoutes.GET("", controllers.AuthMiddleware("ADMIN"), controllers.GetAllBranches)
	branchesRoutes.POST("", controllers.AuthMiddleware("ADMIN"), controllers.InsertBranch)
	branchesRoutes.PUT("/:id", controllers.AuthMiddleware("ADMIN"), controllers.UpdateBranch)
	branchesRoutes.DELETE("/:id", controllers.AuthMiddleware("ADMIN"), controllers.DeleteBranch)

	// CUSTOMER
	customerRoutes := router.Group("/customer")
	customerRoutes.PUT("/", controllers.AuthMiddleware("CUSTOMER"), controllers.UpdateCustomerProfile)
	customerRoutes.PUT("/password", controllers.AuthMiddleware("CUSTOMER"), controllers.UpdateCustomerPassword)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
