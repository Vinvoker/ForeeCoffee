package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	activeUserId := GetUserId(c)

	branchName := c.PostForm("branch_name")
	productName := c.PostFormArray("product_name[]")
	quantity := c.PostFormArray("quantity[]")

	// pilih branch
	var branchId int
	err := db.QueryRow("SELECT id FROM branches WHERE name = ?", branchName).Scan(&branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch name"})
		return
	}

	// insert ke tabel order
	var orderId int
	now := time.Now()
	result, err := db.Exec("INSERT INTO `order` (transactionTime, userId, status, branchId) VALUES (?, ?, 'ONGOING',  ?)", now, activeUserId, branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// mengambil orderId dari order yang baru saja dilakukan
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderId = int(lastInsertId)

	// insert ke tabel orderdetails
	for i, productName := range productName {
		var productId int
		err = db.QueryRow("SELECT id FROM product WHERE name = ?", productName).Scan(&productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product name", "debug": err.Error()})
			return
		}

		quantity, err := strconv.Atoi(quantity[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
			return
		}

		_, err = db.Exec("INSERT INTO orderdetails (orderId, productId, quantity) VALUES (?, ?, ?)", orderId, productId, quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// mengurangi quantity product di cabang tersebut
		_, err = db.Exec("UPDATE branchproduct bp "+
			"JOIN branches b ON bp.branchId = b.id "+
			"JOIN product p ON bp.productId = p.id "+
			"SET bp.productQuantity = bp.productQuantity - ? "+
			"WHERE b.id = ? AND p.id = ?", quantity, branchId, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "order created successfully"}
	c.JSON(http.StatusOK, response)
}

func HistoryOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	activeUserId := GetUserId(c)

	var orders []Order
	query := "SELECT id FROM `order` WHERE userId = ?"
	orderRows, err := db.Query(query, activeUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	defer orderRows.Close()

	for orderRows.Next() {
		totalPrice := 0
		var orderID int
		err := orderRows.Scan(&orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan order ID"})
			return
		}

		var order Order
		var orderDetails []OrderDetails
		detailRows, err := db.Query("SELECT product.name, product.price, orderdetails.quantity FROM product"+
			" JOIN orderdetails ON product.id = orderdetails.productId"+
			" JOIN `order` ON orderdetails.orderid = order.id WHERE order.id = ?", orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch order details"})
			return
		}
		defer detailRows.Close()

		for detailRows.Next() {
			var orderDetail OrderDetails
			err := detailRows.Scan(&orderDetail.Product.Name, &orderDetail.Product.Price, &orderDetail.Quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan order detail"})
				return
			}
			totalPrice += (orderDetail.Product.Price * orderDetail.Quantity)
			orderDetails = append(orderDetails, orderDetail)
		}

		order.TotalPrice = totalPrice
		order.Details = orderDetails
		order.ID = orderID
		err = db.QueryRow("SELECT transactionTime, branchid, `status` FROM `order` WHERE id = ?", orderID).Scan(&order.TransactionTime, &order.BranchID, &order.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		orders = append(orders, order)
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = orders
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Order Status Failed"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Update Order Status Success"}
	c.JSON(http.StatusOK, response)
}
