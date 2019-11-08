package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func hello1() {
	fmt.Print("hello1 ")
}
func main() {
	defer hello()
	defer world()
	hello1()

}
