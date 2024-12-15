package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Event struct {
	Type    string      `json:"type"`
	Repo    Repo        `json:"repo"`
	Payload interface{} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Actitivity struct {
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func checkGitHubUser(user string) error {

	url := fmt.Sprintf("https://api.github.com/users/%s", user)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// read response body
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject ErrorResponse
	json.Unmarshal(responseData, &responseObject)

	if responseObject.Message == "Not Found" {
		return &UserNotFound{}
	}

	return nil
}

func getActivities(events []Event)

func main() {

	// check if github user exists
	if err := checkGitHubUser("janpipan"); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// make http get request
	response, err := http.Get("https://api.github.com/users/janpipan/events")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// read response body
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Event
	json.Unmarshal(responseData, &responseObject)

	for _, event := range responseObject {
		fmt.Println(event.Payload.(map[string]interface{})[""])
	}
}
