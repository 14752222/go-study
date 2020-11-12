package main

import . "fmt"

 //全局变量
 var num int=10
 //error  num :=10  全局变量不能使用类型推导
func fn()  {
	Println(num)
 Println("方法1")
}

func fn2()  {
	fn()
	Println("方法2")
}

func main()  {
	Println("函数使用 递归调用，闭包使用")
	Println("变量还是同其他语言一样，就近使用")
    var num =20
    Println(num)
	//fn2()
   recursion(10)
	Println("1到100 的和")
	Println(recursion1(100))
	Println("5的阶乘")
	Println(recursion2(5))
    Println("闭包函数")
    var fnum1=closure();
    Println(fnum1()) //11
	Println(fnum1())//11
	Println(fnum1())//11
	var fnum2=closure1();
	Println(fnum2(1)) //11
	Println(fnum2(2))//13
	Println(fnum2(3))//16
}

func  recursion(n int)  {
	if n>0{
		Println(n)
		n--
		recursion(n)

	}
}

func recursion1(n int) int {
	if n>1{
		return n+recursion1(n-1)
	}else {
		return 1
	}
}
func recursion2(x int) int {
	 if x>1{
	 	return x*recursion2(x-1)
	 }else {
	 	return 1
	 }
}
//闭包函数
func closure() func() int {
	var A int=10
	return	func () int{
		return  A+1
	}
}
//闭包函数
func closure1() func(y int) int {
	var A int=10
	return	func (y int) int{
		A+=y
		return A
	}
}