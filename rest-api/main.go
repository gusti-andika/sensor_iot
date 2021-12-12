package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gusti-andika/sensor_iot/rest-api/mymqtt"
)

type temperature struct {
	min    float64
	max    float64
	latest float64
}

// store temperature's data
var currentTemp = temperature{min: -1, max: -1, latest: -1}

func minTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"min\": %f}", currentTemp.min)
}

func maxTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"max\": %f}", currentTemp.max)
}

func latestTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"latest\": %f}", currentTemp.latest)
}

func main() {

	// mqtt message listener
	updateLatestTemperature := func(data []byte) {
		latest, err := strconv.ParseFloat(string(data), 32)
		if err != nil {
			log.Printf("Error update temperature: %v", err)
			return
		}

		if currentTemp.max == -1 || latest > currentTemp.max {
			currentTemp.max = latest
		}
		if currentTemp.min == -1 || latest < currentTemp.min {
			currentTemp.min = latest
		}

		currentTemp.latest = latest

		fmt.Printf("Received latest temperature %f. Current temperature %+v", latest, currentTemp)
	}

	// subscribe mqtt to receive min/max/latest temperature
	mymqtt := &mymqtt.MyMqtt{}
	mymqtt.Init()
	mymqtt.Subscribe("temperature/latest", updateLatestTemperature)

	// setup http rest api endpoint
	http.HandleFunc("/sensor/min", minTempHandler)
	http.HandleFunc("/sensor/max", maxTempHandler)
	http.HandleFunc("/sensor/latest", latestTempHandler)

	// start http server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
