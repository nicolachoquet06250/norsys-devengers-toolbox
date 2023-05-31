package history

import "norsys-devengers-toolbox/environement"

func GetCachePath() string {
	return environement.HomePath() + environement.Slash() + "norsys-devengers-toolbox"
}
