package main

import "fmt"

func math(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World")
	fmt.Println(`Hello
World`)

	a := 10
	b := 20
	c := math(a, b)
	fmt.Println(c)
}
