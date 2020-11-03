package main

import . "fmt"

func main() {
	flag := true
	Println("goto语句使用方法：goto <名称(随意)> 然后是需要跳转的位置是跳转到之前声明标签名字，会跳到指定位置执行")
	if flag {
		goto goto_label
		Println("123")
	}
	Println("123")
goto_label:
	Println("to")
	Println(&flag)

	var num = "123"
	var num1 = &num
	Printf("\n %v %v", num, num1)

	Println("lable使用")
lable:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue lable
			}
			Printf("i=%v j=%v\n", i, j)
			Println(" ")
		}
	}

}
