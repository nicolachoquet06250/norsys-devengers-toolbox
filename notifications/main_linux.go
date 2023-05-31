package notifications

import (
	"norsys-devengers-toolbox/notify"
)

func Notify(title, message, appIcon string, appId *string) error {
	return notify.Notify(title, message, appIcon, appId)
}
