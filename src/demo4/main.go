package main

import . "fmt"

func main()  {
	Println("指针使用")
	Println("指针也是一个变量，但是它是一个特殊的变量，它储存的数据不是普通值，而是另外一个变量的内存地址")
	Println("\n------fn-----------\n")
	fn()
	Println("\n------fn1----------\n")
	fn1()
	Println("\n------fn2----------\n")
	fn2()
	Println("\n------fn3----------\n")
    fn3()
	Println("\n------fn4----------\n")
	fn4()
	Println("\n------new()和make()的区别----------\n")
    Println("共同点 二者都是用来做内存分配的")
	Println("不同点 make只能用于slice map channel的初始化，返回的还是这三个引用类型本身")
	Println("不同点 new用于类型的内存分配，并且内存对应的值类型为对应类型的默认值，返回的是指向类型的指针")


}
func fn()  {
	Println("使用指针 &变量名")
	var a int=10
	var b=&a
	Printf("\na=->值:%v 类型:%T 地址:%p \n",a,a,&a)
	Printf("\nb=->值:%v 类型:%T 地址:%p \n",b,b,&b)
}
func fn1()  {
	var a=10
	var b=&a
	Println("原有的a的值",a)
	*b=30 //修改地址中的值
	Println("修改*b=30 a的值",a)
    Printf("值%v 类型%T 地址%p",b,b,&b)
}
func t(x int)  {
	x=30
}
func t1(x *int)  {
	*x=100
}
func fn2()  {
	var a=10
	Printf("\n默认情况下a的值:%v 类型:%T 地址:%p \n",a,a,&a)
	t(a)
	Printf("\n调用t()方法情况后a的值:%v 类型:%T 地址:%p \n",a,a,&a)
	t1(&a)
	Printf("\n调用t1()方法情况后a的值:%v 类型:%T 地址:%p \n",a,a,&a)
}
func fn3()  {
	Println("指针也是应用类型")
	var userInfo=make(map[int]int)
	Println("userInfo地址",&userInfo)
	userInfo[1]=10
	Println(userInfo)
	var age=make([]int,2,2)
    Println("age地址",&age)
	age[0]=10
    Println(age)
}
func fn4()  {
	Println("new函数使用后得到一个指针类型，并且指针默认值是对应类型的默认值")
	A:=new(int)
	var a *int
	b:=new(*int)
	Println("a默认值",A)
	Println("a默认值",*A)
	Println("a默认值",a)
	Println("b默认值",b)
	Println("b默认值",*b)
	//基本上不会用
	a=new(int)
	*a=10
	Println("a修改后",a)
	Println("a修改后",*a)

	/*
	错误写法
	var a *int
	 *a=100
	*/
}
