package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertMenuBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := "%" + c.Param("branchName") + "%"
	productName := "%" + c.PostForm("productName") + "%"
	productStok := c.PostForm("productQuantity")

	// mencari id branch
	var branch Branch
	queryBranch := "SELECT id, name, address FROM `branches` WHERE name LIKE ?"
	row, _ := db.Prepare(queryBranch)
	err := row.QueryRow(branchName).Scan(&branch.ID, &branch.Name, &branch.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama Branch tidak ditemukan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// cek apakah product yang ingin dimasukkan ke branch ada di list all product
	query := "SELECT id, name, price, pictureUrl, category FROM `product` WHERE name LIKE ?"
	rows, _ := db.Prepare(query)
	var product Product
	err = rows.QueryRow(productName).Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Product tidak tersedia di list product...",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 1: ", err)
		return
	}

	// cek apakah product sudah ada di branch tersebut atau belum
	query = "SELECT bp.productQuantity FROM `branchproduct` bp JOIN branches b ON bp.branchId = b.id WHERE bp.productId = ? AND b.id = ?"
	rows, _ = db.Prepare(query)
	var productQuantity int

	err = rows.QueryRow(product.ID, branch.ID).Scan(&productQuantity)
	if err != nil && err != sql.ErrNoRows {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 2: ", err)
		return
	}
	if productQuantity != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product sudah terdaftar di branch tersebut",
			"status":  http.StatusBadRequest,
		})
		return
	}
	// insert product ke branch
	queryInsert := "INSERT INTO `branchproduct`(`branchId`, `productId`, `productQuantity`) VALUES (?,?,?)"
	_, err = db.Exec(queryInsert,
		branch.ID,
		product.ID,
		productStok)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Branch":        branch,
		"Product":       product,
		"Product Stock": productStok,
	})

}

func UpdateMenuBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := "%" + c.Param("branchName") + "%"
	productName := "%" + c.PostForm("productName") + "%"
	plusStok, _ := strconv.Atoi(c.PostForm("plusStok"))

	// mencari id branch
	var updateProduct UpdateProductBranch
	queryBranch := "SELECT id, name, address FROM `branches` WHERE name LIKE ?"
	row, _ := db.Prepare(queryBranch)
	err := row.QueryRow(branchName).Scan(&updateProduct.Branch.ID, &updateProduct.Branch.Name, &updateProduct.Branch.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama Branch tidak ditemukan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// mencari id product
	query := "SELECT id, name, price, pictureUrl, category FROM `product` WHERE name LIKE ?"
	rows, _ := db.Prepare(query)
	err = rows.QueryRow(productName).Scan(&updateProduct.Product.ID, &updateProduct.Product.Name, &updateProduct.Product.Price, &updateProduct.Product.PictureUrl, &updateProduct.Product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Product tidak tersedia di list product...",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 1: ", err)
		return
	}

	// mendapatkan jumlah stok lama product di branch tersebut
	query = "SELECT bp.productQuantity FROM `branchproduct` bp JOIN branches b ON bp.branchId = b.id WHERE bp.productId = ? AND b.id = ?"
	rows, _ = db.Prepare(query)

	err = rows.QueryRow(updateProduct.Product.ID, updateProduct.Branch.ID).Scan(&updateProduct.OldStock)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Product tidak tersedia di branch ini",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 2: ", err)
		return
	}
	updateProduct.NewStock = updateProduct.OldStock + plusStok
	queryUpdate := "UPDATE `branchproduct` SET `productQuantity`= ? WHERE branchId = ? AND productId = ?"
	_, err = db.Exec(queryUpdate,
		updateProduct.NewStock,
		updateProduct.Branch.ID,
		updateProduct.Product.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 2: ", err)
		return
	}
	c.IndentedJSON(http.StatusOK, updateProduct)
}

func DeleteMenuBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := "%" + c.Param("branchName") + "%"
	productName := "%" + c.Query("productName") + "%"

	// mencari id branch
	var branch Branch
	queryBranch := "SELECT id, name, address FROM `branches` WHERE name LIKE ?"
	row, _ := db.Prepare(queryBranch)
	err := row.QueryRow(branchName).Scan(&branch.ID, &branch.Name, &branch.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama Branch tidak ditemukan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// mencari id product
	query := "SELECT id, name, price, pictureUrl, category FROM `product` WHERE name LIKE ?"
	rows, _ := db.Prepare(query)
	var product Product
	err = rows.QueryRow(productName).Scan(&product.ID, &product.Name, &product.Price, &product.PictureUrl, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Product tidak tersedia di list product...",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 1: ", err)
		return
	}

	// delete query
	queryDelete := "DELETE FROM `branchproduct` WHERE productId = ? AND branchId = ?"
	result, err := db.Exec(queryDelete,
		product.ID,
		branch.ID)
	rowsAffected, _ := result.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error: ", err)
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product tidak ditemukan di branch ini",
			"status":  http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete product success",
		"Branch":  branch,
		"Product": product,
	})
}
