package main

//
// import (
// 	"fmt"
// 	"runtime"
// )
//
// var flag1 = 0
// var aa [2000]*int
// var bb *int
//
// func FuncA() {
// 	var tmp1 [2000]*int
// 	var tmp *int
//
// 	fmt.Println("Into dead loop!")
// 	// for flag1 == 0 {
// 	for {
// 		aa = tmp1
// 		bb = tmp
// 		// fmt.Println("1")
// 	}
// }
//
// func main() {
// 	for i := 0; i < 6; i++ {
// 		go func() {
// 			FuncA()
// 		}()
// 	}
// 	for i := 0; i < 4; i++ {
// 		runtime.GC()
// 	}
// }
//
// var count int = 0
//
// func init() {
// 	go checkCount()
// }
//
// //go:nowritebarrierrec
// func monitor() {
// 	systemstack(func() {
// 		newm(monitor, nil, -1)
// 	})
// }
