package environement

import "os/user"

func Slash() string {
	return "/"
}

func PwdVar() string {
	start := "("
	end := ")"

	return "$" + start + "pwd" + end
}

func RootPath() string {
	return "/"
}

func HomePath() string {
	path := RootPath()
	username := func() string {
		u, _ := user.Current()
		return u.Username
	}()

	path += "home" + Slash() + username

	return path
}

func BackLine() string {
	return "\n"
}
