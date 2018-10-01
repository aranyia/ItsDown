package main

import (
	"ItsDown/check"
	"ItsDown/slack"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	onStatusChangeActions := make([]func(status check.ServiceStatus, service check.Service), 0)

	slackWebHookURL, slackParamPresent := os.LookupEnv("slack-webhook-url")
	if slackParamPresent {
		slackWebHook := slack.WebHook{URL: slackWebHookURL}

		onStatusChangeActions = append(onStatusChangeActions, func(status check.ServiceStatus, service check.Service) {
			slackWebHook.PostMessage(service.Name + " is " + strings.ToUpper(status.String()))
		})
	}

	serviceCNN := check.Service{Name: "CNN.com"}
	serviceCNN.StatusCheck = check.HTTPStatusCheck{"http://cnn.com", http.MethodGet}
	serviceCNN.OnChangeActions = onStatusChangeActions

	serviceGoogle := check.Service{Name: "Google"}
	serviceGoogle.StatusCheck = check.HTTPStatusCheck{"http://google.com", http.MethodGet}
	serviceGoogle.OnChangeActions = onStatusChangeActions

	scheduler := check.Scheduler{UpdateCycle: 10 * time.Second}
	scheduler.ToFire = func() {
		serviceCNN.CheckStatus()
		serviceGoogle.CheckStatus()
	}
	scheduler.Start()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
