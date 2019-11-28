package demo2

/*
color.Color 和 color.Model 也是接口，
但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。
这些接口和类型由image/color 包定义。
*/
import (
	"fmt"
	"image"
)

func main17() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
