package gui

import (
	"fmt"
	"io"
	"log"
	"norsys-devengers-toolbox/environement"
	"os"
	"strconv"
	"strings"

	"github.com/asticode/go-astilectron"
)

func UrlBase() string {
	port := strconv.FormatInt(int64(environement.ChosenPort), 10)

	return "http://127.0.0.1:" + port
}

func IsDevEnv() bool {
	return os.Getenv("OPEN_DEVTOOLS") == "1" &&
		strings.Contains(os.Args[0], environement.Slash()+"b001"+environement.Slash()+"exe"+environement.Slash())
}

func OpenDevTools(w *astilectron.Window, l *log.Logger) {
	if IsDevEnv() {
		err := w.OpenDevTools()
		if err != nil {
			l.Fatal(fmt.Errorf("erreur : %s", err.Error()))
		}
	}
}

func GetLoggerWriter() io.Writer {
	if IsDevEnv() {
		return NewLogWriter().Enable().Writer
	}
	return NewLogWriter().Disable().Writer
}
