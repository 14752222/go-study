package main

import "fmt"

func main() {
	// 常量必须赋值 同js 类似 定义赋值不能修改
	fmt.Printf("// 常量必须赋值 同js 类似 定义赋值不能修改")

	const num = 3.14
	/***
	声明多个
	const (
		a="a"
		b="b"
	)
	**/
	//const num 错误
	//num=10 错误
	fmt.Println("开始")

	// iota  go语言计算器
	fmt.Printf("// iota  go语言计算器")

	const _a = iota
	fmt.Print(_a)
	fmt.Printf("------_a--------\n")
	const (
		_b = iota
		_c = 200  //插队
		_d = iota //表示接上之前iota ,如果没有iota 后面的值都等于200
		_e
		_f
	)
	fmt.Println(_b, _c, _d, _e, _f)

	//多个iota一起定义
	fmt.Println("//多个iota一起定义")
	const (
		n1, n2 = iota, iota + 1
		n3, n4
		n5, n6
	)
	fmt.Println(n1, n2)
	fmt.Println(n3, n4)
	fmt.Println(n5, n6)
	fmt.Println("//区分大小写")
	var sex = 10
	var SEX = 20
	fmt.Println(sex, SEX)               //区分大小写
	fmt.Println("go语言 大驼峰表示共有 小驼峰表示私有") //区分大小写

}
