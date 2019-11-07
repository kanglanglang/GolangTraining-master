package main

import "fmt"

type People struct{}

func  (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()

}

/*
result  :::
```
showA
showB
解析：
```
首先明确一点 go中没有继承关系。也不应该提及“继承”这个词，其中Trecher并没有继承Propler，而是嵌套People,
而t.ShowA()是一个语法糖，其实t.ShowA() = t.people.ShowA(),也就是说在嵌套结构中，go会优先调用本身方法，
如果本身没有此方法，就回去调用其所包含结构的方法。
本题中，showA()是Teacher不具有的，但是它所嵌套的People具有，因此回调用People.showA(),People.showA()
中调用了*People 的showB()当然会展示“shwoB”，而不是“teacher showB”
*/