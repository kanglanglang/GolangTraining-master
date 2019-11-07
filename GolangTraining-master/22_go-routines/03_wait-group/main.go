package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t1 := time.Now()
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Nanoseconds())
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
