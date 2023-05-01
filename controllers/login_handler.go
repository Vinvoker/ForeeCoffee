package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	db := connect()
	defer db.Close()

	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")
	role := "CUSTOMER"

	_, err := db.Exec("INSERT INTO `users`(`email`, `password`, `username`, `role`) VALUES (?,?,?,?)",
		email,
		password,
		username,
		role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Sign up FAILED",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Sign Up Success"}
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	db := connect()
	defer db.Close()

	email := c.PostForm("email")
	password := c.PostForm("password")

	rows, err := db.Prepare("SELECT id, username, role FROM `users` WHERE email = ? AND password = ?")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var user User
	err = rows.QueryRow(email, password).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login FAILED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	id := user.ID
	username := user.Username
	role := user.Role
	generateToken(c, id, username, role)

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Login successful"}
	c.JSON(http.StatusOK, response)
}

func Logout(c *gin.Context) {
	resetUserToken(c)

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Logout Success"}
	c.JSON(http.StatusOK, response)
}
