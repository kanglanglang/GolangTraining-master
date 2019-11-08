package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for z := 0; z < len(greeting); z++ {
		//fmt.Println(greeting[z])
		fmt.Printf(greeting[z])
	}

	//fmt.Println(len(greeting))
	//
	//for j := 0; j < len(greeting); j++ {
	//	fmt.Println(greeting[j])
	//}

}
