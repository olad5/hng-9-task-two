package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Profile struct {
	SlackUserName string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int64  `json:"age"`
	Bio           string `json:"bio"`
}

var UserSlackUsername = os.Getenv("SLACK_USERNAME")

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.ParseInt(os.Getenv("AGE"), 10, 64)
	bio := os.Getenv("BIO")

	details := Profile{
		SlackUserName: UserSlackUsername,
		Age:           age,
		Backend:       true,
		Bio:           bio,
	}

	profileJson, err := json.Marshal(details)
	if err != nil {
		panic(err)
	}
	w.Write(profileJson)
}
