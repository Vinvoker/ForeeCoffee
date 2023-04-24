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
	var branchId int
	queryBranch := "SELECT id FROM `branches` WHERE name LIKE ?"
	row, _ := db.Prepare(queryBranch)
	err := row.QueryRow(branchName).Scan(&branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama Branch tidak ditemukan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// cek apakah product yang ingin dimasukkan ke branch ada di list all product
	query := "SELECT id FROM `product` WHERE name LIKE ?"
	rows, _ := db.Prepare(query)
	var productId string
	err = rows.QueryRow(productName).Scan(&productId)
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

	err = rows.QueryRow(productId, branchId).Scan(&productQuantity)
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
		branchId,
		productId,
		productStok)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Product ke Branch Success",
		"status":  http.StatusOK,
	})

}

func UpdateMenuBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := "%" + c.Param("branchName") + "%"
	productName := "%" + c.PostForm("productName") + "%"
	plusStok, _ := strconv.Atoi(c.PostForm("plusStok"))

	// mencari id branch
	var branchId int
	queryBranch := "SELECT id FROM `branches` WHERE name LIKE ?"
	row, _ := db.Prepare(queryBranch)
	err := row.QueryRow(branchName).Scan(&branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama Branch tidak ditemukan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// mencari id product
	query := "SELECT id FROM `product` WHERE name LIKE ?"
	rows, _ := db.Prepare(query)
	var productId string
	err = rows.QueryRow(productName).Scan(&productId)
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

	// mendapatkan jumlah stok product di branch tersebut
	query = "SELECT bp.productQuantity FROM `branchproduct` bp JOIN branches b ON bp.branchId = b.id WHERE bp.productId = ? AND b.id = ?"
	rows, _ = db.Prepare(query)
	var productQuantity int

	err = rows.QueryRow(productId, branchId).Scan(&productQuantity)
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
	newStok := productQuantity + plusStok
	queryUpdate := "UPDATE `branchproduct` SET `productQuantity`= ? WHERE branchId = ? AND productId = ?"
	_, err = db.Exec(queryUpdate,
		newStok,
		branchId,
		productId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("error 2: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update stok product berhasil",
		"status":  http.StatusOK,
	})
}

func DeleteMenuBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := "%" + c.Param("branchName") + "%"
	productName := "%" + c.Query("productName") + "%"

	// delete query
	queryDelete := "DELETE FROM `branchproduct` WHERE productId = (SELECT p.id FROM product p WHERE p.name LIKE ?) AND branchId = (SELECT b.id FROM branches b WHERE b.name LIKE ?)"
	_, err := db.Exec(queryDelete,
		productName,
		branchName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Nama branch atau product tidak dapat ditemukan",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(branchName)
		log.Println(productName)
		log.Println("error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete produk berhasil",
		"status":  http.StatusOK,
	})
}
