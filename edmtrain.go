package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "github.com/tidwall/gjson"
    "strconv"
    "fmt"
    "golang.org/x/oauth2/clientcredentials"
    "github.com/zmb3/spotify"
    "context"
    "strings"
)

func getArtistGenres(artist string, id string, secret string) []string {

    // handle credentials
    config := &clientcredentials.Config{
        ClientID:     id,
        ClientSecret: secret,
        TokenURL:     spotify.TokenURL,
    }
    token, err := config.Token(context.Background())
    if err != nil {
        log.Fatalf("couldn't get token: %v", err)
    }
    client := spotify.Authenticator{}.NewClient(token)

    // search for this artist and get the top result
    results, err := client.Search(artist, spotify.SearchTypeArtist)
    if err != nil {
        log.Fatal(err)
    }

    // list all of the artist's genres
    var g []string
    if results.Artists != nil {
        artists := results.Artists.Artists
        if len(artists) > 0 {
            for _, genre := range artists[0].Genres {
                g = append(g, genre)
            }
        }
    }
    
    return g
}


func getJson(url string) string {

    // API call
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Couldn't fetch results")
    }

    // grab the JSON
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Couldn't read results")
    }

    return string(body)

}
func getLocId(state string, city string, key string) string {
    
    url := "https://edmtrain.com/api/locations?state=" + state + "&city=" + city + "&client=" + key
    body := getJson(url)

    // get the location ID
    value := gjson.Get(string(body), "data.#.id")
    locId := value.Array()[0].String()

    return locId
}

func getArtists(loc_id string, key string, id string, secret string) []Event {

    var a []Event

    url := "https://edmtrain.com/api/events?locationIds=" + loc_id + "&client=" + key
    body := getJson(url)

    json_txt := gjson.Get(string(body), "data")
    test := json_txt.Array()
    num := len(test)

    for ii := 0; ii < num; ii++ {

        date := gjson.Get(string(body), "data." + strconv.Itoa(ii) + ".date")
        artistList := gjson.Get(string(body), "data." + strconv.Itoa(ii) + ".artistList.#.name")
        venue:= gjson.Get(string(body), "data." + strconv.Itoa(ii) + ".venue.name")

        for _, n := range artistList.Array() {
            artist := n.Array()
            if len(artist) > 0 {
                genres := getArtistGenres(artist[0].String(), id, secret)
                for _, g := range genres {
                    if (strings.Contains(g, "trance")) {
                        fmt.Println(artist[0].String())
                        entry := Event{artist[0].String(), date.String(), venue.String()}
                        a = append(a, entry)
                        break
                    }
                }
            }
        }
    }

    return a
}