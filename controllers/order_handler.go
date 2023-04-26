package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context) { //blm selesai
	db := connect()
	defer db.Close()

	activeUserId := GetUserId(c)

	branchName := c.PostForm("branch_name")
	productName := c.PostFormArray("product_name[]")
	quantity := c.PostFormArray("quantity[]")

	var branchId int
	err := db.QueryRow("SELECT id FROM branch WHERE name = ?", branchName).Scan(&branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product name"})
		return
	}

	var orderId int
	now := time.Now()
	err = db.QueryRow("INSERT INTO orders (user_id, branch_id, transaction_time) VALUES (?, ?, ?) RETURNING id", activeUserId, branchId, now).Scan(&orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, productName := range productName {
		var productId int
		err = db.QueryRow("SELECT id FROM products WHERE name = ?", productName).Scan(&productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product name"})
			return
		}

		quantity, err := strconv.Atoi(quantity[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
			return
		}

		_, err = db.Exec("INSERT INTO orderdetails (order_id, product_id, quantity) VALUES (?, ?, ?)", orderId, productId, quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "order created successfully"})
}

func HistoryOrder(c *gin.Context) { //masi error
	db := connect()
	defer db.Close()

	activeUserId := GetUserId(c)

	query := "SELECT o.id, o.transactionTime, o.status, b.name, p.name, od.quantity FROM `order` o " +
		"JOIN `orderdetails` od ON o.id = od.orderId " +
		"JOIN branches b ON o.branchid = b.id " +
		"JOIN product p ON od.productId = p.id WHERE o.userId = ?;"
	rows, err := db.Query(query, activeUserId)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Order History query"})
		return
	}

	var orders []Order
	for rows.Next() {
		var order Order
		var branch Branch
		var product Product
		var orderDetail OrderDetails

		if err := rows.Scan(
			&order.ID,
			&order.TransactionTime,
			&order.Status,
			&branch.Name,
			&product.Name,
			&orderDetail.Quantity,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "no order"})
			return
		}

		order.Details = []OrderDetails{orderDetail}
		order.TotalPrice = orderDetail.Quantity * product.Price
		orders = append(orders, order)
	}

	response := OrderHistory{Order: orders}
	c.IndentedJSON(http.StatusOK, response)
}

func UpdateOrderStatus(c *gin.Context) {
	db := connect()
	defer db.Close()

	status := c.PostForm("status")
	orderId := c.Param("id")

	_, errQueryUpdateOrderStatus := db.Exec("UPDATE `order` SET `status`=? WHERE `id`=?",
		status,
		orderId,
	)

	if errQueryUpdateOrderStatus != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Order Status FAILED"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Order Status SUCCESS"})
}
