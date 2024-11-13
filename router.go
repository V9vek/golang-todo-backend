package main

import "net/http"

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *router {
	router := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
	return router
}

func (r *router) AddRoute(method, path string, handler http.HandlerFunc) {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if methodRoutes, exists := r.routes[req.Method]; exists {
		if handler, found := methodRoutes[req.URL.Path]; found {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
