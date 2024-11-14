package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"todo-backend/controller"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(todoController *controller.TodoController) *router {
	router := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	// api endpoints
	router.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	router.AddRoute("GET", "/api/todos", todoController.FindAll)
	router.AddRoute("POST", "/api/todo", todoController.Create)
	router.AddRoute("GET", "/api/todo/:todoId", todoController.FindById)
	router.AddRoute("DELETE", "/api/todo/:todoId", todoController.Delete)
	router.AddRoute("PATCH", "/api/todo/:todoId", todoController.Update)

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
		for path, handler := range methodRoutes {
			params, matched := matchRoute(path, req.URL.Path)
			if matched {
				// Attach params to request context
				ctx := context.WithValue(req.Context(), "params", params)
				handler(w, req.WithContext(ctx))
				return
			}
		}
	}
	http.NotFound(w, req)
}

func matchRoute(routePath, reqPath string) (map[string]string, bool) {
	// /api/todo/:todoId -> [api, todo, :todoId]
	routeParts := strings.Split(routePath, "/")
	// /api/todo/1 -> [api, todo, 1]
	reqParts := strings.Split(reqPath, "/")

	if len(routeParts) != len(reqParts) {
		return nil, false
	}

	// params["todoId"] = "1"
	params := make(map[string]string)

	for i, part := range routeParts {
		if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			params[paramName] = reqParts[i]
		} else if part != reqParts[i] {
			return nil, false
		}
	}

	return params, true
}
