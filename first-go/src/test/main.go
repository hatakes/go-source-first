package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)

	fmt.Println("hello world");

	var intVal int

	//intVal :=1 // 这时候会产生编译错误

	intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句

	fmt.Println(intVal,intVal1)


	var var1 Student

	fmt.Println(var1)
	var str5 string
	var a_1 *int
	var a12 []int
	//var mapA map[string] int
	var a2 chan int
	var a3 func(string) int
	var a4 error // error 是接口
	fmt.Println(a_1)
	fmt.Println(a12)
	//fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Printf("str5: %q \n", str5)
	c := 3+4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f", cmplx.Pow(math.E, 1i * math.Pi) + 1)
	const a1,b1  = 3,4
	var c1 int
	c1 = int (math.Sqrt(a1 * a1 + b1 * b1))
	fmt.Println(c1)

	enums();
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		java
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, golang, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

type Student struct {

}