package main

import "fmt"

func main() {
	var n News = fakeNews{
		Name: "hello",
	}
	n.Report()
}

type fakeNews = News

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}
