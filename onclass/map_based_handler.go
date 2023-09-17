package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handle func(ctx *Context))
}

type Handler interface {
	http.Handler
	Routable
}

type HandlerBasedOnMap struct {
	// Key: method+url
	handlers map[string]func(ctx *Context)
}

// Implemented http.Handler
func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter,
	req *http.Request) {
	key := h.key(req.Method, req.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		ctx := NewContext(writer, req)
		handler(ctx)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("Not found!"))
	}
}

// Implemented Routable property
func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
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
