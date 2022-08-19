package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/snowflakedb/gosnowflake"
	"log"
	"net/http"
	"os"
	"reflect"
)

func apiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")
		if key != "0b9cdada-cd50-40c3-8c27-e6da7357bf1e" {
			c.AbortWithStatusJSON(401, "Access not granted")
		}
	}
}

func databaseInit() *sql.DB {
	db, err := sql.Open("snowflake", os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getToppings(db *sql.DB) []Toppings {
	rows, err := db.Query("select * from TOPPINGS")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64
	var s string
	var st string
	var b bool
	var toppings []Toppings
	var topping Toppings
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &s, &st, &b)
		}
		if err != nil {
			return nil
		}
		topping.ID = *&i64
		topping.Type = *&s
		topping.Colour = *&st
		topping.Availability = *&b

		toppings = append(toppings, topping)

	}
	fmt.Println(toppings)
	return toppings
}

func getBatters(db *sql.DB) interface{} {
	rows, err := db.Query("select * from BATTERS")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64
	var s string
	var st string
	var b bool
	var toppings []Batters
	var topping Batters
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &s, &st, &b)
		}
		if err != nil {
			return nil
		}
		topping.ID = *&i64
		topping.Type = *&s
		topping.Colour = *&st
		topping.Availability = *&b

		toppings = append(toppings, topping)
	}
	fmt.Println(toppings)
	return toppings
}

func getPayment(db *sql.DB) interface{} {
	rows, err := db.Query("select * from PAYMENT")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64
	var s string
	var f float64
	var payments []Payment
	var payment Payment
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &s, &f)
		}
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		payment.ID = int(*&i64)
		payment.PaymentRef = *&s
		payment.Amount = *&f

		payments = append(payments, payment)
	}
	fmt.Println(payment)
	return payments
}

func getCustomer(db *sql.DB) interface{} {
	rows, err := db.Query("select * from CUSTOMER")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64
	var s string
	var st string
	var stt string
	var f int
	var customers []Customer
	var customer Customer
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &s, &st, &stt, &f)
		}
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		customer.ID = int(*&i64)
		customer.FirstName = *&s
		customer.LastName = *&st
		customer.Email = *&stt
		customer.Phone = int64(*&f)

		customers = append(customers, customer)
	}
	fmt.Println(customer)
	return customers
}

func getBakery(db *sql.DB) interface{} {
	rows, err := db.Query("select bak._id           as ID,       bak.batter_id     as bakery_batter_id,       bak.topping_id    as bakery_topping_id,       bak.type          as bakery_type,      bak.name          as bakery_name,       bak.ppu           as bakery_ppu,      bak.location      as bakery_location,      topp._id          as topping_id,       topp.type         as topping_type,       topp.colour       as topping_colour,      topp.availability as topping_availability,       bat._id           as batter_id,       bat.type          as batter_type,      bat.colour        as batter_colour,       bat.availability  as batter_availability from BAKERY bak,     TOPPINGS topp,     BATTERS bat where bak.TOPPING_ID = topp._ID  and bak.TOPPING_ID = topp._ID;")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64
	var i644 int64
	var i6444 int64
	var s string
	var st string
	var f float64
	var stt string
	var i64444 int64
	var sttt string
	var stttt string
	var b bool
	var i644444 int64
	var sttttt string
	var stttttt string
	var bb bool

	var bakeryResponses []BakeryResponse
	var bakeryResponse BakeryResponse
	var topping Toppings
	var batter Batters
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &i644, &i6444,
				&s, &st,
				&f,
				&stt,

				&i64444,
				&sttt, &stttt,
				&b,

				&i644444,
				&sttttt, &stttttt, &bb)
		}
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		topping.ID = *&i64444
		topping.Type = *&sttt
		topping.Colour = *&stttt
		topping.Availability = *&b

		batter.ID = *&i644444
		batter.Type = *&sttttt
		batter.Colour = *&stttttt
		batter.Availability = *&bb

		bakeryResponse.ID = int(*&i64)
		bakeryResponse.BakeryBatterID = int(*&i644)
		bakeryResponse.BakeryToppingID = int(*&i6444)
		bakeryResponse.BakeryType = *&s
		bakeryResponse.BakeryName = *&st
		bakeryResponse.BakeryPpu = *&f
		bakeryResponse.BakeryLocation = stt

		bakeryResponse.Batters = batter
		bakeryResponse.Toppings = topping

		bakeryResponses = append(bakeryResponses, bakeryResponse)
	}
	fmt.Println(bakeryResponse)
	return bakeryResponses
}

