package main

import "fmt"

func main() {
	greeter("hello")
	greeter("world")
}

func greeter(s string) {
	fmt.Println(s)
}
