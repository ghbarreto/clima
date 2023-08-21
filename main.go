package main

import (
	"clima/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
)

var FORMAT string = "%v %13s %13s %13s %13s %13s\n"
var ARGS []string = os.Args[1:]

func main() {

	if len(ARGS) > 0 && ARGS[0] == "--help" {
		fmt.Println("Usage: clima [t | w | h | up | setup | new]")
		fmt.Println("=====================================================================")
		fmt.Println("t: get today's weather")
		fmt.Println("w: get the weather for the next 7 days [default]")
		fmt.Println("h: show this help message")
		fmt.Println("up: force update the weather data (cached daily)")
		fmt.Println("setup: setup your API key & country & city")
		fmt.Println("new: change country & city")

		os.Exit(0)
	}

	runMock()
}

func runMock() {
	jsonFile, err := os.Open("res.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Weather
	json.Unmarshal([]byte(byteValue), &result)

	result.formatReturn()
}

func (item Weather) header() {
	location := time.Now().In(time.FixedZone("Timezone", item.City.Timezone)).Format("15:04 - 02/01/2006")
	info := color.New(color.FgHiBlue, color.Bold).SprintfFunc()

	fmt.Printf("( %v ) | %v \n  \n", info(item.City.Name+","+item.City.Country), info(location))

	fmt.Printf(FORMAT, "Temp", "Clouds", "Temp Max", "Temp Min", "Feels Like", "Date")
}

func (item Weather) formatReturn() {
	// cyan := color.New(color.FgCyan).SprintfFunc()
	bg_yellow := color.New(color.FgMagenta, color.Bold).SprintfFunc()

	get_today := false

	// print header
	item.header()

	if len(ARGS) > 0 && ARGS[0] == "t" {
		get_today = true
	}

	for _, item := range item.List {
		temperature := item.Main.Temp
		tempf := fmt.Sprintf("%.0f°C", temperature)
		weather := item.Weather[0].Main
		hour := helper.GetHourFromTimestamp(item.Dt) + 1

		if get_today && helper.GetDay(item.Dt) != time.Now().Day() {
			continue
		}

		if hour == 6 || hour == 12 || hour == 18 {
			if temperature < 19 {
				fmt.Printf(FORMAT, helper.Float64ToString(temperature), weather, helper.ColourWeather(weather), weather, helper.Float64ToString(temperature), helper.GetWeekday(item.Dt))
			} else {
				fmt.Printf(FORMAT, bg_yellow(tempf), weather, helper.ColourWeather(weather), helper.ColourWeather(weather), helper.Float64ToString(temperature), helper.GetWeekday(item.Dt))
			}

		}

	}

}

// func fetchData() string {
// 	response, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?q=Vancouver&appid=" + key.Env() + "&units=metric")

// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}

// 	data, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return string(data)
// }
