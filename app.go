package main

import (
	"./check"
	"./config"
	"log"
	"net/http"
)

func main() {
	services, _ := config.LoadServices()

	check.Scheduler{
		UpdateCycle: config.GetCheckIntervaql(),
		ToFire: func() {
			for _, service := range services {
				service.CheckStatus()
			}
		}}.Start()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
