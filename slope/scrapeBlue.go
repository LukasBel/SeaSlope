package slope

import (
	"SeaSlope/models"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
)

//1-11-2024: Implemented a wait group to account for data collection due to the asynchronous nature of web scraping

func ScrapeBlueMountain() (models.Conditions, error) {
	URL := "https://www.skibluemt.com/condition-report/"

	conditions := models.Conditions{}
	conditionList := []string{}
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

	c.OnHTML("p.has-text-align-center", func(element *colly.HTMLElement) {
		conditionList = append(conditionList, element.Text)
	})

	err := c.Visit(URL)
	if err != nil {
		return conditions, err
	}

	conditions.Date = conditionList[0]
	conditions.BaseDepth = conditionList[1]
	conditions.LiftsOperating = conditionList[2]
	conditions.OpenTrails = conditionList[3]
	conditions.TerrainParks = conditionList[4]
	conditions.Tubing = conditionList[5]

	wg.Wait()

	return conditions, nil
}
