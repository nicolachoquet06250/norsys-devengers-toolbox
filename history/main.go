package history

import (
	"norsys-devengers-toolbox/environement"
)

func GetIconPath(_type string) string {
	return GetCachePath() + environement.Slash() + (func() string {
		if _type == "light" {
			return "logo-devengers-light.png"
		}
		return "logo-devengers-dark.png"
	})()
}
