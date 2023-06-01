package router

type ApiRoute string

var (
	ToolsApi      ApiRoute = "/api/tools"
	CategoriesApi ApiRoute = "/api/categories"
)

func ApiRouteToString(route ApiRoute) string {
	return string(route)
}
