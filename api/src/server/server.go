package server

import (
	"net/http"
	"strings"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, p PathParams, h HandlerFunc)

type Server struct {
	CompiledRoutes []CompiledRoute
	Middleware     MiddlewareFunc
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var found bool = false
	var allowedMethods []string
	for _, route := range s.CompiledRoutes {
		match, paramMap := routeMatches(route, r.URL.Path)
		if match {
			found = true
			if route.Method != r.Method {
				allowedMethods = append(allowedMethods, route.Method)
				continue
			}
			if s.Middleware != nil {
				s.Middleware(w, r, paramMap, route.Handler)
			} else {
				route.Handler(w, r, paramMap)
			}
			return
		}
	}
	if found {
		w.Header().Add("Allow", strings.Join(allowedMethods, " "))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, "Path not found", http.StatusNotFound)
}
