package main

import (
	"clima/key"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	data := fetchData()

	fmt.Print(data)
}

func fetchData() string {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?q=Vancouver&appid=" + key.Env())

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
