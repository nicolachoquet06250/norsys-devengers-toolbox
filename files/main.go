package files

import (
	alertPkg "norsys-devengers-toolbox/alert"
)

type (
	FileSystem interface {
		Create() (alert alertPkg.Alert)
		Remove(id int) (alert alertPkg.Alert)
		Exists() (exists bool, err error)
	}

	Project struct {
		Path string `json:"path"`
		Name string `json:"name"`
	}
)
