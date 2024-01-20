package main

import (
	"SeaSlope/models"
	"SeaSlope/sea"
	"SeaSlope/slope"
	"SeaSlope/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetSeaData(c *fiber.Ctx) error {
	seaDataChan := make(chan models.Forecast)

	go func() {
		data, _ := sea.ScapeSurfData()
		seaDataChan <- data
	}()

	seaData := <-seaDataChan

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Sea Data Fetched Successfully", "data": seaData})
	return nil
}

// This works but perhaps measure performance and refactor to use channel
func (r *Repository) GetSlopeData(c *fiber.Ctx) error {
	data, _ := slope.ScrapeBlueMountain()
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Sea Data Fetched Successfully", "data": data})

	weatherResp := &models.WeatherData{}
	err := slope.GetData(weatherResp)
	if err != nil {
		log.Fatal("Failed to get data")
	}

	fmt.Println(weatherResp.Data.Values.Temperature)
	return nil
}

//Get GeneralData func

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/SeaSlope")
	api.Get("/Sea", r.GetSeaData)
	api.Get("/Slope", r.GetSlopeData)
}

func main() {
	//Creating the app
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	//err = models.MigrateSpots(db)
	//if err != nil {
	//	log.Fatal("Failed to migrate database")
	//}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Welcome to SeaSlopes")
		return nil
	})

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080")
	}

	//Slope data

	//data, _ := slope.ScrapeBlueMountain()
	//fmt.Println("Slope Data:")
	//fmt.Println(data)
	//
	//weatherResp := &slope.WeatherData{}
	//err := slope.GetData(weatherResp)
	//if err != nil {
	//	log.Fatal("Failed to get data")
	//}
	//
	//fmt.Println(weatherResp.Data.Values.Temperature)
	//
	////Sea Data
	//
	//fmt.Println("\n\nSea Data:")
	//fmt.Println(sea.ScapeSurfData())

}
