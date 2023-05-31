package window

import (
	"fmt"
	"log"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/history"
	"norsys-devengers-toolbox/router"
	"strconv"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	// "golang.org/x/exp/shiny/widget/theme"
)

func CreateWindow(a *astilectron.Astilectron, l *log.Logger, url string, options *astilectron.WindowOptions, name ...string) (w *astilectron.Window) {
	if len(name) == 0 {
		name = append(name, "main")
	}

	w, err := a.NewWindow(url, options)
	if err != nil {
		l.Fatal(fmt.Errorf("%s: new window failed: %w", name[0], err))
	}

	err = w.Create()
	if err != nil {
		l.Fatal(fmt.Errorf("%s: creating window failed: %w", name[0], err))
	}

	return
}

func CreateLoader(a *astilectron.Astilectron, l *log.Logger) (w *astilectron.Window) {
	l.Println("------------------------ LOADER ------------------------")
	baseUrl := "http://127.0.0.1:" + strconv.FormatInt(int64(environement.ChosenPort), 10)
	w = CreateWindow(
		a, l,
		baseUrl+router.RouteToString(router.LoaderPage),
		&astilectron.WindowOptions{
			Frame:          astikit.BoolPtr(false),
			Center:         astikit.BoolPtr(true),
			Width:          astikit.IntPtr(550 + 20),
			Height:         astikit.IntPtr(174 + 120),
			Resizable:      astikit.BoolPtr(false),
			Fullscreenable: astikit.BoolPtr(false),
			Icon:           astikit.StrPtr(history.GetIconPath("light")),
			Transparent:    astikit.BoolPtr(false),
		},
		"loader",
	)

	return
}

func OpenDevTools(w *astilectron.Window, l *log.Logger) {
	if environement.IsDevEnv() {
		err := w.OpenDevTools()
		if err != nil {
			l.Fatal(fmt.Errorf("erreur : %s", err.Error()))
		}
	}
}
