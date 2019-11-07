package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

//考点：defer执行顺序，panic的执行方式
//结果： 打印后 打印中 “触发异常” 打印前
//解析：“触发异常”可能出现在任何位置。defer是先进后出，逆序执行