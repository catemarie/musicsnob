package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"bufio"
	"strings"
)

var keys []string

type Event struct {
	Artist string `json:"artist"`
	Date string `json:"date"`
	Venue string `json:"venue"`
}

type Search struct {
	Genre string `json:"genre"`
	Location string `json:"location"`
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	var searchData Search

	var searchResData []Event
	//var sampleEvent Event
	//sampleEvent.Artist = "Andrew Rayel"
	//sampleEvent.Date = "11/27/2021"
	//sampleEvent.Venue = "Academy"
	//searchResData = append(searchResData, sampleEvent)

    decoder := json.NewDecoder(r.Body)

    decoder.Decode(&searchData)
    fmt.Println(searchData.Genre)
    fmt.Println(searchData.Location)

    city := strings.Replace(searchData.Location, " ", "%20", -1)

    locId := getLocId("California", city, keys[0])
    searchResData = getArtists(locId, keys[0], keys[1], keys[2])

    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(searchResData); err != nil {
        panic(err)
    }
}

func readKeys(fname string) {
	file, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        keys = append(keys, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
	http.HandleFunc("/", homePageHandler)

	readKeys("keys.txt")

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}