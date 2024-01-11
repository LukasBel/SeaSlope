package slope

import (
	"fmt"
	"github.com/gocolly/colly"
)

func ScrapeBlueMountain() (*Conditions, error) {
	URL := "https://www.skibluemt.com/condition-report/"

	conditions := &Conditions{}
	conditionList := []string{}

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting:", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Status:", response.StatusCode)
	})

	c.OnHTML("p.has-text-align-center", func(element *colly.HTMLElement) {
		conditionList = append(conditionList, element.Text)
		fmt.Println(conditionList)

	})

	fmt.Println(conditionList)

	//conditions.Date = conditionList[0]
	//conditions.BaseDepth = conditionList[1]
	//conditions.LiftsOperating = conditionList[2]
	//conditions.OpenTrails = conditionList[3]
	//conditions.TerrainParks = conditionList[4]
	//conditions.Tubing = conditionList[5]

	err := c.Visit(URL)
	if err != nil {
		return conditions, err
	}

	return conditions, err
}
