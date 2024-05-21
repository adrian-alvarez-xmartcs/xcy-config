package router

import (
	"net/http"
	"strings"

	"xcylla.io/common/log"
)

var logging log.Logger

type Router struct {
	http.ServeMux
	routes map[string]RouterContext
}

func NewRouter() *Router {
	logging = log.NewLogger("Router")

	logging.Trace("Initializing router")

	rtr := &Router{
		routes: make(map[string]RouterContext),
	}

	return rtr
}

func (rtr *Router) addRoute(method string, path string, handlerFunc func(ctx RouterContext)) {
	key := method + " " + path
	rtr.routes[key] = RouterContext{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerFunc(RouterContext{Request: r, Writer: w})
		}),
	}
}

func (rtr *Router) GET(path string, handlerFunc func(ctx RouterContext)) {
	rtr.addRoute("GET", path, handlerFunc)
}

func (rtr *Router) POST(path string, handlerFunc func(ctx RouterContext)) {
	rtr.addRoute("POST", path, handlerFunc)
}

func (rtr *Router) PUT(path string, handlerFunc func(ctx RouterContext)) {
	rtr.addRoute("PUT", path, handlerFunc)
}

func (rtr *Router) DELETE(path string, handlerFunc func(ctx RouterContext)) {
	rtr.addRoute("DELETE", path, handlerFunc)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	key := r.Method + " " + r.URL.Path
	if handler, exists := rtr.routes[key]; exists {
		handler.ServeHTTP(w, r)
		return
	}

	if strings.Contains(r.URL.Path, "static") {
		http.ServeFile(w, r, "public/"+r.URL.Path)
		return
	}

	http.ServeFile(w, r, "public/")
}

func (rtr *Router) Start(port string) error {
	logging.Info("Starting server on port %s", port)
	err := http.ListenAndServe(port, rtr)
	if err != nil {
		return err
	}
	return nil
}
