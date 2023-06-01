package router

type AssetsJs string

var (
	InitAstilectron  AssetsJs = "/assets/js/init_astilectron.js"
	RedirectHandlers AssetsJs = "/assets/js/redirect_handlers.js"
	ToolboxJs        AssetsJs = "/assets/js/toolbox.js"
)

func AssetJsToString(route AssetsJs) string {
	return string(route)
}

type AssetsCss string

var (
	ToolboxCss   AssetsCss = "/assets/css/toolbox.css"
	AboutToolCss AssetsCss = "/assets/css/about-tool.css"
)

func AssetCssToString(route AssetsCss) string {
	return string(route)
}