func getBakeryOrders(db *sql.DB) interface{} {
	rows, err := db.Query("select bak._id           as ID,       bak.batter_id     as bakery_batter_id,       bak.topping_id    as bakery_topping_id,       bak.type          as bakery_type,       bak.name          as bakery_name,       bak.ppu           as bakery_ppu,       bak.location      as bakery_location,       topp._id          as topping_id,       topp.type         as topping_type,       topp.colour       as topping_colour,       topp.availability as topping_availability,       bat._id           as batter_id,       bat.type          as batter_type,       bat.colour        as batter_colour,       bat.availability  as batter_availability,       ord._id           as order_id,       ord.customer_id   as order_customer_id,       ord.bakery_id     as order_bakery_id,       ord.payment_id    as order_payment_id,       ord.timestamp     as order_timestamp,       ord.amount        as order_amount,       pay._ID           as payment_id,       pay.PAYMENT_REF   as payment_ref,       pay.AMOUNT        as payment_amount,\n       cus._ID           as customer_id,\n       cus.FIRST_NAME    as customer_first_name,\n       cus.LAST_NAME     as customer_last_name,\n       cus.EMAIL         as customer_email,\n       cus.PHONE         as customer_phone\nfrom BAKERY bak,\n     TOPPINGS topp,\n     BATTERS bat,\n     ORDERS ord,\n     PAYMENT pay,\n     CUSTOMER cus\nwhere bak.TOPPING_ID = topp._ID\n  and bak.TOPPING_ID = topp._ID\n  and ord.BAKERY_ID = bak._ID\n  and ord.CUSTOMER_ID = cus._ID\n  and ord.PAYMENT_ID = pay._ID")
	if err != nil {
		log.Fatalf("Error doing query")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows")
		}
	}(rows)
	var i64 int64      //ID
	var i644 int64     //BakeryBatterID
	var i6444 int64    //BakeryToppingID
	var s string       //BakeryType
	var st string      //BakeryName
	var f float64      //BakeryPpu
	var stt string     //BakeryLocation
	var i64444 int64   //ToppingID
	var sttt string    //ToppingType
	var stttt string   //ToppingColour
	var b bool         //ToppingAvailability
	var i644444 int64  //BatterID
	var sttttt string  //BatterType
	var stttttt string //BatterColour
	var bb bool        //BatterAvailability

	var i6444444 int64    //OrderID
	var i64444444 int64   //OrderCustomerID
	var i644444444 int64  //OrderBakeryID
	var i6444444444 int64 //OrderPaymentID
	var sttttttt string   //OrderTimestamp
	var ff float64        //OrderAmount

	var i64444444444 int64 //PaymentID
	var stttttttt string   //PaymentRef
	var fff float64        //PaymentAmount

	var i644444444444 int64  //CustomerID
	var sttttttttt string    //CustomerFirstName
	var stttttttttt string   //CustomerLastName
	var sttttttttttt string  //CustomerEmail
	var i6444444444444 int64 //CustomerPhone

	var OrdersResponses []OrdersResponse
	var OrdersResponse OrdersResponse
	var topping Toppings
	var batter Batters
	var orders Orders
	var customer Customer
	var payment Payment
	for rows.Next() {
		column_types, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("ERROR: ColumnTypes() failed. err: %v", err)
		}
		column_type := column_types[0].ScanType()
		switch column_type {
		case reflect.TypeOf(i64):
			err = rows.Scan(&i64, &i644, &i6444,
				&s, &st,
				&f,
				&stt,

				&i64444,
				&sttt, &stttt,
				&b,

				&i644444,

				&sttttt, &stttttt,

				&bb,

				&i6444444, &i64444444, &i644444444, &i6444444444,

				&sttttttt,

				&ff,

				&i64444444444,

				&stttttttt,

				&fff,

				&i644444444444,

				&sttttttttt,
				&stttttttttt,
				&sttttttttttt,

				&i6444444444444,
			)
		}
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		topping.ID = *&i64444
		topping.Type = *&sttt
		topping.Colour = *&stttt
		topping.Availability = *&b

		batter.ID = *&i644444
		batter.Type = *&sttttt
		batter.Colour = *&stttttt
		batter.Availability = *&bb

		orders.ID = int(*&i6444444)
		orders.CustomerID = int(*&i64444444)
		orders.BakeryID = int(*&i644444444)
		orders.PaymentID = int(*&i6444444444)
		orders.Timestamp = *&sttttttt
		orders.Amount = *&ff

		payment.ID = int(*&i64444444444)
		payment.PaymentRef = *&stttttttt
		payment.Amount = *&fff

		customer.ID = int(*&i644444444444)
		customer.FirstName = *&sttttttttt
		customer.LastName = *&stttttttttt
		customer.Email = *&sttttttttttt
		customer.Phone = *&i6444444444444

		OrdersResponse.ID = int(*&i64)
		OrdersResponse.BakeryBatterID = int(*&i644)
		OrdersResponse.BakeryToppingID = int(*&i6444)
		OrdersResponse.BakeryType = *&s
		OrdersResponse.BakeryName = *&st
		OrdersResponse.BakeryPpu = *&f
		OrdersResponse.BakeryLocation = stt

		OrdersResponse.Batters = batter
		OrdersResponse.Toppings = topping
		OrdersResponse.Orders = orders
		OrdersResponse.Customer = customer
		OrdersResponse.Payment = payment

		OrdersResponses = append(OrdersResponses, OrdersResponse)
	}
	fmt.Println(OrdersResponse)
	return OrdersResponses
}

func main() {
	r := gin.Default()
	db := databaseInit()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	r.GET("/bakery/toppings", apiKeyAuth(), func(c *gin.Context) {
		var data = getToppings(db)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/batters", apiKeyAuth(), func(c *gin.Context) {

		var data = getBatters(db)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/payment", apiKeyAuth(), func(c *gin.Context) {

		var data = getPayment(db)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/customer", apiKeyAuth(), func(c *gin.Context) {

		var data = getCustomer(db)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery", apiKeyAuth(), func(c *gin.Context) {

		var data = getBakery(db)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/bakery/orders", apiKeyAuth(), func(c *gin.Context) {

		var data = getBakeryOrders(db)
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
