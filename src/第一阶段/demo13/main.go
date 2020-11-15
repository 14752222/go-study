package main

import . "fmt"

func main() {
	Println("数组使用")
	Println("方法1")
	Println("var arr [2]int")
	var arr [2]int
	arr[0] = 22
	arr[1] = 11
	Printf("arr类型=%T arr=%v", arr, arr)
	Println("说明golang中数组长度也是类型的一部分，但是在golang中数组值的类型是值类型，而不是引用类型")
	Println("数组的长度被定义就不能添加超过长度的数据了")
	Println("方法2")
	Println("var arr1 = [3]int{1, 23, 4}")
	var arr1 = [3]int{1, 23, 4}
	Println(arr1)
	Println("方法3")
	Println("arr2 := [2]int{1, 2}")
	arr2 := [2]int{1, 2}
	Println(arr2)
	Println("方法4")
	Println("arr3 := [...]int{1, 2}")
	arr3 := [...]int{1, 2}
	Println(arr3)

}
