package alerting

type Alert struct {
	Err error
}

func NewAlert(err error) *Alert {
	return &Alert{
		Err: err,
	}
}
