package server

import (
	"fmt"
	"net/http"
	"net/url"
)

type PathParams map[string]string

type HandlerFunc func(w http.ResponseWriter, r *http.Request, u url.URL, p PathParams)

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

// type CompiledRoute struct {
// 	Method      string
// 	PathMatcher regexp.Regexp
// 	Handler     HandlerFunc
// }

type Server struct {
	Routes []Route
}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I'm called")
}
