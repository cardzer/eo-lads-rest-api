package main

type BakeryResponse struct {
	ID             string     `json:"_ID"`
	BakeryType     string     `json:"TYPE"`
	BakeryName     string     `json:"NAME"`
	BakeryLocation string     `json:"LOCATION"`
	Batters        []Batters  `json:"BATTERS"`
	Toppings       []Toppings `json:"TOPPINGS"`
}

type Toppings struct {
	ID           string `json:"_ID"`
	Type         string `json:"TYPE"`
	Colour       string `json:"COLOUR"`
	Availability bool   `json:"AVAILABILITY"`
}

type Batters struct {
	ID           string `json:"_ID"`
	Type         string `json:"TYPE"`
	Colour       string `json:"COLOUR"`
	Availability bool   `json:"AVAILABILITY"`
}

type Payment struct {
	ID         string  `json:"_ID"`
	PaymentRef string  `json:"PAYMENT_REF"`
	Amount     float64 `json:"AMOUNT"`
}

type Customer struct {
	ID        string `json:"_ID"`
	FirstName string `json:"FIRST_NAME"`
	LastName  string `json:"LAST_NAME"`
	Email     string `json:"EMAIL"`
	Phone     int64  `json:"PHONE"`
}

type Orders struct {
	ID         string `json:"_ID"`
	CustomerID string `json:"CUSTOMER_ID"`
	BakeryID   string `json:"BAKERY_ID"`
	PaymentID  string `json:"PAYMENT_ID"`
	Timestamp  string `json:"TIMESTAMP"`
}

type OrdersResponse struct {
	ID              int      `json:"_ID"`
	BakeryBatterID  int      `json:"BAKERY_BATTER_ID"`
	BakeryToppingID int      `json:"BAKERY_TOPPING_ID"`
	BakeryType      string   `json:"BAKERY_TYPE"`
	BakeryName      string   `json:"BAKERY_NAME"`
	BakeryPpu       float64  `json:"BAKERY_PPU"`
	BakeryLocation  string   `json:"BAKERY_LOCATION"`
	Toppings        Toppings `json:"TOPPINGS"`
	Batters         Batters  `json:"BATTERS"`
	Orders          Orders   `json:"ORDERS"`
	Payment         Payment  `json:"PAYMENT"`
	Customer        Customer `json:"CUSTOMER"`
}
