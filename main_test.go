package main

import (
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("users", GetUsers)
		v1.GET("users/:id", GetUser)
		v1.POST("users", CreateUser)
		v1.PUT("users/:id", UpdateUser)
		v1.DELETE("users/:id", DeleteUser)
	}
	return router
}

func setup() {

	if err := models.ConnectDatabase(); err != nil {
		log.Fatalf("Oops")
	}
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func TestCreateUser(t *testing.T) {
	setup()
	newUser := models.User{
		Id:           3,
		LastName:     "Bobberton",
		FirstName:    "Bob",
		Email:        "bob@bob.com",
		HasLoan:      false,
		HasOtherLoan: true,
	}
	writer := makeRequest("POST", "/api/v1/users", newUser)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetUsers(t *testing.T) {
	setup()

	writer := makeRequest("GET", "/api/v1/users", models.User{})
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUpdateUser(t *testing.T) {
	setup()

	update := models.User{
		FirstName:    "Bobbly",
		LastName:     "Bobble",
		Email:        "bob@bob.com",
		HasLoan:      false,
		HasOtherLoan: true,
	}

	writer := makeRequest("PUT", "/api/v1/users/1", update)
	assert.Equal(t, http.StatusAccepted, writer.Code)
}

func TestGetUser(t *testing.T) {
	writer := makeRequest("GET", "/api/v1/users/1", models.User{})
	assert.Equal(t, http.StatusOK, writer.Code)
}

func BenchmarkCreateUser(b *testing.B) {
	users := []models.User{
		{
			FirstName:    "Bobbly",
			LastName:     "Bobble",
			Email:        "bob@bob.com",
			HasLoan:      false,
			HasOtherLoan: true,
		},
		{
			FirstName:    "Bobberella",
			LastName:     "Bobble",
			Email:        "bobtastic@bob.com",
			HasLoan:      true,
			HasOtherLoan: true,
		},
		{
			FirstName:    "Bo",
			LastName:     "Bobble",
			Email:        "bo@bob.com",
			HasLoan:      false,
			HasOtherLoan: true,
		},
	}
	for _, user := range users {
		writer := makeRequest("POST", "/api/v1/users", user)
		assert.Equal(b, http.StatusOK, writer.Code)
	}
}

func BenchmarkUpdateUser(b *testing.B) {
	users := []models.User{
		{
			FirstName:    "Bobbly",
			LastName:     "Bobble",
			Email:        "bob@bob.com",
			HasLoan:      false,
			HasOtherLoan: true,
		},
		{
			FirstName:    "Bobberella",
			LastName:     "Bobble",
			Email:        "bobtastic@bob.com",
			HasLoan:      true,
			HasOtherLoan: true,
		},
		{
			FirstName:    "Bo",
			LastName:     "Bobble",
			Email:        "bo@bob.com",
			HasLoan:      false,
			HasOtherLoan: true,
		},
	}
	for _, user := range users {
		writer := makeRequest("PUT", "/api/v1/users/1", user)
		assert.Equal(b, http.StatusAccepted, writer.Code)
	}
}
