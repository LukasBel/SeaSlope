package main

import (
	"SeaSlope/slope"
	"fmt"
)

func main() {
	/*
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			c.SendString("Welcome to SeaSlopes")
			return nil
		})

		err := app.Listen(":8080")
		if err != nil {
			log.Fatal("Failed to listen on port 8080")
		}

	*/

	fmt.Println(slope.ScrapeBlueMountain())

}
