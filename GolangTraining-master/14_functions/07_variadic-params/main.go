package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	fmt.Printf("%T \n", sf)
	var total float64
	fmt.Println(len(sf))
//https://zhidao.baidu.com/question/1990317293498053027.html
	for i, value := range sf {
		fmt.Printf("i=%v, value=%v\n", i, value)
	}

	for _, v := range sf {


		total += v
	}
	return total / float64(len(sf))
}
