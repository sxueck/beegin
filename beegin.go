package beegin

import (
	"fmt"
	"net/http"
)

type (
	HandlerFunc func(*Context) error
	routerMap   map[string]HandlerFunc

	Engine struct {
		router                  *router
		DefaultErrorHandlerFunc HandlerFunc
	}
)

func New() *Engine {
	return &Engine{
		router: newRouter(),
		DefaultErrorHandlerFunc:
		defaultErrorHandlerFunc,
	}
}

func defaultErrorHandlerFunc(c *Context) error {
	return c.String(http.StatusNotFound, fmt.Errorf("404 Not Found"))
}

func (e *Engine) GET(pattern string, h HandlerFunc) {
	e.router.addRoute(http.MethodGet, pattern, h)
}

func (e *Engine) POST(pattern string, h HandlerFunc) {
	e.router.addRoute(http.MethodPost, pattern, h)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
