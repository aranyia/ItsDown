package config

import (
	"../check"
	"../slack"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CheckInterval   = "interval"
	Services        = "services"
	SlackWebHookURL = "slack-webhook-url"
)

const checkIntervalDefault = 10

func GetCheckIntervaql() time.Duration {
	checkInterval, checkIntervalParamPresent := os.LookupEnv(CheckInterval)
	if checkIntervalParamPresent {
		checkIntervalInSec, err := strconv.Atoi(checkInterval)
		if err != nil {
			log.Fatal("check interval parse error", err)
		}
		return time.Duration(checkIntervalInSec) * time.Second
	} else {
		return checkIntervalDefault * time.Second
	}
}

func LoadServices() ([]check.Service, error) {
	servicesFileParam, servicesParamPresent := os.LookupEnv(Services)
	if !servicesParamPresent {
		log.Fatal("services configuration not provided")
	}

	var err error
	jsonFile, err := os.Open(servicesFileParam)
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("failed to parse configuration for services", err)
	}

	onChangeActions := loadOnStatusChangeActions()

	services := make([]check.Service, 0)
	for _, service := range services {
		service.OnChangeActions = onChangeActions
	}
	return services, json.Unmarshal(jsonBytes, &services)
}

func loadOnStatusChangeActions() []func(status check.ServiceStatus, service check.Service) {
	onStatusChangeActions := make([]func(status check.ServiceStatus, service check.Service), 0)

	slackWebHookURL, slackParamPresent := os.LookupEnv(SlackWebHookURL)
	if slackParamPresent {
		slackWebHook := slack.WebHook{URL: slackWebHookURL}

		onStatusChangeActions = append(onStatusChangeActions, func(status check.ServiceStatus, service check.Service) {
			slackWebHook.PostMessage("*" + service.Name + "* is " + strings.ToUpper(status.String()))
		})
	}
	return onStatusChangeActions
}
