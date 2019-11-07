package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	foo()
	bar()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Nanoseconds())
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
