package main

import . "fmt"

func main() {
	var num = 10
	var num1 = 12
	Println("---运算符---")
	Printf("\n%v+%v=%v", num1, num, num1+num)
	Printf("\n%v-%v=%v", num1, num, num1-num)
	Printf("\n%v*%v=%v", num1, num, num1*num)
	Printf("\n%v/%v=%v", num1, num, num1/num)
	Printf("\n%v%v%v=%v", num1, "%", num, num1%num)
	Println("余数计算方法 eg 10%3 计算方法：10-（10/3）*3 ")
	Println("\n1除法运算时 如果是除后会去掉小数，只保留整数")
	var fnum = 10.0
	var fnum1 = 5.0
	Printf("%v/%v=%v", fnum, fnum1, fnum/fnum1)
	Println("自增自减是单独语句，不能赋值,只能单独使用，没有前置++ -- ")
	num++
	Printf("%v++", num)
	num--
	Printf("%v--", num)

	Println("---关系运算符---")

	Printf("\n%v==%v=%v", num1, num, num1 == num)
	Printf("\n%v!=%v=%v", num1, num, num1 != num)
	Printf("\n%v>=%v=%v", num1, num, num1 >= num)
	Printf("\n%v>%v=%v", num1, num, num1 > num)
	Printf("\n%v<%v=%v", num1, num, num1 <= num)
	Printf("\n%v<=%v=%v", num1, num, num1 < num)

	flag := num > num1
	if flag {
		Println("\nnum>num1")
	} else {
		Println("\nnum<num1")

	}

	Println("逻辑运算")
	Println(num > 1 && num1 > 2)   //true
	Println(num < 21 && num1 < 20) //true
	Println(num > 10 || num1 > 10) //true
	Println(num < 10 || num1 < 10) //false
	Println(num < 10 || num1 < 10) //false
	Println(!false)                //ture
	Println("\n也可以执行短路操作")
	if flag && test() {
		Println("&&开始")
	}
	if flag || test() {
		Println("||开始")
	}

	Println("赋值运算")
	var num2 = 10
	var num3 = 10
	var num4 = num2 + num3
	Printf("\nnum2 + num3=%v \n", num4)
	num2 += 10
	Println(num2)
	Printf("\nnum2 +=%v", num2)
	num2 -= 4
	Printf("\nnum2 -=%v", num2)
	num2 /= 4
	Printf("\nnum2 /=%v", num2)
	num2 *= 4
	Printf("\nnum2 *=%v", num2)
	num2 %= 4
	Printf("\n num2 %v= %v", "%", num2)

}
func test() bool {
	Println("执行过")
	return true
}
