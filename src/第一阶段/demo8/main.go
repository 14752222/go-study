package main

import . "fmt"

func main() {
	Print("定义字符 字符在go语言中是int类型的")
	var a = 'a' //字符只能一个，多个就是字符串
	Println("//字符只能一个，多个就是字符串")
	Printf("值%v 类型%T", a, a)
	Println("\n汉字是占三个字节的 ，字母是一个字节")
	Println("// 字符有俩种类型 \n1、uint8 ascll类型字符 \n2、rune类型 UTF-8类型字符")
	var str = "this"
	Printf("\n值%v 类型%T", str[2], str[2])
	var str1 = "这是一段文本,hhhhhhhh"
	Print("\nstr1字节大小:")
	Println(len(str1))
	var str2 = '共'
	Printf("\n值%v 类型%T 原样%c", str2, str2, str2)
	Println("\n值20849 类型int32 因为汉字是unicode编码的")
	Println("循环字符串")

	for i := 0; i < len(str1); i++ {
		Printf("\n%v (%c)", str1[i], str1[i])
	}
	Println("\n-------------")
	for _, i := range str1 {
		Printf("\n%v=%c 类型%T", i, i, i)
	}
	Println("有汉字用uine方法，没有就可以用byte")
	var str3 = "fwef"
	var tes = []byte(str3)
	tes[0] = 'c'
	Println(string(tes))
	Println("有汉字的方法")
	var str4 = "有汉字fwef"
	var tes2 = []rune(str4)
	tes2[0] = '没'
	Println(string(tes2))

}
