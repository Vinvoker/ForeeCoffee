package controllers

import (
	"fmt"
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

func HistoryOrder(c *gin.Context) {

}

func UpdateOrderStatus(c *gin.Context) {

}
