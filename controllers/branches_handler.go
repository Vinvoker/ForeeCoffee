package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBranches(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * " +
		"FROM branches "
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Branch query"})
		return
	}

	var branches []Branch
	var branch Branch
	for rows.Next() {
		if err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
		); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "branches not found"})
			return
		} else {
			branches = append(branches, branch)
		}
	}

	c.IndentedJSON(http.StatusOK, branches)
}

func InsertBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchName := c.PostForm("name")
	branchAddress := c.PostForm("address")

	if branchName == "" || branchAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	// Insert new branch into database
	_, err := db.Exec("INSERT INTO branches (name, address) VALUES (?, ?)", branchName, branchAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in insert query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch insert successful"})
}

func UpdateBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchId := c.Param("id")
	branchName := c.PostForm("name")
	branchAddress := c.PostForm("address")

	_, errGetOldBranch := db.Exec("SELECT * FROM branches WHERE id=?", branchId)
	if errGetOldBranch != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Branch doesn't exist"})
		return
	}

	if branchName == "" || branchAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch cannot be empty"})
		return
	}

	_, err := db.Exec("UPDATE branches SET name= ?, address= ? WHERE id=?", branchName, branchAddress, branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in update query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch update successful"})
}

func DeleteBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchId := c.Param("id")

	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch id cannot be empty"})
		return
	}

	_, err := db.Exec("DELETE FROM branches WHERE id = ?", branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in delete query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch delete successful"})
}
