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

func TestRouteMatch_ValidMatch(t *testing.T) {
	route := compileRoute(Route{"GET", "/first/:id/second/:id2/value", nil})
	match, params := routeMatches(route, "/first/123/second/abc/value")
	if match != true {
		t.Error("Paths should match")
	}
	if params["id"] != "123" {
		t.Error("First path param is not set")
	}
	if params["id2"] != "abc" {
		t.Error("Second path param is not set")
	}
}

func TestRouteMatch_InvalidMatch(t *testing.T) {
	var route CompiledRoute
	var match bool

	route = compileRoute(Route{"GET", "/first/:id", nil})
	match, _ = routeMatches(route, "/first")
	if match == true {
		t.Error("Route should not match")
	}
	match, _ = routeMatches(route, "/first/123/second")
	if match == true {
		t.Error("Route should not match")
	}
	match, _ = routeMatches(route, "/")
	if match == true {
		t.Error("Route should not match")
	}
}
