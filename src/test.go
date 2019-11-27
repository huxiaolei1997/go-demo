package main

import "fmt"

// 声明全局变量
var x, y int

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int  = 3
	b bool = true
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 常量
const m = 2.72828

// 常量
const (
	p = 2
	// s会使用p的值
	s
	q = "3333333"
	// t会使用q的值
	t
)

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {
	fmt.Println("Hello, World!")

	// 声明一个局部变量
	var a string
	// 初始化a变量，Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
	a = "RUNNOB"
	fmt.Println(a)

	// 没有初始化就为灵值
	var b int
	b = 0
	fmt.Println(b)

	// bool 零值为 false
	var c bool = false
	// 或者可以这样写
	var d = true
	fmt.Println(c)
	fmt.Println(d)

	//
	e := true
	fmt.Println(e)

	g, h := 123, "hello"

	println(x, y, a, b, c, d, e, f, g, h)

	// for循环
	for a := 0; a < 10; a++ {
		if a == 1 {
			fmt.Println("a=", a)
			continue
		}
		fmt.Println(a)
	}

	// 在Go语言中，自增操作符不再是一个操作符，而是一个语句。
	// 因此，在Go语言中自增只有一种写法：i++
	// 如果写成前置自增++i，或者赋值后自增a=i++都将导致编译错误。

	sum := sum(1, 2)
	fmt.Printf("sum=%d\n", sum)

	fmt.Printf("%.2f\n", m)

	// 将a声明为NewInt类型
	var w NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", w)
	// 将a2声明为IntAlias类型
	var w2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", w2)
}

/**
定义一个函数
*/
func sum(a, b int) int {
	num := a + b
	return num
}
