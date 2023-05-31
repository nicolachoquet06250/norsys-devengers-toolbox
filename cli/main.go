package cli

import (
	"fmt"
	"log"
	"norsys-devengers-toolbox/errors"
	"os/exec"
	"runtime"
	"strings"
)

func ExeCmd(cmd string) []byte {
	parts := strings.Fields(cmd)

	out, err := exec.Command(parts[0], parts[1]).Output()
	errors.MaybeError(err, func(err error) *int64 {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
		var r *int64
		return r
	})

	return out
}

type Browser struct{}

func (_ Browser) Open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	errors.MaybeError(err, func(err error) *int64 {
		log.Fatal(err)
		var r *int64
		return r
	})
}
