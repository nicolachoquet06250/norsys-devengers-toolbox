package receiver

import (
	"fmt"
	"log"
	"norsys-devengers-toolbox/alert"
	"norsys-devengers-toolbox/decoder"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/history"
	"norsys-devengers-toolbox/modals"
	"norsys-devengers-toolbox/notifications"
	"norsys-devengers-toolbox/router"
	"norsys-devengers-toolbox/window"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func NotificationChannel(l *log.Logger, message *decoder.JsonMessageType) {
	body := fmt.Sprintf("Bonjour\n%s", message.Data["name"])
	if message.Data["body"] != "" {
		body = message.Data["body"]
	}

	title := "Test notif"
	if message.Data["title"] != "" {
		title = message.Data["title"]
	}

	println("NOTIFICATION : " + title + ", " + body)

	notifications.CreateNotification(l, notifications.NotificationOption{
		Title: title,
		Body:  body,
	})
}

func ChooseFolderChannel(w *astilectron.Window, l *log.Logger, message *decoder.JsonMessageType, main *astilectron.Window) {
	folderPath := message.Data["path"]

	notifications.CreateNotification(l, notifications.NotificationOption{
		Title: "Répertoire choisis",
		Body:  folderPath,
	})

	if w != nil {
		err := w.Destroy()
		if err != nil {
			l.Fatal(fmt.Errorf("fail to destroy window %s", err))
		}

		err = main.SendMessage(
			decoder.NewMessage(decoder.PutFolder, map[string]string{
				"folder": folderPath,
			}),
		)
		if err != nil {
			l.Fatal(fmt.Errorf("fail to send message to main window %s", err))
		}
	}
}

func OpenFolderSelectorModalChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger) *astilectron.Window {
	modals.Modal = window.CreateWindow(a, l, environement.UrlBase()+router.RouteToString(router. /*FolderSelectorPage*/ LoaderPage), &astilectron.WindowOptions{
		Center:          astikit.BoolPtr(true),
		Height:          astikit.IntPtr(700),
		Width:           astikit.IntPtr(700),
		Icon:            astikit.StrPtr(history.GetIconPath("light")),
		Show:            astikit.BoolPtr(false),
		Modal:           astikit.BoolPtr(true),
		Transparent:     astikit.BoolPtr(false),
		BackgroundColor: astikit.StrPtr("white"),
	})

	window.OpenDevTools(modals.Modal, l)

	err := modals.Modal.Show()
	if err != nil {
		l.Fatal(fmt.Errorf("fail to show modal %s", err))
	}

	modals.Modal.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decoder.JsonMessage(
			decoder.Message(m),
		)

		switch jsonMessage.Channel {
		case string(decoder.ChooseFolder):
			ChooseFolderChannel(modals.Modal, l, &jsonMessage, w)
		case string(decoder.OpenFolder):
			go OpenFolderChannel(modals.Modal, l, &jsonMessage)
		}

		if alert.GeneratedPreventionAlert != nil {
			err = w.SendMessage(decoder.NewMessage(decoder.ShowAlert, map[string]string{
				"message": alert.GeneratedPreventionAlert.Message,
				"type":    string(alert.GeneratedPreventionAlert.Type),
			}))
			if err != nil {
				log.Println(fmt.Errorf("main window send message error : %s", err))
			}
		}

		if alert.GeneratedAlert != nil {
			err = w.SendMessage(decoder.NewMessage(decoder.ShowAlert, map[string]string{
				"message": alert.GeneratedAlert.Message,
				"type":    string(alert.GeneratedAlert.Type),
			}))
			if err != nil {
				log.Println(fmt.Errorf("main window send message error : %s", err))
			}
		}

		return
	})

	return modals.Modal
}

func OpenFolderChannel(w *astilectron.Window, l *log.Logger, message *decoder.JsonMessageType) {
	if message.Data["folder"] == "" {
		message.Data["folder"] = environement.HomePath()
	}

	files, err := os.ReadDir(message.Data["folder"])
	if err != nil {
		log.Fatal(err)
	}

	var tree []string

	for _, f := range files {
		if f.IsDir() {
			tree = append(tree, f.Name())
		}
	}

	err = w.SendMessage(
		decoder.NewArrayMessage(decoder.GetTree, map[string]interface{}{
			"basePath":      message.Data["folder"],
			"tree":          tree,
			"pathSeparator": environement.Slash(),
			"isHome":        message.Data["folder"] == environement.HomePath(),
		}),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("fail to send message to window %s", err))
	}
}

func DestroyLoaderChannel(w *astilectron.Window, l *log.Logger) {
	if w != nil {
		err := w.Destroy()

		m := "destroy loader window"
		if err != nil {
			m = "loader window can't be destroy because doesn't exists"
		}

		l.Println(m)
		l.Println("------------------------ END LOADER ------------------------")
	}
}
