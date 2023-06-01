package router

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"norsys-devengers-toolbox/api"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/parser"
	"strconv"
)

//go:embed templates/index.html
var home string

//go:embed templates/loader.html
var load string

//go:embed templates/help.html
var help string

//go:embed templates/about-tool.html
var aboutTool string

func Home(w http.ResponseWriter, r *http.Request) {
	result, err := parser.ShowHtmlPage(parser.Page{
		CurrentPage: "index",
		Template:    home,
		Title: parser.Title{
			Tab: "Devengers Toolbox",
		},
		CssFiles: parser.CssFiles{
			AssetCssToString(ToolboxCss),
		},
		Vars: &map[string]interface{}{
			"Scripts": []struct{ Script string }{
				{Script: AssetJsToString(InitAstilectron)},
				{Script: AssetJsToString(RedirectHandlers)},
				{Script: AssetJsToString(ToolboxJs)},
			},
			"Logo":  base64.StdEncoding.EncodeToString([]byte(IconDark)),
			"Tools": r.RequestURI,
		},
	}, "")

	if err != nil {
		log.Fatal(err.Error())
	}

	parser.Text(&w, result)
}

func Loader(w http.ResponseWriter, r *http.Request) {
	result, _ := parser.ShowHtmlPage(parser.Page{
		CurrentPage: "load",
		Template:    load,
		CssFiles:    parser.CssFiles{},
		MetaData:    parser.MetaData{},
		Title: parser.Title{
			Tab:  "Chargement...",
			Page: "Chargement...",
		},
		Vars: &map[string]interface{}{
			"Image": base64.StdEncoding.EncodeToString([]byte(IconDark)),
			"Uri":   r.RequestURI,
		},
	}, "")

	parser.Text(&w, result)
}

func Help(w http.ResponseWriter, r *http.Request) {
	result, _ := parser.ShowHtmlPage(parser.Page{
		CurrentPage: "help",
		Template:    help,
		CssFiles:    parser.CssFiles{},
		MetaData:    parser.MetaData{},
		Title: parser.Title{
			Tab:  "Help",
			Page: "Help",
		},
		Vars: &map[string]interface{}{
			"Scripts": []struct{ Script string }{
				{Script: AssetJsToString(InitAstilectron)},
				{Script: AssetJsToString(RedirectHandlers)},
			},
			"Uri": r.RequestURI,
		},
	}, "")

	parser.Text(&w, result)
}

func AboutTool(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:" + strconv.FormatInt(int64(environement.ChosenPort), 10) + "/api/tools?id=" + r.URL.Query()["id"][0])

	result := ""
	if err == nil {
		body, err := io.ReadAll(resp.Body)

		if err == nil {
			var tool api.Tool
			err = json.Unmarshal(body, &tool)

			if err != nil {
				println(err.Error())
			}

			result, _ = parser.ShowHtmlPage(parser.Page{
				CurrentPage: "about-tool",
				Template:    aboutTool,
				CssFiles: parser.CssFiles{
					AssetCssToString(AboutToolCss),
				},
				MetaData: parser.MetaData{},
				Title: parser.Title{
					Tab:  "À propos de",
					Page: "À propos de",
				},
				Vars: &map[string]interface{}{
					// "Scripts": []struct{ Script string }{
					// 	{Script: AssetJsToString(InitAstilectron)},
					// 	{Script: AssetJsToString(RedirectHandlers)},
					// },
					"Tool": tool,
				},
			}, "")
		}
	}

	parser.Text(&w, result)
}
