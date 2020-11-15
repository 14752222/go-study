package main

import "fmt"

func main() {
	// var a, b string
	// a = "123"
	// b = "aaa"
	// fmt.Println(a, b);
	var (
		name string
		age  int = 10
		sex  string
	)
	name = "123"
	age = 20
	sex = "男"
	fmt.Println(name, age, sex)

	username := "zzl" //函数中使用  但是不能用于全局

}
