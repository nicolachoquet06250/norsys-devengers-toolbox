package notifications

import (
	"fmt"
	"log"
	"norsys-devengers-toolbox/alert"
	"norsys-devengers-toolbox/history"
	"strings"

	"github.com/asticode/go-astikit"
)

var GeneratedPreventionAlert *alert.Alert
var GeneratedAlert *alert.Alert

type NotificationOption struct {
	Title string
	Body  string
}

func CreateNotification(l *log.Logger, o NotificationOption) {
	err := Notify(
		o.Title, o.Body,
		history.GetIconPath("light"),
		astikit.StrPtr("Norsys Project Base Generator"),
	)
	if err != nil {
		if strings.Contains(err.Error(), "not compatible") {
			GeneratedPreventionAlert = alert.GenerateAlert(err.Error(), alert.ERROR)

			message := o.Title + "<br />" + o.Body
			var status alert.AlertStatus
			if strings.Contains(strings.ToLower(message), "succès") {
				status = alert.SUCCESS
			} else if strings.Contains(strings.ToLower(message), "erreur") {
				status = alert.ERROR
			} else {
				status = alert.INFO
			}

			GeneratedAlert = alert.GenerateAlert(message, status)
		} else {
			l.Fatal(fmt.Errorf("error : %s", err))
		}
	}
}
