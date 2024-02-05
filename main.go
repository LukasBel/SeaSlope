package main

import (
	"SeaSlope/models"
	"SeaSlope/sea"
	"SeaSlope/slope"
	"SeaSlope/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		data, err := sea.ScapeSurfData()
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed To Fetch Sea Data"})
			log.Fatal("Sea Data GoRoutine Failed")
		}
		seaDataChan <- data
	}()

	seaData := <-seaDataChan
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Sea Data Fetched Successfully", "data": seaData})

	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Sea Data Fetched Successfully", "data": seaData})
}

func (r *Repository) GetSlopeData(c *fiber.Ctx) error {
	slopeDataChan := make(chan models.Conditions)

	go func() {
		data, err := slope.ScrapeBlueMountain()
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed To Fetch Slope Data"})
			log.Fatal("Slope Data GoRoutine Failed")
		}
		slopeDataChan <- data
	}()

	slopeData := <-slopeDataChan
	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Slope Data Fetched Successfully", "data": slopeData})
}

func (r *Repository) GetSlopeWeather(c *fiber.Ctx) error {
	weatherResp := &models.WeatherData{}
	done := make(chan struct{})
	go func() {
		defer close(done)

		err := slope.GetData(weatherResp)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed To Fetch Slope Weather"})
			log.Fatal("Failed to get data")
		}
	}()

	<-done
	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Slope Weather Fetched Successfully", "data": &weatherResp})
}

//func (r *Repository) SaveSeaData(c *fiber.Ctx) error {
//	err := r.DB.Create(&models.Forecast{})
//	if err != nil {
//		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Failed To Save Data"})
//		log.Fatal("Failed to get data")
//	}
//
//	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Sea Data Saved Successfully"})
//	return nil
//}

//Get GeneralData func

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/SeaSlope")
	api.Get("/Sea", r.GetSeaData)
	api.Get("/Slope", r.GetSlopeData)
	api.Get("/Slope/Weather", r.GetSlopeWeather)
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

	app.Use(cors.New())

	//app.Get("/", func(c *fiber.Ctx) error {
	//	c.SendString("Welcome to SeaSlopes")
	//	//Trying something new
	//	return c.SendStatus(200) //Everything's ok
	//})

	app.Get("/api/data", r.GetSeaData)

	//app.Get("/api/message", func(c *fiber.Ctx) error {
	//	message := "Hello from Golang Backend!"
	//	return c.JSON(fiber.Map{"message": message})
	//})

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080")
	}

}
