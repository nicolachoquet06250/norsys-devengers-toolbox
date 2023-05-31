package alert

type AlertStatus string

const (
	SUCCESS AlertStatus = "success"
	ERROR   AlertStatus = "danger"
	INFO    AlertStatus = "info"
)

var (
	GeneratedPreventionAlert *Alert
	GeneratedAlert           *Alert
)

type Alert struct {
	Message string
	Type    string
}

func NewAlert(message string, status AlertStatus) Alert {
	return Alert{
		Message: message,
		Type:    string(status),
	}
}

func GenerateAlert(message string, status AlertStatus) *Alert {
	a := NewAlert(message, status)
	return &a
}
