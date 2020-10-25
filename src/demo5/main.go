package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//支持 float32 float64
	// 64位计算机浮点类型默认是float64
	var n1 float32 = 12.3
	var n2 float64 = 13.3
	var n3 float64 = 3.1415926
	var n4 = 3.14e2  //表示3.14*10的二次方
	var n5 = 3.14e-2 //表示3.14/10的二次方
	var n6 float64 = 8.2
	var n7 float64 = 3.8
	fmt.Println("-------float32------")
	fmt.Printf("n1=%v 类型%T", n1, n1)
	fmt.Println("\n储存空间大小", unsafe.Sizeof(n1))
	fmt.Println("-------float64------")
	fmt.Printf("n2=%v 类型%T", n2, n2)
	fmt.Println("\n储存空间大小", unsafe.Sizeof(n2))
	fmt.Println("-------%/v %/f-------")
	fmt.Printf("%v\n", n3)
	fmt.Printf("%f", n3)
	fmt.Println("-------%/v %/.2f-------")
	fmt.Println("%.2-f 点几f 就是保留几位小数输出")
	fmt.Printf("原样 %v 保留2位小数n3=%.2f", n3, n3)
	fmt.Println("--------- //3.14e2表示3.14*10的二次方------------")
	fmt.Printf("n4=%v 类型n4=%T \n", n4, n4)
	fmt.Println("--------- //3.14e-2表示3.14/10的二次方------------")
	fmt.Printf("n5=%v 类型n5=%T \n", n5, n5)

	fmt.Println("精度丢失8.2-3.8")
	fmt.Println(n6 - n7)
	fmt.Println("使用第三方包解决 暂时！！！")
	fmt.Println("int 转 float64")
	var num = 10
	var num2 = float64(num)
	fmt.Printf("num=%v num类型%T num2=%v num2类型%T", num, num, num2, num2)
	fmt.Println("float 转换 int 会造成精度丢失")

}
