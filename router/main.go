package router

import (
	_ "embed"
	"net/http"
	httpPort "norsys-devengers-toolbox/http"
	"strconv"
)

//go:embed logo-devengers-dark.png
var IconDark string

//go:embed logo-devengers-light.png
var IconLight string

func Routes() {
	http.HandleFunc(RouteToString(HomePage), Home)
	http.HandleFunc(RouteToString(HelpPage), Help)
	http.HandleFunc(RouteToString(LoaderPage), Loader)
	http.HandleFunc(RouteToString(AboutToolPage), AboutTool)

	http.HandleFunc(AssetJsToString(InitAstilectron), AssetJsInitAstilectron)
	http.HandleFunc(AssetJsToString(RedirectHandlers), AssetJsRedirectHandlers)
	http.HandleFunc(AssetJsToString(ToolboxJs), AssetJsToolbox)
	http.HandleFunc(AssetCssToString(ToolboxCss), AssetCssToolbox)
	http.HandleFunc(AssetCssToString(AboutToolCss), AssetCssAboutTool)

	http.HandleFunc(ApiRouteToString(ToolsApi), ApiToolsPage)
	http.HandleFunc(ApiRouteToString(CategoriesApi), ApiCategoriesPage)

	http.ListenAndServe("0.0.0.0:"+strconv.FormatInt(int64(httpPort.ChooseUnusedPort()), 10), nil)
}
