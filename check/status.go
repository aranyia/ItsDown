package check

import (
	"log"
	"net/http"
)

type StatusCheck interface {
	Check() (ServiceStatus, interface{})
}

type HTTPStatusCheck struct {
	URL        string `json:"url"`
	HTTPMethod string `json:"httpMethod"`
}

type HTTPStatusCheckResult struct {
	StatusCheck HTTPStatusCheck
	Response    *http.Response
}

func (statusCheck *HTTPStatusCheck) Check() (ServiceStatus, interface{}) {
	var response *http.Response
	var err error

	switch statusCheck.HTTPMethod {
	case http.MethodGet:
		response, err = http.Get(statusCheck.URL)
	case http.MethodPost:
		response, err = http.Post(statusCheck.URL, "text/html", nil)
	}

	if err == nil {
		return mapStatus(response), HTTPStatusCheckResult{StatusCheck: *statusCheck, Response: response}
	} else {
		log.Fatalln(err)
		return Down, nil
	}
}

func mapStatus(response *http.Response) ServiceStatus {
	httpStatus := response.StatusCode
	if httpStatus >= 200 && httpStatus < 500 {
		return Up
	} else {
		return Down
	}
}
