package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	go foo()
	go bar()
	fmt.Println(time.Now())
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
