package gui

import (
	_ "embed"
	"fmt"
	"log"
	"norsys-devengers-toolbox/decoder"
	"norsys-devengers-toolbox/history"
	"norsys-devengers-toolbox/modals"
	"norsys-devengers-toolbox/receiver"
	"norsys-devengers-toolbox/router"
	"norsys-devengers-toolbox/server"
	"norsys-devengers-toolbox/window"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/atotto/clipboard"
)

func Process() {
	logger := log.New(GetLoggerWriter(), log.Prefix(), log.Flags())

	go server.Process(true, logger)

	GenerateIcon()

	app := CreateApp("Devengers Toolbox", logger)

	defer app.Close()

	urlBase := UrlBase()

	var Loader = window.CreateLoader(app, logger)

	var Window = window.CreateWindow(
		app, logger,
		urlBase,
		&astilectron.WindowOptions{
			Center:    astikit.BoolPtr(true),
			Height:    astikit.IntPtr(500),
			Width:     astikit.IntPtr(450),
			Icon:      astikit.StrPtr(history.GetIconPath("light")),
			Show:      astikit.BoolPtr(false),
			Resizable: astikit.BoolPtr(false),
		},
	)

	err := Window.Show()
	logger.Println("show main Window")
	if err != nil {
		logger.Fatal(fmt.Errorf("main Window can't be showed"))
	} else {
		err = Loader.Destroy()
		if err != nil {
			logger.Fatal(fmt.Errorf("fail to destroy loader %s", err))
		}
	}

	err = app.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("Toollist"),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				var err = Window.SendMessage(
					decoder.NewMessage(decoder.Redirect, map[string]string{"uri": router.RouteToString(router.HomePage)}),
				)
				if err != nil {
					log.Fatal(fmt.Errorf("error : %s", err.Error()))
				}
				return
			},
			Accelerator: astilectron.NewAccelerator("Control", "T"),
		},
		{
			Label: astikit.StrPtr("Aide"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Aide"),
					Role:  astilectron.MenuItemRoleHelp,
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						var err = Window.SendMessage(
							decoder.NewMessage(decoder.Redirect, map[string]string{"uri": router.RouteToString(router.HelpPage)}),
						)
						if err != nil {
							log.Fatal(fmt.Errorf("error : %s", err.Error()))
						}
						return
					},
					Accelerator: astilectron.NewAccelerator("Control", "H"),
				},
				{
					Label:       astikit.StrPtr("Fermer"),
					Role:        astilectron.MenuItemRoleClose,
					Accelerator: astilectron.NewAccelerator("Control", "Q"),
				},
			},
		},
	}).Create()
	if err != nil {
		logger.Fatal(fmt.Errorf("fail to create menu %s", err))
	}

	OpenDevTools(Window, logger)

	Window.On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
		err = Window.Show()
		logger.Println("show main Window")
		if err != nil {
			logger.Fatal(fmt.Errorf("main Window can't be showed"))
		}

		err = Loader.Destroy()
		if err != nil {
			logger.Println(fmt.Errorf("loader can't be destroyed"))
		}
		return
	})

	Window.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		if Loader != nil {
			err = Loader.Destroy()
			if err != nil {
				logger.Fatal(fmt.Errorf("fail to destroy loader %s", err))
			}
		}

		if modals.Modal != nil {
			err = modals.Modal.Destroy()
			if err != nil {
				logger.Fatal(fmt.Errorf("fail to destroy modal %s", err))
			}
		}
		return
	})

	Window.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decoder.JsonMessage(
			decoder.Message(m),
		)

		// fmt.Printf("%v", jsonMessage)

		switch jsonMessage.Channel {
		// case string(decoder.Notification):
		// 	receiver.NotificationChannel(logger, &jsonMessage)
		case string(decoder.DestroyLoader):
			receiver.DestroyLoaderChannel(Loader, logger)
		case string(decoder.Clipboard):
			err := clipboard.WriteAll(jsonMessage.Data["text"])
			if err != nil {
				println("erreur")
			}
			Window.SendMessage(decoder.NewMessage(decoder.Clipboard, map[string]string{
				"success": "true",
			}))
		}

		return
	})

	defer (func() {
		// recursive := func()  {
		if Loader == nil {
			return
		}
		err = Loader.Destroy()
		if err != nil {
			logger.Fatal(fmt.Errorf("fail to destroy loader %s", err))
		}
		// }
	})()

	// Blocking pattern
	app.Wait()
}
