package main

import "fmt"

func main() {
	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello, panic: %v\n", data)
		}
		fmt.Println("恢复之后从这里继续执行")
	}()

	panic("Boom")
	// 这里不会执行
	fmt.Println("这里将不会执行下来")
}
