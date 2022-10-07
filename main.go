package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/snowflakedb/gosnowflake"
	"net/http"
	"os"
)

func apiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")
		if key != "0b9cdada-cd50-40c3-8c27-e6da7357bf1e" {
			c.AbortWithStatusJSON(401, "Access not granted")
		}
	}
}

func getToppings() []Toppings {
	var toppings []Toppings // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/TOPPINGS.json")
	err := json.Unmarshal(fileBytes, &toppings)
	if err != nil {
		fmt.Println(err)
	}
	return toppings
}

func getBatters() interface{} {
	var batters []Batters // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/BATTERS.json")
	err := json.Unmarshal(fileBytes, &batters)
	if err != nil {
		fmt.Println(err)
	}
	return batters
}

func getPayment() interface{} {
	var payments []Payment // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/PAYMENTS.json")
	err := json.Unmarshal(fileBytes, &payments)
	if err != nil {
		fmt.Println(err)
	}
	return payments
}

func getCustomer() interface{} {
	var customers []Customer // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/CUSTOMER.json")
	err := json.Unmarshal(fileBytes, &customers)
	if err != nil {
		fmt.Println(err)
	}
	return customers
}

func getBakery() interface{} {
	var bakery []BakeryResponse // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/BAKERY.json")
	err := json.Unmarshal(fileBytes, &bakery)
	if err != nil {
		fmt.Println(err)
	}
	return bakery
}

func getOrders() interface{} {
	var orders []Orders // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("data/ORDERS.json")
	err := json.Unmarshal(fileBytes, &orders)
	if err != nil {
		fmt.Println(err)
	}
	return orders
}

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	r.GET("/bakery/toppings", apiKeyAuth(), func(c *gin.Context) {
		var data = getToppings()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/batters", apiKeyAuth(), func(c *gin.Context) {

		var data = getBatters()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/payment", apiKeyAuth(), func(c *gin.Context) {

		var data = getPayment()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/customer", apiKeyAuth(), func(c *gin.Context) {

		var data = getCustomer()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery", apiKeyAuth(), func(c *gin.Context) {

		var data = getBakery()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/orders", apiKeyAuth(), func(c *gin.Context) {
		var data = getOrders()
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
