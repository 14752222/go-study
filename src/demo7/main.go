package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("定义字符串")
	var str string = "这是一段文本"
	var str1 = "\nF:\\go\\src\\demo4"
	var str2 = `we
	qweqw
	eqwee`
	fmt.Println(str)
	fmt.Printf("str=%v str类型%T", str, str)

	fmt.Println("//输出目录")
	fmt.Println("//输出目录", str1)
	fmt.Println("输出多行字符串")
	fmt.Println("str2", str2)

	fmt.Println("str长度")
	fmt.Println(len(str))
	fmt.Println("一个汉字占3个字符串")
	fmt.Println("拼接字符串 普通方法+号连接")
	fmt.Println(str + str1)
	fmt.Println("拼接字符串 Sprintf方法")
	var test = fmt.Sprintf("%v %v", str, str2)
	fmt.Println("test", test)

	fmt.Println("--------------------")
	fmt.Println("字符串分割 需要导入strings包")
	var str3 = "123-123-234-546-675"
	var arr = strings.Split(str3, "-")
	fmt.Println(arr)
	fmt.Println("切片连接 需要导入strings包")
	var str4 = strings.Join(arr, "@")
	fmt.Println(str4)
	fmt.Println("--------包含 constains----------")
	var str5 = "sdgeryheghehe"
	var str6 = "e"
	var flag = strings.Contains(str5, str6)
	fmt.Println("str5中是否包含str6", flag)

	fmt.Println("--------前缀判断/后缀判断---------")
	var flag1 = strings.HasPrefix(str5, str6)
	var flag2 = strings.HasSuffix(str5, str6)
	fmt.Println("前缀判断是否包含str6", flag1)
	fmt.Println("后缀判断是否包含str6", flag2)
	fmt.Println("--------Index lastIndex 方法 查找不到返回-1 找到返回索引值---------")
	var num = strings.Index(str5, str6)
	var num2 = strings.LastIndex(str5, str6)
	fmt.Println("子字符串出现的从头位置", num)
	fmt.Println("子字符串出现的结尾位置", num2)

}
