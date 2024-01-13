package main

import (
	"SeaSlope/sea"
	"SeaSlope/slope"
	"fmt"
	"log"
)

func main() {
	//Creating the app
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

	//Slope data

	data, _ := slope.ScrapeBlueMountain()
	fmt.Println("Slope Data:")
	fmt.Println(data)

	weatherResp := &slope.WeatherData{}
	err := slope.GetData(weatherResp)
	if err != nil {
		log.Fatal("Failed to get data")
	}

	fmt.Println(weatherResp.Data.Values.Temperature)

	fmt.Println("\n\nSea Data:")
	fmt.Println(sea.ScapeSurfData())

}
