package fn

import (
	"fmt"
)

// Weather : struct to hold json response
type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		TempF       float64 `json:"temp_f"`
		WindKph     float64 `json:"wind_kph"`
		Humidity    int     `json:"humidity"`
		AirQuality  struct {
			Pm25 float64 `json:"pm2_5"`
			Pm10 float64 `json:"pm10"`
		} `json:"air_quality"`
	} `json:"current"`
}

// ToString : to convert struct to string
func ToString(w *Weather) string {

	return fmt.Sprintf("Location \n Name : %s\n Country   : %s\n Localtime : %s\nCurrent\n LastUpdated : %s\n TempC  : %f\n TempF  : %f\n WindKph  : %f\n Humidity : %d\n AirQuality :\n   Pm25 : %f\n   Pm10 : %f\n",
		w.Location.Name, w.Location.Country, w.Location.Localtime,
		w.Current.LastUpdated, w.Current.TempC, w.Current.TempF,
		w.Current.WindKph, w.Current.Humidity, w.Current.AirQuality.Pm25,
		w.Current.AirQuality.Pm10)
}
