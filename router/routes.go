package router

type Route string

var (
	HomePage      Route = "/"
	LoaderPage    Route = "/loading"
	HelpPage      Route = "/help"
	AboutToolPage Route = "/about"
)

func RouteToString(route Route) string {
	return string(route)
}
