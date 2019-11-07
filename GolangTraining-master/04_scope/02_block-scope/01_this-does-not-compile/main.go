package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	// no access to x
	// this does not compile
	x :=2222
	fmt.Println(x)
}
