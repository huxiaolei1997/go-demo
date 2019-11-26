package main

import "fmt"

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int = 3
    b bool = true
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
    fmt.Println("Hello, World!")

    // 声明一个变量并初始化
    var a string = "RUNNOOB"
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
}
