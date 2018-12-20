package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func addA(x, y int) int {
	return x + y + y
}

func addB(x, y int) int {
	return x + x + y
}

/**
 *	字符串颠倒，并返回多个参数
 */
func swap(x, y string) (string, string) {
	return y, x
}

/**
* 隐含返回值,多个参数
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//声明多个参数，参数类型在最后
//var c, python, java bool

var i, j int = 1, 2

var c, python, java = true, false, "noe"

//Pi eeee
const Pi = 3.143

func main() {

	fmt.Println(math.Pi)

	a := byte(255)
	b := uint8(255)
	c := int8(127)
	d := int8(a)
	e := int8(c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// fmt.Println("My favorite nu  mber kis dd", math.Sqrt(7))
	// fmt.Println(add(123, 2))
	// fmt.Println(addA(1, 3))
	// fmt.Println(addB(2, 4))
	// //变量定义
	// k := 22
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(split(17))
	// var i int
	// fmt.Println(i, c, python, java, k)
	//test1.ddd()

	// var ut uint = 112444444
	// var ut64 uint64 = 1<<64 - 1
	// var z complex128 = cmplx.Sqrt(-5 + 12i)
	// fmt.Println(ut, "#", ut64)
	// fmt.Println(z)

	//默认值信息
	// var (
	// 	i int
	// 	f float64
	// 	b bool
	// 	s = "ddd"
	// )
	// fmt.Println(i, f, b, s)

	// i := 42
	// f := i
	// u := f

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

}
