package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func dummyHandler(w http.ResponseWriter, r *http.Request, p PathParams) {
	w.Write([]byte("AWK"))
}
func TestHTTPServer_ValidRoute(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/unique/path/123/value", nil)

	routes := CompileRoutes([]Route{
		{"GET", "/unique/path/:id/value", dummyHandler},
	})

	server := Server{CompiledRoutes: routes}
	server.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Error("HTTPServer did not return an OK status on valid request")
	}
	if string(body) != "AWK" {
		t.Error("HTTPServer did not return the proper body for the request")
	}
}

func TestHTTPServer_MethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/myPath", nil)

	routes := CompileRoutes([]Route{
		{"GET", "/myPath", dummyHandler},
	})

	server := Server{CompiledRoutes: routes}
	server.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Error("HTTPServer did not reject method not allowed correctly")
	}
	if string(body) != "Method not allowed\n" {
		t.Error("HTTPServer did not provide error for method not allowed")
	}
}

func TestHTTPServer_PathNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/wrongPath", nil)

	routes := CompileRoutes([]Route{
		{"GET", "/myPath", dummyHandler},
	})

	server := Server{CompiledRoutes: routes}
	server.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusNotFound {
		t.Error("HTTPServer did not reject path not found correctly")
	}
	if string(body) != "Path not found\n" {
		t.Error("HTTPServer did not provide error for path not found")
	}
}
