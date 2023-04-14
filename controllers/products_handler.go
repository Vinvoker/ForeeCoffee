package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products Product
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Price, &products.PictureUrl, &products.Category); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			c.IndentedJSON(http.StatusOK, products)
		}
	}
}

func GetAllProductsByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=?", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products Product
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Price, &products.PictureUrl, &products.Category); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			var response BranchProducts
			response.Data = products
			response.Branch = branchName
			c.IndentedJSON(http.StatusOK, response)
		}
	}
}

func GetProductsCoffee(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='COFFEE'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products Product
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Price, &products.PictureUrl, &products.Category); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			var response BranchProducts
			response.Data = products
			response.Branch = branchName
			c.IndentedJSON(http.StatusOK, response)
		}
	}
}

func GetProductsYakult(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='YAKULT'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products Product
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Price, &products.PictureUrl, &products.Category); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			var response BranchProducts
			response.Data = products
			response.Branch = branchName
			c.IndentedJSON(http.StatusOK, response)
		}
	}
}

func GetProductsTea(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='TEA'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products Product
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Price, &products.PictureUrl, &products.Category); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			var response BranchProducts
			response.Data = products
			response.Branch = branchName
			c.IndentedJSON(http.StatusOK, response)
		}
	}
}

func GetProduct(c *gin.Context) {

}

func InsertProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	var product Product

	if err := c.Bind(&product); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := db.Query("INSERT INTO product (name, price, pictureUrl, category) VALUES (?,?,?,?)",
		product.Name,
		product.Price,
		product.PictureUrl,
		product.Category,
	)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "insert failed"})
	} else {
		c.IndentedJSON(http.StatusCreated, product)
	}
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
