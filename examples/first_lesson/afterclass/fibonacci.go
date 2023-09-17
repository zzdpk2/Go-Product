package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(fibonacci(i), ",")
	}

}

func fibonacci(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}
