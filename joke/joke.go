package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Joke struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	url := "https://08ad1pao69.execute-api.us-east-1.amazonaws.com/dev/random_joke"
	/*
	   response, err := http.Get(API_URL)
	   response := `{"id":3, "type": "general", "setup": "New setup", "punchline": "New punchline"}`
	   response_bytes := []byte(response)
	   err := json.Unmarshal(response_bytes, &joke)
	*/
	response := GetJoke(url)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	joke := Joke{}
	jsonErr := json.Unmarshal(body, &joke)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}
	fmt.Println("Setup:\t\t", joke.Setup)
	fmt.Println("Punchline:\t", joke.Punchline)
}

func GetJoke(url string) *http.Response {
	jokeClient := http.Client{
		Timeout: time.Second * 5,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "jokes-cli")

	response, getErr := jokeClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}
	return response
}
