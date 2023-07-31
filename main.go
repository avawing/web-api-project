package main

import (
	"awesomeProject/models"
	"log"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("users", getUsers)
		v1.GET("users/:id", getUser)
		v1.POST("users", createUser)
		v1.PUT("users/:id", updateUser)
		v1.DELETE("users/:id", deleteUser)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

func getUsers(c *gin.Context) {
	users, err := models.GetUsers(10)
	checkErr(err)

	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := models.GetUserById(id)
	checkErr(err)
	// if the name is blank we can assume nothing is found
	if user.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func createUser(c *gin.Context) {
	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddUser(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func updateUser(c *gin.Context) {
	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.UpdateUser(json, userId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.DeleteUser(userId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
