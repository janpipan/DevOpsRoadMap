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
	Type    string                 `json:"type"`
	Repo    Repo                   `json:"repo"`
	Payload map[string]interface{} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
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

func getUserEvents(user string) ([]Event, error) {
	// check if github user exists
	if err := checkGitHubUser(user); err != nil {
		// fmt.Print(err.Error())
		// os.Exit(1)
		return nil, err
	}

	// create url
	url := fmt.Sprintf("https://api.github.com/users/%s/events", user)

	// make http get request
	response, err := http.Get(url)

	if err != nil {
		return nil, err
		// fmt.Print(err.Error())
		// os.Exit(1)
	}
	// read response body
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
		// log.Fatal(err)
	}

	// parse json data
	var responseJson []Event
	json.Unmarshal(responseData, &responseJson)

	return responseJson, nil

}

func printEventMessage(event Event) string {
	payload := event.Payload
	switch event.Type {
	case "CommitCommentEvent":
		return fmt.Sprintf("Create a comment %s in %s repo", payload["comment"], event.Repo.Name)
	case "CreateEvent":
		if payload["ref_type"] == "branch" {
			return fmt.Sprintf("Created a branch %s in repo %s", payload["ref"], event.Repo.Name)
		} else if payload["ref_type"] == "tag" {
			return fmt.Sprintf("Created a tag %s in repo %s", payload["ref"], event.Repo.Name)
		} else {
			return fmt.Sprintf("Created a repository %s", event.Repo.Name)
		}
	case "DeleteEvent":
		if payload["ref_type"] == "branch" {
			return fmt.Sprintf("Deleted a branch %s in repo %s", payload["ref"], event.Repo.Name)
		} else if payload["ref_type"] == "tag" {
			return fmt.Sprintf("Deleted a tag %s in repo %s", payload["ref"], event.Repo.Name)
		}
		return ""
	case "ForkEvent":
		return fmt.Sprintf("Forked a repository %s", event.Repo.Name)
	case "GollumEvent":
		return "GollumEvent"
	case "IssueCommentEvent":
		return "IssueCommentEvent"
	case "IssuesEvent":
		return "IssuesEvent"
	case "MemberEvent":
		return "MemberEvent"
	case "PublicEvent":
		return "PublicEvent"
	case "PullRequestEvent":
		return "PullRequestEvent"
	case "PullRequestReviewEvent":
		return "PullRequestReviewEvent"
	case "PullRequestReviewCommentEvent":
		return "PullRequestReviewCommentEvent"
	case "PullRequestReviewThreadEvent":
		return "PullRequestReviewThreadEvent"
	case "PushEvent":
		if size, ok := payload["size"].(float64); ok {
			return fmt.Sprintf("Pushed %d commits to %s", int(size), event.Repo.Name)
		}
		return ""
	case "ReleaseEvent":
		return "ReleaseEvent"
	case "SponsorshipEvent":
		return "SponsorshipEvent"
	case "WatchEvent":
		return fmt.Sprintf("Starred %s repo", event.Repo.Name)
	default:
		return "Unknown activity"
	}
}

func main() {

	// check if user set input arg
	if len(os.Args) != 2 {
		log.Fatal("GitHub user is missing")
	}

	user := os.Args[1]

	events, err := getUserEvents(user)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for _, event := range events {
		fmt.Println(printEventMessage(event))

	}
}
