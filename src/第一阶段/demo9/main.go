package main

import (
	. "fmt"
	strconv "strconv"
)

func main() {
	Println("转换类型")
	Println("int 1、类型建议低位转高位，因为高位转低位可能会导致超出范围\n2、浮点类型也是一样，但是浮点类型转整型会导致数据失真，\n所以将整型转为浮点型在运算")
	Println("--整型转换--")
	var num int16 = 20
	var num1 int32 = 26
	Printf("int32(num)+num1=%v", int32(num)+num1)
	Println("\n--浮点转换--")
	var fnum float32 = 13.1
	var fnum1 float64 = 23.22
	Printf("float32(fnum)+fnum1=%v", float64(fnum)+fnum1)
	Println("--其他类型转换String类型 Sprint--")
	Println("--使用中需要注意格式 int类型 为百分号d float类型 为 百分号f bool类型")
	Print("为 百分号t byte类型 为 百分号c")
	var num2 int = 32
	var fnum2 float32 = 13.32
	var flag bool = true
	var str byte = 'a'

	str1 := Sprintf("%d", num2)
	Printf("\n值：%vstr1类型是：%T ", str1, str1)
	str2 := Sprintf("%f", fnum2)
	Printf("\n值：%v str2类型是：%T", str2, str2)

	str3 := Sprintf("%t", flag)
	Printf("\n值：%vstr3类型是：%T ", str3, str3)

	str4 := Sprintf("%c", str)
	Printf("\n值：%v str4类型是：%T", str4, str4)

	var num3 int = 30
	Println("strconv.FormatInt \n有俩个参数 \n1 是int64类型的数字 \n2是int类型的进制")
	str5 := strconv.FormatInt(int64(num3), 10)
	Printf("\n值%v 类型: %T", str5, str5)

	var fnum3 float32 = 2.32
	Println("strconv.FormatFloat \n有四个参数 \n1 是float64类型的数字 \n2格式输出的样式(f,b,e,E,g,G)，\n3保留的小数点，\n4格式的类型\n")
	str6 := strconv.FormatFloat(float64(fnum3), 'f', 2, 64)
	Printf("\n值%v 类型: %T", str6, str6)
	var flag1 = true
	str7 := strconv.FormatBool(flag1)
	Printf("\n值%v 类型: %T", str7, str7)

	var bystr = 'a'
	Println("strconv.FormatUint \n有俩个参数 \n1 是uint64类型的数字 \n2是字符类型的进制")

	str8 := strconv.FormatUint(uint64(bystr), 10)
	Printf("\n值%v 类型: %T", str8, str8)
	Println("转换数字的方法")
	Println("方法1 fmt包的Sprintf")
	Println("方法2 strconv包的Format+类型")

	Println("----------------string转浮点型--------------------------")
	Println("方法1 strconv.ParseInt \n有三个参数 \n1需要转换的字符串 \n2 转换的进制 \n3 转换的位数 int64 int32等等 \n注意返回两个值，转换失败返回0!!!!")
	var fstr = "12.23"
	var istr = "23"
	Println("方法1 strconv.ParseFloat")
	str9, _ := strconv.ParseInt(istr, 10, 64)
	Printf("\n值%v 类型%T", str9, str9)

	Println("方法2 strconv.ParseFloat \n有两个参数 \n1需要转换的字符串 \n2 转换的位数 Float64 Float32等等 \n注意返回两个值，转换失败返回0!!!!")

	str10, _ := strconv.ParseFloat(fstr, 64)

	Printf("\n值%v 类型%T", str10, str10)

	flag2, _ := strconv.ParseBool("true")
	Printf("\n值%v 类型%T", flag2, flag2)
	flag3, _ := strconv.ParseBool("false")
	Printf("\n值%v 类型%T", flag3, flag3)
	flag4, _ := strconv.ParseBool("xxxx")
	Printf("\n值%v 类型%T", flag4, flag4)
	Println("\n如果字符值不是true false 转义会失败 不建议使用 go语言中没法把数值类型转换为bool类型 bool类型也没法转换成数值类型")

}
