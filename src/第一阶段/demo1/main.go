package main

import "fmt"

func main() {
	// 注释
	/*注释*/
	fmt.Println('A', "A") //结果 65 A
	//  "" 字符串   '' uncode编码
	fmt.Println("A", "B", "C")  // A B C
	fmt.Println("hello golang") //可以换行
	fmt.Print("A", "B", "C")    //ABC
	fmt.Println("\n---------------------------------------")
	var a = "aaa" //定义必须使用
	var b int = 10
	var c int = 5
	fmt.Printf("a=%v b=%v  c=%v", a, b, c) //可以查看类型
	fmt.Printf("a=%v a的类型是%T", a, a)

}
