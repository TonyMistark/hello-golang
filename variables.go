package main

import "fmt"

var x, y int

var (
	a int
	b bool

)


var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现

func main() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

}
