package app

import (
	"net/http"
	"strings"
)

type Router struct {
	routes   []*route
	NotFound http.Handler
}

func NewRouter() *Router {
	return &Router{
		NotFound: http.NotFoundHandler(),
	}
}

func (r *Router) Handle(method, pattern string, handler http.Handler) {
	route := &route{
		method:  strings.ToLower(method),
		pattern: pattern,
		handler: handler,
	}

	r.routes = append(r.routes, route)
}

func (r *Router) HandleFunc(method, pattern string, fn http.HandlerFunc) {
	r.Handle(method, pattern, fn)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := strings.ToLower(req.Method)
	pattern := req.URL.Path
	for _, route := range r.routes {
		if route.match(method, pattern) {
			route.handler.ServeHTTP(w, req)
			return
		}
	}
	r.NotFound.ServeHTTP(w, req)

}

type route struct {
	method  string
	pattern string
	handler http.Handler
}

func (r *route) match(method, pattern string) bool {
	if r.method != method {
		return false
	}

	if r.pattern != pattern {
		return false
	}

	return true
}
