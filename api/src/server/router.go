package server

import (
	"net/http"
	"regexp"
	"strings"
)

type PathParams map[string]string

type HandlerFunc func(w http.ResponseWriter, r *http.Request, p PathParams)

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

type CompiledRoute struct {
	Method      string
	PathMatcher *regexp.Regexp
	Handler     HandlerFunc
}

func CompileRoutes(routes []Route) []CompiledRoute {
	compiledRoutes := make([]CompiledRoute, len(routes))
	for i, route := range routes {
		compiledRoutes[i] = compileRoute(route)
	}
	return compiledRoutes
}

func compileRoute(route Route) CompiledRoute {
	re := regexp.MustCompile(":[^/]*")
	components := strings.Split(route.Path, "/")
	for i, component := range components {
		match := re.FindString(component)
		if match != "" {
			components[i] = "(?P<" + match[1:] + ">[^/]*)"
		}
	}
	path := "^" + strings.Join(components, "/") + "$"
	pathMatcher := regexp.MustCompile(path)
	return CompiledRoute{Method: route.Method, PathMatcher: pathMatcher, Handler: route.Handler}
}

func routeMatches(r CompiledRoute, u string) (bool, PathParams) {
	res := r.PathMatcher.FindStringSubmatch(u)
	var match bool = false
	pathParams := make(map[string]string)
	if len(res) > 0 {
		match = true
		pathNames := r.PathMatcher.SubexpNames()[1:]
		pathValues := res[1:]
		for i := range pathNames {
			pathParams[pathNames[i]] = pathValues[i]
		}
	}
	return match, pathParams
}
