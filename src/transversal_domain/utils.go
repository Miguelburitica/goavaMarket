package transversal_domain

import "net/http"

func GetGroupedRoutes(routes ...Route) Routes {
	return Routes(routes)
}

func GetQueryParamWithDefault(r *http.Request, param, defaultValue string) string {
	value := r.URL.Query().Get(param)
	if value == "" {
		return defaultValue
	}
	return value
}
