package fn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	weatherAPIURL = "http://api.weatherapi.com/v1/current.json"
)

func requestWeather(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	return response, nil
}

// GetWeather : to get weather by making request
func GetWeather(city string) (*Weather, error) {

	var weatherInfo Weather

	reqURL := fmt.Sprintf("%s?key=%s&q=%s&aqi=yes", weatherAPIURL, WeatherAPIKey, city)
	res, err := requestWeather(reqURL)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Unable to make request:" + err.Error())
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Unable to read response:" + err.Error())
		}

		err = json.Unmarshal(data, &weatherInfo)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Unable to unmarshal JSON:" + err.Error())
		}
		return &weatherInfo, nil
	}

	return nil, errors.New("Unexpected response")
}
