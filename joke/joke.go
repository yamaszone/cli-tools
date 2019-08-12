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
	//Id      string `json:"id"`
	Joke     string `json:"joke"`
	//Status string `json:"status"`
}

func main() {
	url := "https://icanhazdadjoke.com"
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
	fmt.Println("Joke:\t\t", joke.Joke)
	//fmt.Println("Punchline:\t", joke.Punchline)
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
	request.Header.Set("Accept", "application/json")

	response, getErr := jokeClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}
	return response
}
