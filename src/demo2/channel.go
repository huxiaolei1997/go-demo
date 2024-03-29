package demo2

import "fmt"

/**
channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

ch <- v    // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。
（“箭头”就是数据流的方向。）

和 map 与 slice 一样，channel 使用前必须创建：

ch := make(chan int)
默认情况下，在另一端准备好之前，发送和接收都会阻塞。
这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。
*/
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main19() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, x+y)

	/**
	channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

	ch := make(chan int, 100)
	向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。

	修改例子使得缓冲区被填满，然后看看会发生什么。
	*/
	cc := make(chan int, 2)
	cc <- 1
	cc <- 2
	fmt.Println(<-cc)
	fmt.Println(<-cc)
}
