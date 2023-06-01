package router

import (
	_ "embed"
	"net/http"
	"norsys-devengers-toolbox/parser"
)

//go:embed templates/assets/js/init_astilectron.js
var assetJsInitAstilectron string

//go:embed templates/assets/js/redirect_handlers.js
var assetJsRedirectHandlers string

//go:embed templates/assets/js/toolbox.js
var assetJsToolbox string

//go:embed templates/assets/css/toolbox.css
var assetCssToolbox string

//go:embed templates/assets/css/about-tool.css
var assetCssAboutTool string

func AssetJsInitAstilectron(w http.ResponseWriter, r *http.Request) {
	parser.Js(&w, assetJsInitAstilectron)
}

func AssetJsRedirectHandlers(w http.ResponseWriter, r *http.Request) {
	parser.Js(&w, assetJsRedirectHandlers)
}

func AssetJsToolbox(w http.ResponseWriter, r *http.Request) {
	parser.Js(&w, assetJsToolbox)
}

func AssetCssToolbox(w http.ResponseWriter, r *http.Request) {
	parser.Css(&w, assetCssToolbox)
}

func AssetCssAboutTool(w http.ResponseWriter, r *http.Request) {
	parser.Css(&w, assetCssAboutTool)
}
