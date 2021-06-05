package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	Path           string `json:"path"`
	Timezone       string `json:"timezone"`
	TimezoneOffset string `json:"timezone_offset"`
}

type Weather struct {
	Code        string `json:"code"`
	Text        string `json:"text"`
	Temperature string `json:"temperature"`
}

type WeatherResult struct {
	Location Location `json:"location"`
	Now      Weather  `json:"now"`
	Update   string   `json:"last_update"`
}

type WeatherResults struct {
	Results []WeatherResult
}

func Now(key string, city string) *Weather {
	url := fmt.Sprintf("https://api.seniverse.com/v3/weather/now.json?key=%s&location=%s&language=zh-Hans&unit=c", key, city)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results WeatherResults
	json.Unmarshal(data, &results)
	if len(results.Results) == 0 {
		return nil
	}
	return &results.Results[0].Now
}

type DailyWeather struct {
	Date            string `json:"date"`
	DayCode         string `json:"code_day"`
	DayText         string `json:"text_day"`
	NightCode       string `json:"code_night"`
	NightText       string `json:"text_night"`
	HighTemperature string `json:"high"`
	LowTemperature  string `json:"low"`
}

type DailyWeatherResult struct {
	Location Location       `json:"location"`
	Daily    []DailyWeather `json:"daily"`
}

type DailyWeatherResults struct {
	Results []DailyWeatherResult `json:"results"`
}

func Forecast(key string, city string, days uint) []DailyWeather {
	url := fmt.Sprintf("https://api.seniverse.com/v3/weather/daily.json?key=%s&location=%s&language=zh-Hans&unit=c&start=0&days=%d", key, city, days)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results DailyWeatherResults
	json.Unmarshal(data, &results)
	if len(results.Results) == 0 || len(results.Results[0].Daily) == 0 {
		return nil
	}
	return results.Results[0].Daily
}
