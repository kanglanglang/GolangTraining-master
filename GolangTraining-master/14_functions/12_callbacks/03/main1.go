package main

import "fmt"

func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	addName("lamar", func(nm string) {
		fmt.Printf("hi may name is %v \n", nm)
	})
}
