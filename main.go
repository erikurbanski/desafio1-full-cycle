package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/erikurbanski/desafio1-full-cycle/models"
)

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
	c.JSON(200, gin.H{"message": "A new Record Created!"})
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
