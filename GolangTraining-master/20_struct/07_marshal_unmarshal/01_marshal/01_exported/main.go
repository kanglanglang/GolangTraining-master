package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	//p1 := person{"James", "Bond", 20, 007}
	//fmt.Println(p1)
	//bs, _ := json.Marshal(p1)
	////fmt.Println(bs)
	////fmt.Printf("%T \n", bs)
	//fmt.Println(string(bs))



	p1 := person{"James", "Bond", 20,000}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}





