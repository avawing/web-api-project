package main

import (
	"awesomeProject/models"
	"log"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("users", GetUsers)
		v1.GET("users/:id", GetUser)
		v1.POST("users", CreateUser)
		v1.PUT("users/:id", UpdateUser)
		v1.DELETE("users/:id", DeleteUser)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(10)
	CheckErr(err)

	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := models.GetUserById(id)
	CheckErr(err)
	// if the name is blank we can assume nothing is found
	if user.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func CreateUser(c *gin.Context) {
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

func UpdateUser(c *gin.Context) {
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
		c.JSON(http.StatusAccepted, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteUser(c *gin.Context) {
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
