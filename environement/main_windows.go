package environement

import (
	"os/user"
	"strings"
)

func Slash() string {
	return "\\"
}

func PwdVar() string {
	start := "{"
	end := "}"

	return "$" + start + "pwd" + end
}

func RootPath() string {
	return "C:\\"
}

func HomePath() string {
	path := RootPath()
	username := func() string {
		u, _ := user.Current()
		return u.Username
	}()

	path += "Users" + Slash() + strings.Split(username, Slash())[1]

	return path
}

func BackLine() string {
	return "\r\n"
}
