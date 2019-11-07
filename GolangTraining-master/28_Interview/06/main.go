package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*

结果：
```
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```

解析：
```
1、当程序运行到defer函数时，不会执行函数实现，但会将defer函数中的参数代码进行执行。
因此首先执行的是calc("10", a, b))，随后执行的是calc("2", a, calc("20", a, b))
得到第一行和第二行结果。
2、defer的执行结果是先进后出，从函数尾部向函数头部以此执行。因此会首先执行calc("2", a, calc("20", a, b))，
然后执行defer calc("1", a, calc("10", a, b))，相应打印第三行和第四行
```

*/