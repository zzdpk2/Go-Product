package main

import (
	"fmt"
	webv1 "goProduct/dummy-web/pkg/v1"
	"net/http"
)

func main() {
	server := webv1.NewSdkHttpServer("test-server")
	// server.Route("/", home)
	// server.Route("/user", user)
	// server.Route("/user/create", createUser)
	// POST
	server.Route(http.MethodGet, "/user/signup", SignUp)
	// server.Route("/order", order)
	err := server.Start(":8080")
	// er := errors.New("abc")
	if err != nil {
		panic(err)
	}
}

func SignUp(ctx *webv1.Context) {
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
