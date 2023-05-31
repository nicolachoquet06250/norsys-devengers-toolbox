package history

import "norsys-devengers-toolbox/environement"

func GetCachePath() string {
	return environement.RootPath() + "norsys-devengers-toolbox"
}
