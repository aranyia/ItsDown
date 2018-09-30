package main

import (
	"ItsDown/slack"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type StatusData struct {
	Status string
	Color  string
}

var status = "Unknown"

var statusTemplate = template.Must(template.ParseFiles("status.html"))

func GetStatus(w http.ResponseWriter, req *http.Request) {
	var color string
	if strings.EqualFold(status, "UP") {
		color = "#00ff00"
	} else if strings.EqualFold(status, "DOWN") {
		color = "#ff0000"
	}

	statusTemplate.Execute(w, StatusData{status, color})
}

func SetStatus(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	status := params.Get("s")

	log.Println("status set to:", status)
	status = strings.ToUpper(status)

	slackWebHookURL, slackParamPresent := os.LookupEnv("slack-webhook-url")
	if slackParamPresent {
		slackWebHook := slack.WebHook{URL: slackWebHookURL}
		slackWebHook.PostMessage("*Integration point* is " + strings.ToUpper(status))
	}
	io.WriteString(w, "New integration point status is: "+status)
}

func main() {
	http.HandleFunc("/", GetStatus)
	http.HandleFunc("/set", SetStatus)

	http.ListenAndServe(":8080", nil)
}
