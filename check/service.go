package check

import "log"

type ServiceStatus uint8

const (
	Unknown ServiceStatus = 0
	Down    ServiceStatus = 1
	Up      ServiceStatus = 2
)

func (status ServiceStatus) String() string {
	switch status {
	case 0:
		return "unknown"
	case 1:
		return "down"
	case 2:
		return "up"
	default:
		return ""
	}
}

type Service struct {
	Name            string        `json:"name"`
	Status          ServiceStatus `json:"status"`
	StatusCheck     StatusCheck
	OnChangeActions []func(status ServiceStatus, service Service)
}

func (service *Service) CheckStatus() {
	newStatus, _ := service.StatusCheck.Check()
	log.Println(service.Name+" check HTTP status:", newStatus)

	previousStatus := service.Status
	service.Status = newStatus

	if newStatus != previousStatus {
		service.onStatusChange(previousStatus)
	}
}

func (service *Service) onStatusChange(previous ServiceStatus) {
	log.Println(service.Name + " service status changed from " + previous.String() + " to " + service.Status.String())

	if previous == Unknown && service.Status == Up {
		log.Println("no action taken as service confirmed " + Up.String())
		return
	} else {
		for _, fun := range service.OnChangeActions {
			fun(service.Status, *service)
		}
	}
}
