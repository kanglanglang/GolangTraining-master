package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	s = append(s, 4,5,6)
	fmt.Println(s)
	fmt.Println(len(s))

}

/*

结果：
```
[0 0 0 0 0 1 2 3]
8
[0 0 0 0 0 1 2 3 4 5 6]
11
```
解析：
```
make([]int,5)的含义是创建数组，并且数组初始化5个元素，5个元素的值为类型零值。
append
*/