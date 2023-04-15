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

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products []ProductsDetails
	var product ProductsDetails
	for rows.Next() {
		var branch Branch
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {

			// Check if the current product already exists in the products slice
			// If it does, append the branch to the existing product's Branch slice
			// If it doesn't, create a new product and append it to the products slice
			if len(products) > 0 && products[len(products)-1].ID == product.ID {
				products[len(products)-1].Branch = append(products[len(products)-1].Branch, branch)
			} else {
				product.Branch = []Branch{branch}
				products = append(products, product)
			}

		}
	}
	c.IndentedJSON(http.StatusOK, products)
}

func GetAllProductsByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=?", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product Product
	var products []Product
	var branch Branch
	var error bool
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			products = append(products, product)
		}
	}

	if !error {
		var response BranchProductsForMenu
		response.Product = products
		response.Branch = branch
		c.IndentedJSON(http.StatusOK, response)
	}
}

func GetProductsCoffeeByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='COFFEE'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product Product
	var products []Product
	var branch Branch
	var error bool
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			products = append(products, product)
		}
	}

	if !error {
		var response BranchProductsForMenu
		response.Product = products
		response.Branch = branch
		c.IndentedJSON(http.StatusOK, response)
	}
}

func GetProductsYakultByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='YAKULT'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product Product
	var products []Product
	var branch Branch
	var error bool
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			products = append(products, product)
		}
	}

	if !error {
		var response BranchProductsForMenu
		response.Product = products
		response.Branch = branch
		c.IndentedJSON(http.StatusOK, response)
	}
}

func GetProductsTeaByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	fmt.Print("branchName = ", branchName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.category='TEA'", branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product Product
	var products []Product
	var branch Branch
	var error bool
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			products = append(products, product)
		}
	}

	if !error {
		var response BranchProductsForMenu
		response.Product = products
		response.Branch = branch
		c.IndentedJSON(http.StatusOK, response)
	}
}

func GetProductByNameAndBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")
	productName := "%" + c.Query("Name") + "%"

	fmt.Print("branchName = ", branchName)
	fmt.Print("productName = ", productName)

	rows, err := db.Query("SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address FROM product p JOIN branchproduct bp ON p.id=bp.productId JOIN branches b ON bp.branchId=b.id WHERE b.name=? AND p.name LIKE ?", branchName, productName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product Product
	var products []Product
	var branch Branch
	var error bool
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category, &branch.ID, &branch.Name, &branch.Address); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
		} else {
			products = append(products, product)
		}
	}

	if !error {
		var response BranchProductsForMenu
		response.Product = products
		response.Branch = branch
		c.IndentedJSON(http.StatusOK, response)
	}
}

func InsertProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	var newProduct BranchProductForInsert

	if err := c.Bind(&newProduct); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// INSERT TO PRODUCT TABLE
	_, err := db.Query("INSERT INTO product (name, price, category, pictureUrl) VALUES (?,?,?,?)",
		newProduct.Product.Name,
		newProduct.Product.Price,
		newProduct.Product.Category,
		newProduct.Product.PictureUrl,
	)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "insert to product table failed"})
		return
	}

	// GET ID BRANCH and NEW PRODUCT ID
	var branchId int
	var productId int

	rows, errQueryGetIDs := db.Query("SELECT p.id, b.id FROM product p, branches b WHERE b.name=? AND p.name=?", newProduct.Branch, newProduct.Product.Name)
	if errQueryGetIDs != nil {
		log.Println(errQueryGetIDs)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	for rows.Next() {
		if errorGetIDs := rows.Scan(&productId, &branchId); err != nil {
			log.Println(errorGetIDs)
			c.JSON(400, gin.H{"error": "id product and branch not found"})
			return
		}
	}

	// INSERT TO BRANCH-PRODUCT TABLE
	_, errInsertBranchProduct := db.Query("INSERT INTO branchproduct (branchId, productId, productQuantity) VALUES (?,?,?)",
		branchId,
		productId,
		newProduct.Quantity,
	)

	if errInsertBranchProduct != nil {
		log.Println(errInsertBranchProduct)
		c.JSON(400, gin.H{"error": "insert to branch-product table failed"})
		return
	} else {
		newProduct.Product.ID = productId
		c.IndentedJSON(http.StatusCreated, newProduct)
	}

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
