package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductsAndTheirBranches(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var products []ProductsDetails
	var product ProductsDetails
	for rows.Next() {
		var branch Branch
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {

			var productIdFound = false
			if len(products) > 0 {
				for i := 0; i < len(products); i++ {
					if products[i].ID == product.ID {
						products[i].Branch = append(products[i].Branch, branch)
						productIdFound = true
						break
					}
				}
			}

			if !productIdFound {
				product.Branch = []Branch{branch}
				products = append(products, product)
			}
		}
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = products
	c.JSON(http.StatusOK, response)

}

func GetAllProductsByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	if branchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address, bp.productQuantity " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id " +
		"WHERE b.name=?"

	rows, err := db.Query(query, branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product ProductForMenu
	var products []ProductForMenu
	var productQuantity int
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&productQuantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {
			if productQuantity > 0 {
				product.Status = "AVAILABLE"
			} else {
				product.Status = "UNAVAILABLE"
			}
			products = append(products, product)
		}
	}

	var data BranchProductsForMenu
	data.Product = products
	data.Branch = branch

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func GetProductsCoffeeByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	if branchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address, bp.productQuantity " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id " +
		"WHERE b.name=? AND p.category='COFFEE'"

	rows, err := db.Query(query, branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product ProductForMenu
	var products []ProductForMenu
	var productQuantity int
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&productQuantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {
			if productQuantity > 0 {
				product.Status = "AVAILABLE"
			} else {
				product.Status = "UNAVAILABLE"
			}
			products = append(products, product)
		}
	}

	var data BranchProductsForMenu
	data.Product = products
	data.Branch = branch

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func GetProductsYakultByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

	if branchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address, bp.productQuantity " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id " +
		"WHERE b.name=? AND p.category='YAKULT'"

	rows, err := db.Query(query, branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product ProductForMenu
	var products []ProductForMenu
	var productQuantity int
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&productQuantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {
			if productQuantity > 0 {
				product.Status = "AVAILABLE"
			} else {
				product.Status = "UNAVAILABLE"
			}
			products = append(products, product)
		}
	}

	var data BranchProductsForMenu
	data.Product = products
	data.Branch = branch

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func GetProductsTeaByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")
	if branchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address, bp.productQuantity " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id " +
		"WHERE b.name=? AND p.category='TEA'"

	rows, err := db.Query(query, branchName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product ProductForMenu
	var products []ProductForMenu
	var productQuantity int
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&productQuantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {
			if productQuantity > 0 {
				product.Status = "AVAILABLE"
			} else {
				product.Status = "UNAVAILABLE"
			}
			products = append(products, product)
		}
	}

	var data BranchProductsForMenu
	data.Product = products
	data.Branch = branch

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func GetProductByNameAndBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")
	productName := c.Query("Name")

	if branchName == "" || productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch and name cannot be empty"})
		return
	}

	productName = "%" + c.Query("Name") + "%"

	query := "SELECT p.id, p.name, p.price, p.pictureUrl, p.category, b.id, b.name, b.address, bp.productQuantity " +
		"FROM product p " +
		"JOIN branchproduct bp ON p.id=bp.productId " +
		"JOIN branches b ON bp.branchId=b.id " +
		"WHERE b.name=? AND p.name LIKE ?"

	rows, err := db.Query(query, branchName, productName)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	var product ProductForMenu
	var products []ProductForMenu
	var productQuantity int
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.PictureUrl,
			&product.Category,
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&productQuantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "products not found"})
			return
		} else {
			if productQuantity > 0 {
				product.Status = "AVAILABLE"
			} else {
				product.Status = "UNAVAILABLE"
			}
			products = append(products, product)
		}
	}

	var data BranchProductsForMenu
	data.Product = products
	data.Branch = branch

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func InsertProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	var newProduct Product

	if errBind := c.Bind(&newProduct); errBind != nil {
		fmt.Print(errBind)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if newProduct.Name == "" || newProduct.Price <= 0 || newProduct.Category == "" || newProduct.PictureUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input product cannot be empty"})
		return
	}

	// Check if new product really does not have its name already
	var oldProduct int
	errGetOldProduct := db.QueryRow("SELECT `id` FROM `product` WHERE `name`=?", newProduct.Name).Scan(&oldProduct)
	if errGetOldProduct != sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product existed already"})
		return
	}

	// INSERT TO PRODUCT TABLE
	_, errQueryInsertProduct := db.Query("INSERT INTO product (name, price, category, pictureUrl) VALUES (?,?,?,?)",
		newProduct.Name,
		newProduct.Price,
		newProduct.Category,
		newProduct.PictureUrl,
	)

	if errQueryInsertProduct != nil {
		log.Println(errQueryInsertProduct)
		c.JSON(400, gin.H{"error": "insert to product table failed"})
		return
	}

	errGetNewProductId := db.QueryRow("SELECT `id` FROM `product` WHERE `name`=?", newProduct.Name).Scan(&newProduct.ID)
	if errGetNewProductId != nil {
		fmt.Println(errGetNewProductId)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Get new product ID failed"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = newProduct
	c.JSON(http.StatusOK, response)
}

func UpdateProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	productId := c.Param("id")
	productName := c.PostForm("name")
	productPrice := c.PostForm("price")
	productUrl := c.PostForm("url")
	productCategory := c.PostForm("category")

	var product Product
	//Check if product exists
	errGetOldProduct := db.QueryRow("SELECT id, name, price FROM product WHERE id = ?", productId).Scan(&product.ID, &product.Name, &product.Price)
	if errGetOldProduct == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product does not exist"})
		return
	} else if errGetOldProduct != nil {
		log.Fatal(errGetOldProduct)
		return
	}

	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product ID cannot be empty"})
		return
	}

	_, err := db.Exec("UPDATE product SET name= ?, price= ?, pictureurl= ?, category= ? WHERE id=?", productName, productPrice, productUrl, productCategory, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in update query"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Product update successful"}
	c.JSON(http.StatusOK, response)
}

func DeleteProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	productId := c.Param("id")

	var product Product
	//Check if product exists
	errGetOldProduct := db.QueryRow("SELECT id, name, price FROM product WHERE id = ?", productId).Scan(&product.ID, &product.Name, &product.Price)
	if errGetOldProduct == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product does not exist"})
		return
	} else if errGetOldProduct != nil {
		log.Fatal(errGetOldProduct)
		return
	}

	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product ID cannot be empty"})
		return
	}

	_, errDelFromProductBranch := db.Exec("DELETE FROM branchproduct WHERE productid = ?", productId)
	if errDelFromProductBranch != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in delete query for product branch"})
		return
	}

	_, errDelFromOrderDetails := db.Exec("DELETE FROM orderdetails WHERE productid = ?", productId)
	if errDelFromOrderDetails != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in delete query for product branch"})
		return
	}

	_, err := db.Exec("DELETE FROM product WHERE id = ?", productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in delete query"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Product delete successful"}
	c.JSON(http.StatusOK, response)
}
