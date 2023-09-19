package main

import (
	"fmt"
	"time"
)

type HandlerFunc func(c *Context)

// 实现责任链的AOP
type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

// type Filter1 interface {
// 	Filter(c *Context)
// }
//
// type Server1 struct {
// 	filters []Filter1
// }

// type Interceptor interface {
// 	Before(c *Context)
// 	After(c *Context)
// 	Surrounding(c *Context)
// 	OnError()
// 	OnResponse()
// 	OnReturn()
// }

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("Spent %d nanoseconds", end-start)
	}
}
