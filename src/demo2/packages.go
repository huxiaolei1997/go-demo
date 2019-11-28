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

	//defer 语句会延迟函数的执行直到上层函数返回。
	//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	//defer fmt.Println("world")
	//fmt.Println("hello ")

	//延迟的函数调用被压入一个栈中。当函数返回时，
	// 会按照后进先出的顺序调用被延迟的函数调用。
	// 下面输出 10,9,8,7,6,5，4,3,2,1,0
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	pointer()

	structFunc()

	arrayFunc()

	sliceFunc()

	rangeFunc()

	mapFunc()

	Funcfunc()

	// 闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
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

func pointer() {
	/**
	Go 具有指针。 指针保存了变量的内存地址。

	类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。

	var p *int
	& 符号会生成一个指向其作用对象的指针。

	i := 42
	p = &i
	* 符号表示指针指向的底层的值。

	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 21         // 通过指针 p 设置 i
	这也就是通常所说的“间接引用”或“非直接引用”。

	与 C 不同，Go 没有指针运算。
	*/
	i, j := 42, 2701
	// 获取变量 i 的地址，并赋值给p
	p := &i

	// *p 代表取出这个地址存放的值，也就是读取到了 i 的值
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func structFunc() {
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	p := &v
	p.X = 2333
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		p1 = &Vertex{1, 2} // 类型为 *Vertex
	)

	fmt.Println(v1, p1, v2, v3)
}

func arrayFunc() {
	// 类型 [n]T 是一个有 n 个类型为 T 的值的数组。
	// 数组的长度是其类型的一部分，因此数组不能改变大小。
	// 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func sliceFunc() {
	//一个 slice 会指向一个序列的值，并且包含了长度信息。
	//[]T 是一个元素类型为 T 的 slice。
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	// 对slice切片
	//slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
	//
	//表达式
	//
	//s[lo:hi]
	//表示从 lo 到 hi-1 的 slice 元素，含两端。因此
	//
	//s[lo:lo]
	//是空的，而
	//
	//s[lo:lo+1]
	//有一个元素。
	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])

	/**
	slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：

	a := make([]int, 5)  // len(a)=5
	为了指定容量，可传递第三个参数到 `make`：

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	*/
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	// slice 的零值是 `nil`
	// 一个 nil 的slice 的长度和容量是 0
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// 向 slice 添加元素
	// append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。
	//
	//append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
	//
	//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
	var cc []int
	printSlice("cc", cc)

	cc = append(cc, 0)
	printSlice("cc", cc)

	cc = append(cc, 1)
	printSlice("cc", cc)

	cc = append(cc, 2, 3, 4)
	printSlice("cc", cc)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func rangeFunc() {
	// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 可以通过赋值给 _ 来忽略序号和值。
	// 如果只需要索引值，去掉“, value”的部分即可。
	for _, value := range pow {
		fmt.Printf("2**%d = %d\n", value, value)
	}

	pow2 := make([]int, 10)

	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	dx1 := make([][]uint8, dx)
	for i := range dx1 {
		dy1 := make([]uint8, dy)
		for j := range dy1 {
			dy1[j] = uint8(i*j - 1)
		}
		dx1[i] = dy1
	}
	return dx1
}

func mapFunc() {
	// map 映射键到值。
	//
	//map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat:  40.68333,
		Long: -74.323359,
	}

	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m["Bell Labs"])

	//在 map m 中插入或修改一个元素：
	//
	//m[key] = elem
	//获得元素：
	//
	//elem = m[key]
	//删除元素：
	//
	//delete(m, key)
	//通过双赋值检测某个键存在：
	//
	//elem, ok = m[key]
	//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	//
	//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	mm := make(map[string]int)
	mm["Answer"] = 42
	fmt.Println("The value:", mm["Answer"])

	mm["Answer"] = 48
	fmt.Println("The value:", mm["Answer"])

	delete(mm, "Answer")
	fmt.Println("The value:", mm["Answer"])

	v, ok := mm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

func Funcfunc() {
	// 函数也是值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

/**
Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
