package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("//定义int 类型")
	var num int = 10 //整形
	fmt.Printf("num= %T 类型", num)

	/*错误
	num=22.2
	num="123"

	*/

	/*取范围 有符号整型
	int8 -128~127
	int16 -32768-32767
	int32 -2^31~2^31-1
	*/
	var num2 int8 = 100
	fmt.Println("unsafe.Sizeof(num2) 需要导入unsafe包")
	fmt.Printf("num2= %T 类型\n", num2)
	fmt.Println("num存储空间", unsafe.Sizeof(num))
	fmt.Println("num2存储空间", unsafe.Sizeof(num2))
	fmt.Println("int在 64位计算机表示int64")
	/*
		错误
		 num2=129 超出范围了

	*/
	fmt.Println("无符号整型")
	var num3 uint8 = 255
	//var num3 uint8 = 555
	fmt.Println(num3)
	fmt.Println("ware 两种不同类型的数字不能相加 eg: num2+num3 //错误")
	fmt.Println("如果需要相加需要使用转换方法")
	fmt.Println(uint8(num2) + num3)

	var n1 int16 = 130
	fmt.Println(int8(n1))
	fmt.Println("高位转低位 需要查看是否超出范围")
	var n2 = 10
	fmt.Printf("num=%v\n", n2) //默认原样输出
	fmt.Printf("num=%d\n", n2) //%d十进制输出
	fmt.Printf("num=%b\n", n2) //%b二进制输出
	fmt.Printf("num=%o\n", n2) //%o八进制输出
	fmt.Printf("num=%x\n", n2) //%x十六进制输出

}
