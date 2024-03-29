package main

import "fmt"

//
//func increment() int {
//	x++
//	return x
//}
//
//func main() {
//	fmt.Println(increment())
//	fmt.Println(increment())
//}



func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func  test() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

}
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	test()
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
