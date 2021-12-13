package main

import (
	"flag"
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

var Host = flag.String("host", "w7de211b.us-east-1.emqx.cloud", "server hostname or IP")
var Port = flag.Int("port", 15301, "server port")
var Username = flag.String("username", "emqx", "username")
var Password = flag.String("password", "public", "password")

// store temperature's data
var currentTemp = temperature{min: -1, max: -1, latest: -1}

func minTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"value\": %f}", currentTemp.min)
}

func maxTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"value\": %f}", currentTemp.max)
}

func latestTempHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "{\"value\": %f}", currentTemp.latest)
}

func updateLatestTemperature(data []byte) {
	latest, err := strconv.ParseFloat(string(data), 32)
	if err != nil {
		log.Printf("Error update temperature: %v \n", err)
		return
	}

	if currentTemp.max == -1 || latest > currentTemp.max {
		currentTemp.max = latest
	}
	if currentTemp.min == -1 || latest < currentTemp.min {
		currentTemp.min = latest
	}

	currentTemp.latest = latest

	fmt.Printf("Received latest temperature %f. Current temperature %+v \n", latest, currentTemp)
}

func main() {
	flag.Parse()
	config := mymqtt.Config{Host: *Host, Port: *Port, Username: *Username, Password: *Password}

	// subscribe mqtt to receive min/max/latest temperature
	mymqtt := &mymqtt.MyMqtt{}
	mymqtt.Connect(config)
	mymqtt.Subscribe("temperature/latest", updateLatestTemperature)

	// setup http rest api endpoint
	http.HandleFunc("/sensor/min", minTempHandler)
	http.HandleFunc("/sensor/max", maxTempHandler)
	http.HandleFunc("/sensor/latest", latestTempHandler)

	// start http server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
