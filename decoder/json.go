package decoder

import (
	"encoding/base64"
	jsonPkg "encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/asticode/go-astilectron"
)

type Channel string

const (
	Redirect                Channel = "Redirect"
	Notification            Channel = "Notification"
	ChooseFolder            Channel = "ChooseFolder"
	PutFolder               Channel = "PutFolder"
	OpenFolderSelectorModal Channel = "OpenFolderSelectorModal"
	OpenFolder              Channel = "OpenFolder"
	GetTree                 Channel = "GetTree"
	DestroyLoader           Channel = "DestroyLoader"
	ShowAlert               Channel = "ShowAlert"
	Clipboard               Channel = "Clipboard"
	OpenLink                Channel = "OpenLink"
)

type JsonMessageType struct {
	Channel string            `json:"channel"`
	Data    map[string]string `json:"data"`
}

func NewMessage(channel Channel, data map[string]string) JsonMessageType {
	return JsonMessageType{
		Channel: string(channel),
		Data:    data,
	}
}

func Message(m *astilectron.EventMessage) []byte {
	json, err := m.MarshalJSON()
	if err != nil {
		log.Fatal(fmt.Errorf("error : %s", err.Error()))
	}

	json, err = base64.StdEncoding.DecodeString(strings.Replace(string(json), "\"", "", 2))
	if err != nil {
		log.Fatal(fmt.Errorf("error : %s", err.Error()))
	}

	return json
}

func JsonMessage(message []byte) (jsonMessage JsonMessageType) {
	err := jsonPkg.Unmarshal(message, &jsonMessage)
	if err != nil {
		log.Fatal(fmt.Errorf("error : %s", err.Error()))
	}

	return jsonMessage
}

type JsonArrayMessage struct {
	Channel string         `json:"channel"`
	Data    map[string]any `json:"data"`
}

func NewArrayMessage(channel Channel, data map[string]any) JsonArrayMessage {
	return JsonArrayMessage{
		Channel: string(channel),
		Data:    data,
	}
}
