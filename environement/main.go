package environement

import (
	"os"
	"strconv"
	"strings"
)

var ChosenPort int

func CharBack() string {
	return "\n"
}

func Tab() string {
	return "\t"
}

func IsBuild() bool {
	return os.Getenv("GOBUILD") != "1"
}

func UrlBase() string {
	port := strconv.FormatInt(int64(ChosenPort), 10)

	return "http://127.0.0.1:" + port
}

func IsDevEnv() bool {
	return os.Getenv("OPEN_DEVTOOLS") == "1" &&
		strings.Contains(os.Args[0], Slash()+"b001"+Slash()+"exe"+Slash())
}
