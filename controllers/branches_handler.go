package controllers

import (
	"database/sql"
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

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = branches
	c.IndentedJSON(http.StatusOK, response)
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

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Branch insert successful"}
	c.JSON(http.StatusOK, response)
}

func UpdateBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchId := c.Param("id")
	branchName := c.PostForm("name")
	branchAddress := c.PostForm("address")

	var branch Branch
	//Check if branch exists
	errGetOldBranch := db.QueryRow("SELECT id, name, address FROM branches WHERE id = ?", branchId).Scan(&branch.ID, &branch.Name, &branch.Address)
	if errGetOldBranch == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch does not exist"})
		return
	} else if errGetOldBranch != nil {
		log.Fatal(errGetOldBranch)
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

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Branch update successful"}
	c.JSON(http.StatusOK, response)
}

func DeleteBranch(c *gin.Context) {
	db := connect()
	defer db.Close()

	branchId := c.Param("id")

	var branch Branch
	//Check if branch exists
	errGetOldBranch := db.QueryRow("SELECT id, name, address FROM branches WHERE id = ?", branchId).Scan(&branch.ID, &branch.Name, &branch.Address)
	if errGetOldBranch == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	} else if errGetOldBranch != nil {
		log.Fatal(errGetOldBranch)
		return
	}

	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input branch id cannot be empty"})
		return
	}

	_, err := db.Exec("DELETE FROM branches WHERE id = ?", branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in delete query"})
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "Branch delete successful"}
	c.JSON(http.StatusOK, response)
}
