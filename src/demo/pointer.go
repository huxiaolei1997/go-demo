package demo

import "fmt"

/**
此函数会在包被导入时执行，例如如果是在main中导入包，包中存在init()方法，那么init()中的代码会在main()函数执行前执行，用于
初始化包所需要的特定资料。
*/
func init() {
	fmt.Println("pointer.go demo init")
}

func Main() {
	var a int = 1
	b := 2
	fmt.Printf("%p %p %d\n", &a, &b, *&a)

	var n *int
	n = &a
	fmt.Println("n = ", n)
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	var o *string
	o = &house
	fmt.Println("o = ", *o)
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	// new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)

}
