package modals

import (
	"log"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/history"
	"norsys-devengers-toolbox/router"
	"norsys-devengers-toolbox/window"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

var Modal *astilectron.Window

func CreateModal(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger) *astilectron.Window {
	Modal = window.CreateWindow(a, l, environement.UrlBase()+router.RouteToString(router. /*FolderSelectorPage*/ LoaderPage), &astilectron.WindowOptions{
		Center:          astikit.BoolPtr(true),
		Height:          astikit.IntPtr(700),
		Width:           astikit.IntPtr(700),
		Icon:            astikit.StrPtr(history.GetIconPath("light")),
		Show:            astikit.BoolPtr(false),
		Modal:           astikit.BoolPtr(true),
		Transparent:     astikit.BoolPtr(false),
		BackgroundColor: astikit.StrPtr("white"),
	})

	window.OpenDevTools(Modal, l)

	return Modal
}
