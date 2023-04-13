package main

import (
	"gin/controllers"
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
	router.POST("/cron", controllers.AuthMiddleware("ADMIN"), controllers.ActivateCRON)
	router.POST("/email", controllers.AuthMiddleware("ADMIN"), controllers.SendEmail)

	// ORDERS
	router.POST("/order", controllers.AuthMiddleware("CUSTOMER"), controllers.InsertOrder)

	// LOGIN
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.POST("/signup", controllers.Signup)

	// PRODUCTS
	productsRoutes := router.Group("/products")
	productsRoutes.GET("", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetAllProducts)
	productsRoutes.GET("/:name", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProduct)
	productsRoutes.GET("/coffee", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductsCoffee)
	productsRoutes.GET("/noncoffee", controllers.AuthMiddleware("ADMIN", "CUSTOMER"), controllers.GetProductsNonCoffee)

	productsRoutes.POST("", controllers.AuthMiddleware("ADMIN"), controllers.InsertProduct)
	productsRoutes.PUT("", controllers.AuthMiddleware("ADMIN"), controllers.UpdateProduct)
	productsRoutes.DELETE("", controllers.AuthMiddleware("ADMIN"), controllers.DeleteProduct)

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
