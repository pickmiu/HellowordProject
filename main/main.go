package main

import (
	"fmt"
	"math"
)

type Chicken struct {
	color  string
	isMale bool
}

type Animal interface {
	run()
}

func (c Chicken) run() {
	fmt.Println("chicken run")
}

func main() {
	// var a int
	// var b float32
	// var c, d float64
	// e, f := 9, 10
	// var g = "Ricardo"
	// fmt.Print("hello word")

	// 返回值部分不使用 匿名变量占位
	//_, v, _ := getData()
	// fmt.Println(v)

	// showDefer()
	// showMake()

	// 结构体 多态
	// var animal Animal
	// animal = Chicken{}
	// animal.run()
}

// go 中的 if else
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim
	}
}

// 多返回值的函数
func getData() (int, int, int) {
	// 返回3个值
	return 2, 4, 8
}

func cycle() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for {
		fmt.Println("")
	}
}

func add(x, y int) int {
	return x + y
}

// 返回值命名
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func showArray() {
	var a [10]int
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func showDefer() {
	defer fmt.Println("world")
	defer fmt.Println("the")
	fmt.Println("hello")
}

func showMake() {
	// 0值数组
	a := make([]int, 5)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = make([]int, 0, 5)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
