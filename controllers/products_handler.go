package controllers

import (
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

	c.IndentedJSON(http.StatusOK, products)

}

func GetAllProductsByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

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

	var response BranchProductsForMenu
	response.Product = products
	response.Branch = branch
	c.IndentedJSON(http.StatusOK, response)
}

func GetProductsCoffeeByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

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

	var response BranchProductsForMenu
	response.Product = products
	response.Branch = branch
	c.IndentedJSON(http.StatusOK, response)
}

func GetProductsYakultByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

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

	var response BranchProductsForMenu
	response.Product = products
	response.Branch = branch
	c.IndentedJSON(http.StatusOK, response)
}

func GetProductsTeaByBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")

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

	var response BranchProductsForMenu
	response.Product = products
	response.Branch = branch
	c.IndentedJSON(http.StatusOK, response)
}

func GetProductByNameAndBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.Query("Branch")
	productName := "%" + c.Query("Name") + "%"

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

	var response BranchProductsForMenu
	response.Product = products
	response.Branch = branch
	c.IndentedJSON(http.StatusOK, response)
}

func InsertProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	var newProduct BranchProductForInsert

	if errBind := c.Bind(&newProduct); errBind != nil {
		fmt.Print(errBind)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// INSERT TO PRODUCT TABLE
	_, errQueryInsertProduct := db.Query("INSERT INTO product (name, price, category, pictureUrl) VALUES (?,?,?,?)",
		newProduct.Product.Name,
		newProduct.Product.Price,
		newProduct.Product.Category,
		newProduct.Product.PictureUrl,
	)

	if errQueryInsertProduct != nil {
		log.Println(errQueryInsertProduct)
		c.JSON(400, gin.H{"error": "insert to product table failed"})
		return
	}

	// GET ID BRANCH and NEW PRODUCT ID
	var branchId int
	var productId int

	rows, errQueryGetIDs := db.Query("SELECT p.id, b.id FROM product p, branches b WHERE b.name=? AND p.name=?",
		newProduct.Branch,
		newProduct.Product.Name,
	)

	if errQueryGetIDs != nil {
		log.Println(errQueryGetIDs)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Product query"})
		return
	}

	for rows.Next() {
		if errorGetIDs := rows.Scan(&productId, &branchId); errQueryInsertProduct != nil {
			log.Println(errorGetIDs)
			c.JSON(400, gin.H{"error": "id product and branch not found"})
			return
		}
	}

	// INSERT TO BRANCH-PRODUCT TABLE
	_, errQueryInsertBranchProduct := db.Query("INSERT INTO branchproduct (branchId, productId, productQuantity) VALUES (?,?,?)",
		branchId,
		productId,
		newProduct.Quantity,
	)

	if errQueryInsertBranchProduct != nil {
		log.Println(errQueryInsertBranchProduct)
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
