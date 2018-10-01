package check

import "net/http"

type StatusCheck interface {
	Check() (ServiceStatus, interface{})
}

type HTTPStatusCheck struct {
	URL        string
	HTTPMethod string
}

type HTTPStatusCheckResult struct {
	StatusCheck HTTPStatusCheck
	Response    *http.Response
}

func (statusCheck HTTPStatusCheck) Check() (ServiceStatus, interface{}) {
	var response *http.Response
	switch statusCheck.HTTPMethod {
	case http.MethodGet:
		response, _ = http.Get(statusCheck.URL)
	case http.MethodPost:
		response, _ = http.Post(statusCheck.URL, "text/html", nil)
	}
	return mapStatus(response), HTTPStatusCheckResult{StatusCheck: statusCheck, Response: response}
}

func mapStatus(response *http.Response) ServiceStatus {
	httpStatus := response.StatusCode
	if httpStatus >= 200 && httpStatus < 500 {
		return Up
	} else {
		return Down
	}
}
