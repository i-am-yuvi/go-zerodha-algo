package main

import (
	"github.com/gofiber/fiber/v2"
)

type Balance map[string]float64

type User struct {
	id       string
	balances Balance
}

type Order struct {
	userId   string
	prices   float64
	quantity float64
}

var users = []User{
	{
		id: "1",
		balances: Balance{
			"USD":    1000,
			"GOOGLE": 0,
		},
	},
	{
		id: "2",
		balances: Balance{
			"USD":    1000,
			"GOOGLE": 0,
		},
	},
}

var bids = []Order{}
var asks = []Order{}

const TICKER = "GOOGLE"

func limitOrder(req any, res any) error {

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/order", limitOrder)
	app.Get("/depth", getDepth)
	app.Get("/balance/:userId", getBalance)
	app.Get("/quote", getQuote)

	app.Listen(":3000")
}
