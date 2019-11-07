

package main

import (
"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}


/*

结果：
```
编译不过去
```
解析：
```
原因是Student并没有继承People，people中有Speak(string)string方法，而Student类型中没有Speak()方法，
代码中的Student方法是*Student类型的，所有  var peo People = Student{}是不符合规范的。
```

*/