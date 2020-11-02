package main

import . "fmt"

func main() {
	var arr = [...]int{1, 3, 4, 54, 6, 678, 123, 1231, 123, 1}
	Print(arr)
	Println("\n条件判断 if else if else")
	if len(arr) < 2 {
		Println("小于2")
	} else if len(arr) < 5 {
		Println("小于5")
	} else {
		Println(len(arr))

	}
	var sceos = 80
	if sceos >= 90 {
		Println("A")
	} else if sceos >= 80 {
		Println("B")

	} else if sceos >= 70 {
		Println("c")
	} else {
		Println("d")
	}

	Println("当判断体内部只有一句话，但是花括号不能省略")
	println("for循环")
	Println("方法1")
	for i := 0; i < len(arr); i++ {
		Println(arr[i])
	}
	Println("打印99乘法表")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			Printf("%v*%v=%v \t", i, j, i*j)
		}
		Println(" ")
	}

	Println("\n无限循环")
	i := 1
	for {
		Println(i)
		if i > 5 {
			break
		}
		i++
	}
	Println("\n1-50的偶数和奇数")
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			Println(i, "偶数")
		} else {
			Println(i, "奇数")

		}
	}

	Println("\n1-10的是9的倍数")
	var sum = 0
	var count = 0
	for i := 1; i < 100; i++ {
		if i%9 == 0 {
			sum += i
			count++
			Println(i)
		}
	}
	Println("count", count)

	Println("5的阶乘")

	var sum1 = 1
	for i := 1; i <= 5; i++ {
		sum1 *= i
		Println("sum1", sum1, "i", i)

	}
	Println("sum1", sum1)

	Println("打印正方形")
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			Print("*")
		}
		Println(" ")
	}

	Println("打印三角形")

	for i := 1; i < 5; i++ {
		for j := 1; j <= i; j++ {
			Print("*")
		}
		Println("*")
	}
	Println("for range 语句使用 有点forEach的感觉但是比forEach 好用")
	for key, value := range arr {
		Printf("key=%v value=%v", key, value)
		Println(" ")
	}
	var str = "A"
	Println("golang的switch不用break也不会穿透")
	switch str {
	case "a", "A":
		Println("A", "a")
	case "b", "B":
		Println("B", "b")
	case "c", "C":
		Println("C", "c")
	default:
		Println("未知")
	}
	Println("switch的不同用法")
	var str1 = "c"

	switch str1 {
	case "a":
		Println("A", "a")
	case "b":
		Println("B", "b")
	case "c":
		Println("C", "c")
	default:
		Println("未知")
	}
}
