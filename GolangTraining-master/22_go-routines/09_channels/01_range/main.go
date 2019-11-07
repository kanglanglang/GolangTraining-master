package main

import (
	"fmt"
	"time"
)

func sample(c chan string){
	c <- "hello 1"
	c <- "hello 2"
	c <- "hello 3"
	c <- "hello 4"

}


func sample2(c chan string){
	time.Sleep(2*time.Second)
	str := <-c
	fmt.Println(str)
	fmt.Println(str)
	str=str+"11111"
	c <-str
	close(c)
}
func main() {
	c := make(chan string, 5)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	go sample(c)
	go sample2(c)
	time.Sleep(3*time.Second)
	for n := range c {
		fmt.Println(n)
	}
}
