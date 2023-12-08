package main

import "fmt"

func Foo(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Foo(1, 2))
}
