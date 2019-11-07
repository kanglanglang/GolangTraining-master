package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	if val, kk := myGreeting[2];kk{
		delete(myGreeting, 2)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", kk)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", kk)
	}

	fmt.Println(myGreeting)
}
