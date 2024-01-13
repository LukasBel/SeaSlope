package sea

type Forecast struct {
	Date      int    `json:"date"`
	Weather   string `json:"weather"`
	Tide      string `json:"tide"`
	Buoy      string `json:"buoy"`
	WaterTemp string `json:"waterTemp"`
	Report    string `json:"report"`
}
