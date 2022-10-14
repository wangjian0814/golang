package main

import (
	"fmt"
)

func main() {

	//ascii循环取值
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("ascii: %c %d\n", a[i], a[i])
	//}

	//unicode循环取值
	//for _, s := range a {
	//	fmt.Printf("Unicode: %c %d\n", s, s)
	//}
	//查找关键字
	//comma := strings.LastIndex(a, ",")
	//pos := strings.LastIndex(a[comma:], "死神")
	//fmt.Println(comma, pos, a[comma+pos:])

	var a = 2
	var b = 8
	title := fmt.Sprintf("已采集%d个草药, 还需要%d个完成任务", a, b)
	fmt.Println(title)

	pi := 3.14159
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)
}
