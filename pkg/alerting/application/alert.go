package alerting

import (
	"coffee-machine-dz/pkg/alerting/domain/alerting"
	"fmt"
)

type AlertingService struct{}

// New
func NewAlertingService() *AlertingService {
	return &AlertingService{}
}

type AlertingSvc interface {
	// Alert the err
	Alert(err error)
}

func (as *AlertingService) Alert(err error) {
	alert := alerting.NewAlert(err)
	fmt.Println(alert.Err.Error())
}
