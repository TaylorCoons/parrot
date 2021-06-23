package server

import (
	"testing"
)

func TestCompileRoutes(t *testing.T) {
	compiledRoutes := CompileRoutes([]Route{
		{"GET", "/first/:firstId/second/:secondId/path", nil},
		{"GET", "/third/:thirdId/fourth/:fourthId/path", nil},
	})
	if len(compiledRoutes) != 2 {
		t.Error("Failed to return the same number of compiled routes as routes")
	}
}

func TestCompileRoute(t *testing.T) {
	compiledRoute := compileRoute(Route{"GET", "/first/:firstId/second/:secondId/path", nil})
	if compiledRoute.Method != "GET" {
		t.Error("Method set incorrectly")
	}
	if compiledRoute.PathMatcher.String() != "^/first/(?P<firstId>[^/]*)/second/(?P<secondId>[^/]*)/path$" {
		t.Error("Incorrect regex set")
	}
}
