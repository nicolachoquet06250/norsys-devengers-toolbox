package gui

import (
	"fmt"
	"log"
	"norsys-devengers-toolbox/alert"
	"norsys-devengers-toolbox/files"
	"norsys-devengers-toolbox/history"
	"norsys-devengers-toolbox/router"

	"github.com/asticode/go-astilectron"
)

func init() {
	alert.GeneratedPreventionAlert = new(alert.Alert)
	alert.GeneratedAlert = new(alert.Alert)
}

func GenerateIcon() {
	_icon := files.NewFile(history.GetIconPath("light"))

	exists, _ := _icon.Exists()
	if !exists {
		err := _icon.Create(router.IconLight, true)
		if err != nil {
			log.Fatal(fmt.Errorf("can't create icon light locally"))
		}
	}

	_icon = files.NewFile(history.GetIconPath("dark"))

	exists, _ = _icon.Exists()
	if !exists {
		err := _icon.Create(router.IconDark, true)
		if err != nil {
			log.Fatal(fmt.Errorf("can't create icon dark locally"))
		}
	}
}

func CreateApp(name string, l *log.Logger) (a *astilectron.Astilectron) {
	// Initialize astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           name,
		BaseDirectoryPath: "gui",
	})

	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}

	// Add a listener on Astilectron
	a.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		l.Println("App has crashed")
		return
	})

	// Handle signals
	a.HandleSignals()

	// Start
	err = a.Start()
	if err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	return
}

func GenerateAlert(message string, status alert.AlertStatus) *alert.Alert {
	a := alert.NewAlert(message, status)
	return &a
}
