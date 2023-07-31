package main

import (
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HasLoan      bool   `json:"has_loan"`
	HasOtherLoan bool   `json:"has_other_loan"`
}

//type Users struct {
//	Users []User `json:"users"`
//}

var users = []User{
	{ID: "1", FirstName: "Blue Train", LastName: "John Coltrane", Email: "Blue@Trane.com", HasLoan: false, HasOtherLoan: false},
	{ID: "2", FirstName: "Bob Train", LastName: "John Warmtrane", Email: "Bob@Trane.com", HasLoan: true, HasOtherLoan: false},
	{ID: "3", FirstName: "Borg Train", LastName: "Johntrane", Email: "Borg@Trane.com", HasLoan: true, HasOtherLoan: true},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUser)
	router.GET("/user/:id", getUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Println(err)
		return
	}
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var newUser User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			user = u
		}
	}
	c.IndentedJSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			u = updatedUser
		}
	}
	c.IndentedJSON(http.StatusAccepted, updatedUser)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	newUsers := make([]User, len(users)-1)
	k := 0

	for i := 0; i < len(users); i++ {
		if users[i].ID != id {
			newUsers[i] = users[k]
			k++
		} else {
			k++
		}
		i++
	}
	users = newUsers

	c.IndentedJSON(http.StatusNoContent, "Deleted")
}
