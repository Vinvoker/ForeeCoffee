package main

import (
	"foreecoffee/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("ROUTER_PORT")

	// EMAIL
	router.GET("/restart-email-cron", controllers.AuthMiddleware("ADMIN"), controllers.StartCRON)
	router.GET("/email", controllers.AuthMiddleware("ADMIN"), controllers.SendEmail)

	// ORDERS
	ordersRoutes := router.Group("/orders")
	ordersRoutes.POST("", controllers.AuthMiddleware("CUSTOMER"), controllers.InsertOrder)
	ordersRoutes.GET("/history", controllers.AuthMiddleware("CUSTOMER"), controllers.HistoryOrder)
	ordersRoutes.PUT("/status/:id", controllers.AuthMiddleware("ADMIN"), controllers.UpdateOrderStatus)

	// LOGIN
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.AuthMiddleware("CUSTOMER", "ADMIN", "INVESTOR"), controllers.Logout)
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
	customerRoutes := router.Group("/customers")
	customerRoutes.PUT("/", controllers.AuthMiddleware("CUSTOMER"), controllers.UpdateCustomerProfile)
	customerRoutes.PUT("/password", controllers.AuthMiddleware("CUSTOMER"), controllers.UpdateCustomerPassword)

	// Product Branch
	productBranchRoutes := router.Group("/productBranch")
	productBranchRoutes.POST("/:branchName", controllers.AuthMiddleware("ADMIN"), controllers.InsertMenuBranch)
	productBranchRoutes.PUT("/:branchName", controllers.AuthMiddleware("ADMIN"), controllers.UpdateMenuBranch)
	productBranchRoutes.DELETE("/:branchName", controllers.AuthMiddleware("ADMIN"), controllers.DeleteMenuBranch)

	controllers.StartCRON(nil)

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
