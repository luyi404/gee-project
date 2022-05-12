package gee

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
