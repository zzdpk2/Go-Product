package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	server := NewHttpServer("test-server")
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

// type Server interface {
// 	Route(pattern string, handlerFunc http.HandlerFunc)
// 	Start(address string) error
// }
//
// type sdkHttpServer struct {
// 	Name string
// }
