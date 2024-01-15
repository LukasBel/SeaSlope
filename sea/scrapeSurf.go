package sea

import (
	"SeaSlope/handlers"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
	"time"
)

func ScapeSurfData() (Forecast, error) {
	URL := "https://surfcaptain.com/forecast/atlantic-city-new-jersey"

	forecast := Forecast{}
	forecastList := []string{}
	var wg sync.WaitGroup

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting:", request.URL)
		wg.Add(1)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Status:", response.StatusCode)
		wg.Done()
	})

	c.OnHTML(".fcst-current-header", func(element *colly.HTMLElement) {
		forecastList = append(forecastList, element.Text)
	})

	c.OnHTML(".current-data-info", func(element *colly.HTMLElement) {
		forecastList = append(forecastList, element.Text)
	})

	err := c.Visit(URL)
	if err != nil {
		return forecast, err
	}

	forecastList = handlers.CleanText(forecastList) //filters out white space from data

	forecast.Date = time.Now().Day()
	forecast.Report = forecastList[0]
	forecast.Weather = forecastList[1]
	forecast.Tide = forecastList[2]
	forecast.Buoy = forecastList[3]
	forecast.WaterTemp = forecastList[4]

	wg.Wait()

	return forecast, err
}
