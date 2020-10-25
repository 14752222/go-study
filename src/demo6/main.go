package main

import "fmt"

func main() {
	fmt.Println("布尔类型")
	/**
	1默认值是false  int默认值 0  string默认值 "" float 默认值 0
	2不支持强制将整型转换布尔类型
	3布尔值无法参与运算 也无法与其他类型进行转化
	**/
	fmt.Printf("1默认值是false  int默认值 0  string默认值 \"\" float 默认值 0\n2不支持强制将整型转换布尔类型\n3布尔值无法参与运算 也无法与其他类型进行转化\n")
	var flag bool = true
	fmt.Printf("flag=%v flag类型%T", flag, flag)

}
