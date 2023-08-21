package helper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func Helper() {
	fmt.Print("Hello worlds\n")
}

func getTimeFromTimestamp(t int64) time.Time {
	t, err := strconv.ParseInt(fmt.Sprint(t), 10, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tm := time.Unix(t, 0)

	return tm
}

func GetDay(t int64) int {
	parsedTime := getTimeFromTimestamp(t)

	day := parsedTime.Day()
	return day
}

func GetHourFromTimestamp(t int64) int {
	parsedTime := getTimeFromTimestamp(t)

	hour := parsedTime.Hour()
	return hour
}

func GetWeekday(t int64) string {
	parsedTime := getTimeFromTimestamp(t)

	weekday := parsedTime.Weekday().String()

	return weekday
}

func Float64ToString(i float64) string {
	return fmt.Sprintf("%.0f°C", i)
}

func ColourWeather(i string) string {
	bg_yellow := color.New(color.FgWhite, color.BgRed, color.Bold).SprintfFunc()
	cyan := color.New(color.BgGreen, color.Underline, color.Bold).SprintfFunc()

	if i == "Rain" {
		return "         " + bg_yellow(i)
	} else if i == "Clouds" {
		return "       " + cyan(i)
	}

	return i
}
