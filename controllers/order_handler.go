package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	var newOrder Order

	if errBind := c.Bind(&newOrder); errBind != nil {
		fmt.Print(errBind)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
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
	// currentUserRole := c.GetHeader("role")
	// fmt.Println("currentUserRole:", currentUserRole)

	// if currentUserRole != "ADMIN" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Only admins can access this endpoint"})
	// 	return
	// }

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
