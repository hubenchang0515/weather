package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/hubenchang0515/weather"
)

type Config struct {
	Key  string
	City string
}

func cache(key *string, city *string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFile := path.Join(home, ".config", "weather.config")
	data, _ := ioutil.ReadFile(configFile)

	var config Config
	json.Unmarshal(data, &config)
	if *key != "" {
		config.Key = *key
	}

	if *city != "" {
		config.City = *city
	}

	data, _ = json.Marshal(&config)
	err = ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	*key = config.Key
	*city = config.City
}

func combine(w weather.DailyWeather) {
	if w.DayText != w.NightText {
		fmt.Printf("%s转%s ", w.DayText, w.NightText)
	} else {
		fmt.Printf("%s", w.DayText)
	}
}

func main() {
	key := flag.String("key", "", "Webhook key of seniverse.com, will be saved in ~/.config/weather.config")
	city := flag.String("city", "", "City, will be saved in ~/.config/weather.config")
	flag.Parse()
	cache(key, city)

	w := weather.Now(*key, *city)
	if w == nil {
		fmt.Printf("cannot get weather of '%s' with key '%s'", *city, *key)
		return
	}
	fmt.Printf("实时: %s %s℃\n", w.Text, w.Temperature)
	ws := weather.Forecast(*key, *city, 2)
	if len(ws) == 2 {
		fmt.Printf("今日: %s\n", ws[0].ToString())
		fmt.Printf("明日: %s\n", ws[1].ToString())
	}
}
