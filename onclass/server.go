package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	// Route(method string, pattern string, handlerFunc func(ctx *Context))
	Routable
	Start(address string) error
}

// sdkHttpServer is implemented on http library
type sdkHttpServer struct {
	Name string
	// handler *HandlerBasedOnMap
	handler Handler
	root    Filter
}

// Route register route function
// func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
// 	http.HandleFunc(pattern,
// 		func(writer http.ResponseWriter, request *http.Request) {
// 			ctx := NewContext(writer, request)
// 			handleFunc(ctx)
// 		})
// }

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	// key := s.handler.key(method, pattern)
	// s.handler.handlers[key] = handleFunc
	// s.handler.Route(method, pattern, handleFunc)
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	// handler := &HandlerBasedOnMap{}
	// http.Handle("/", handler)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	var root Filter = handler.ServeHTTP

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

type signUpReq struct {
	// `json:"email"` : tag
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	// ctx := &Context{
	// 	W: w,
	// 	R: r,
	// }

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("Write failed: %v", err)
	}
}

// func NewServer() Server {
// 	return &sdkHttpServer{}
// }

type Factory func() Server

var factory Factory

func RegisterFactory(f Factory) {
	factory = f
}

func NewServer() Server {
	return factory()
}

// type Header map[string][]string
