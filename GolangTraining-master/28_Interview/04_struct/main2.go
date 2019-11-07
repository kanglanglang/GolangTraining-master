package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Human struct{}

func (h *Human) ShowA() {
	fmt.Println("Human showA")
}

type Teacher struct {
	Human
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
	//t.People.ShowA()
}

/*

如果嵌套有两个结构，并且两个结构具有相同的方法，如何执行的？
答案是 编译报错，不支持这种情况的。

*/