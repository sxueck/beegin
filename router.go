package beegin

import (
	"fmt"
	"log"
)

type router struct {
	handlers routerMap
}

func newRouter() *router {
	return &router{
		handlers: make(routerMap),
	}
}

func (r *router) addRoute(method string, pattern string, h HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = h
}

func (r *router) handle(c *Context) error {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		err := handler(c)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	return defaultErrorHandlerFunc(c)
}
