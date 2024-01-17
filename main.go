package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository struct {
	DB gorm.DB
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

	err = models.MigrateSpots(db)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetUpRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Welcome to SeaSlopes")
		return nil
	})

	err := app.Listen(":8080")
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
