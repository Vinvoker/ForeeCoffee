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

	router.POST("/cron", controllers.ActivateCRON)
	router.POST("/email", controllers.SendEmail)

	// PRODUCTS
	router.GET("/products", controllers.GetAllProducts)                 // autentikasi : admin, customer
	router.GET("/products-coffee", controllers.GetProductsCoffee)       // autentikasi : admin, customer  // tidak ada input
	router.GET("/products-noncoffee", controllers.GetProductsNonCoffee) // autentikasi : admin, customer  // tidak ada input
	router.GET("/product", controllers.GetProduct)                      // autentikasi : admin, customer  // input : nama product
	router.POST("/product", controllers.InsertProduct)                  // autentikasi : admin
	router.PUT("/product", controllers.UpdateProduct)                   // autentikasi : admin
	router.DELETE("/product", controllers.DeleteProduct)                // autentikasi : admin

	// ORDERS
	router.POST("/order", controllers.InsertOrder) // autentikasi : customer

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
