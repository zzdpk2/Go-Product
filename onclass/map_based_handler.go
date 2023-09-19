package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handle HandlerFunc)
}

type Handler interface {
	// http.Handler
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// Key: method+url
	handlers map[string]func(ctx *Context)
}

// Implemented http.Handler
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not found!"))
	}
}

// Implemented Routable property
func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc HandlerFunc) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

// Confirm the HandlerBasedOnMap implemented Handler interface
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
