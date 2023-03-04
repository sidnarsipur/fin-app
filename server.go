package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users", getUser)
	router.GET("/users/:user_id", getUserAccounts)
	router.POST("/users", postUser)

	router.Run("localhost:8080")
}

type user struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

var users = []user{
	{
		UserID: 1,
		Name:   "John",
		Email:  "test@gmail.com",
	},
}

type account struct {
	AccountID int    `json:"account_id"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	UserID    int    `json:"user_id"`
}

var accounts = []account{
	{
		AccountID: 100,
		Name:      "John's Checking",
		Balance:   1000,
		UserID:    1,
	},
}

type creditCard struct {
	CardID  int    `json:"card_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	Limit   int    `json:"limit"`
	UserID  int    `json:"user_id"`
}

var creditCards = []creditCard{
	{
		CardID:  200,
		Name:    "John's Credit Card",
		Balance: 1000,
		Limit:   10000,
		UserID:  1,
	},
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserAccounts(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, a := range accounts {
		if a.UserID == userId {
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	for _, b := range creditCards {
		if b.UserID == userId {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

}
