package demo2

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 全局变量
var i, j int = 1, 2

/**
常量
常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。
*/
const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func Test() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("My favorite number is %g\n", math.Nextafter(2, 10))
	//在 Go 中，首字母大写的名称是被导出的。
	//Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(add2(3, 2, 3))

	a, b := swap("aaa", "bbb")
	fmt.Println(a, b)

	//短声明变量
	//在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	//函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	c, d := split(36)
	fmt.Println(c, d)

	var i, j int
	fmt.Println(i, j)

	var javascript, python, java = true, false, "no!"
	fmt.Println(javascript, python, java)

	//bool
	//
	//string
	//
	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//
	//byte // uint8 的别名
	//rune // int32 的别名
	//     // 代表一个Unicode码
	//float32 float64
	//complex64 complex128
	fmt.Printf("\n")
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	fmt.Printf("\n")
	conversions()

	fmt.Printf("\n")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
	fmt.Printf("\n")
	forloop()
	fmt.Printf("\n")
	fmt.Println(sqrt(2))
	fmt.Printf("\n")
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
	fmt.Printf("\n")
	switchFunc()
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y, z int) int {
	return x + y + z
}

// 返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

//命名返回值
//Go 的返回值可以被命名，并且像变量那样使用。
//返回值的名称应当具有一定的意义，可以作为文档使用。
//没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
//直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 此时x, y的值将会被返回
	return
}

func conversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// for
//Go 只有一种循环结构——`for` 循环。
//
//基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。

func forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 跟 C 或者 Java 中一样，可以让前置、后置语句为空。
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// 死循环
	//for {
	//}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func switchFunc() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
