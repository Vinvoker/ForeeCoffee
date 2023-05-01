package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCustomerProfile(c *gin.Context) {
	db := connect()
	defer db.Close()

	email := c.PostForm("Email")
	newUsername := c.PostForm("Username")

	if email == "" || newUsername == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input email or username cannot be empty"})
		return
	}

	activeUserId := GetUserId(c)

	var oldUsername string
	errGetOldUsername := db.QueryRow("SELECT `username` FROM `users` WHERE `id`=?", activeUserId).Scan(&oldUsername)
	if errGetOldUsername != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Get old username failed"})
		return
	}

	_, errQueryUpdateProfile := db.Exec("UPDATE `users` SET `email`=?,`username`=? WHERE `id`=?",
		email,
		newUsername,
		activeUserId,
	)

	if errQueryUpdateProfile != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Customer Profile FAILED"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Update Customer Profile Success"}
	c.JSON(http.StatusOK, response)

	if oldUsername != newUsername {
		fmt.Println("change token!!")
		generateToken(c, activeUserId, newUsername, "CUSTOMER")
	}
}

func UpdateCustomerPassword(c *gin.Context) {
	db := connect()
	defer db.Close()

	activeUserId := GetUserId(c)

	oldPassword := c.PostForm("Old password")
	newPassword := c.PostForm("New password")
	repeatNewPassword := c.PostForm("Repeat new password")

	if oldPassword == "" || newPassword == "" || repeatNewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input password cannot be empty"})
		return
	}

	if newPassword != repeatNewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "New passwords do not match"})
		return
	}

	var oldPasswordFromDb string
	errGetRealOldPass := db.QueryRow("SELECT `password` FROM `users` WHERE id=?", activeUserId).Scan(&oldPasswordFromDb)
	if errGetRealOldPass != nil {
		fmt.Println(errGetRealOldPass)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if oldPassword != oldPasswordFromDb {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Old password is wrong"})
		return
	}

	if oldPassword == newPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "New password is same with old password. Please pick other new password",
		})
		return
	}

	_, errQueryUpdatePass := db.Exec("UPDATE `users` SET `password`=? WHERE `id`=?",
		newPassword,
		activeUserId,
	)

	if errQueryUpdatePass != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Customer Password FAILED"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Update Customer Password Success"}
	c.JSON(http.StatusOK, response)
}
