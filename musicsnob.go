package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strings"
)

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

    decoder := json.NewDecoder(r.Body)

    decoder.Decode(&searchData)
    fmt.Println(searchData.Genre)
    fmt.Println(searchData.Location)

    city := strings.Replace(searchData.Location, " ", "%20", -1)

    locId := getLocId("California", city)
    searchResData = getArtists(locId)

    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(searchResData); err != nil {
        panic(err)
    }
}

func main() {
	http.HandleFunc("/", homePageHandler)
	port := os.Getenv("PORT")


	fmt.Println("EDMTRAIN:", os.Getenv("EDMTRAIN"))
	fmt.Println("SPOTIFYID:", os.Getenv("SPOTIFYID"))
	fmt.Println("SPOTIFYSECRET:", os.Getenv("SPOTIFYSECRET"))

	fmt.Println("Server listening on port " + ":" + port)
	log.Panic(
		http.ListenAndServe(":" + port, nil),
	)
}