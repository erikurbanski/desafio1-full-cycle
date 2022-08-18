package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/erikurbanski/desafio1-full-cycle/models"
)

type AddAccountRequestBody struct {
	Number string  `json:"account_number"`
	Amount float64 `json:"amount"`
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()
	router := r.Group("/bank-accounts")
	{
		router.POST("/", createAccount)
		router.POST("/transfer", transfer)
		router.GET("/", getAllAccounts)
	}
	r.Run(":8000")
}

func createAccount(c *gin.Context) {
	body := AddAccountRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var account models.Account

	account.Number = body.Number
	account.Amount = body.Amount

	accountId := models.InsertAccount(account)
	if accountId == 0 {
		c.JSON(404, gin.H{"error": "Insert error!"})
	} else {
		c.JSON(http.StatusCreated, accountId)
	}
}

func transfer(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}

func getAllAccounts(c *gin.Context) {
	accounts, err := models.GetAccounts()
	checkErr(err)
	if accounts == nil {
		c.JSON(404, gin.H{"error": "No records found!"})
		return
	} else {
		c.JSON(200, gin.H{"data": accounts})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
