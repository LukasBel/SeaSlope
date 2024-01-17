package slope

import (
	"SeaSlope/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetData(weatherResp *models.WeatherData) error {
	url := "https://api.tomorrow.io/v4/weather/realtime?location=blue%20mountain%20resort&apikey=DcIy5mM8YHNwH94Ow4FsQTLZjKaoMcxo"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	//weatherResp := &WeatherData{}
	//var weatherResp WeatherData
	body, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(body, weatherResp)
	if err != nil {
		log.Fatal("Error decoding JSON", err)
		return err
	}
	return nil
	//return weatherResp.Data.Values.PrecipitationProbability, weatherResp.Data.Values.Temperature
}
