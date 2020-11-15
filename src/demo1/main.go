package main

import . "fmt"

func main()  {
	 Println("defer语句使用")
	 Println("derfer会将defer定义的语句延迟执行，在定义函数返回时执行，如果有多个，按照defer的定义执行的逆序执行，就是最先定义的最后执行，最后定义的最先执行")
	 Println("eg:正常执行流程---------")
	 Println("函数中return语句底层实现，return x ->1、返回值=x 2、然后执行->return 指令")
	 Println("defer函数语句底层实现")
	 Println("函数中return语句底层实现，return x ->1,返回值=x 2、运行defer 3、然后执行->return 指令")

	Println("开始")
	 Println(1)
	 Println(2)
	 Println(3)
	 Println("结束")
	 Println("eg:defer流程----------")
	 defer Println("开始")
	 defer Println(1)
	 defer Println(2)
	 defer Println(3)
	 defer Println("结束")

	 fn()
	 var num=fn1();
	 Println(num,"fn1()匿名返回值函数执行结果，匿名返回值是操作之前的值")//0
	 var num1=fn2();
	 Println(num1,"fn2()命名返回值函数执行结果，命名返回值是操作之后的值")//1

	 Println("练习-------------------start")
	 Println(f1())
	 Println(f2())
	 Println(f3())
     Println(f4())
	 Println("练习-------------------end")
	 Println("练习2-------------------start")
	//defer 注册要延迟执行的函数时该函数所有的参数都需要确定值
	 Println("!!!defer 注册要延迟执行的函数时该函数所有的参数都需要确定值")
	 Println("!!!defer 注册要延迟执行的函数时该函数所有的参数都需要确定值")
	 Println("!!!defer 注册要延迟执行的函数时该函数所有的参数都需要确定值")
     x :=1
     y:=2
     defer calc("AA",x,calc("A",x,y)) //x的值已经确定了 x=>1
     x=10
     defer calc("BB",x,calc("B",x,y)) //x的值已经被重新修改了 x =>10
	 Println("练习2-------------------start")
     y=20
/*
1注册顺序
        defer calc("AA",x,calc("A",x,y))
        defer calc("BB",x,calc("B",x,y))
2执行顺序
        defer calc("BB",x,calc("B",x,y))
        defer calc("AA",x,calc("A",x,y))
3执行步骤
   3.1 calc("A",x,y) //结果是  A 1 2 3
   3.2 calc("B",x,y) //结果是  A 10 2 12
   3.3 calc("BB",x,calc("B",x,y))等同于==>calc("BB",x,12)  //结果是 BB 10 12 22
   3.4 calc("AA",x,calc("A",x,y))等同于==>calc("AA",x,3)   //结果是 AA 1 3 4
*/

}

func calc(index string,a,b int)int  {
	ret :=a+b
	Println(index,a,b,ret)
	return ret
}


func fn()  {
	Println("开始")
	defer func() {
		Println("自执行函数执行。。。")
	}()
	Println("结束")

}
func fn1() int {
	var num int
	defer func() {
		num++
	}()
	return num //返回0 这里应该是先返回，然后在执行defer函数
}

func fn2() (a int) {
	 defer func(){
	 	a++
	 }()
	 return
}

//练习

func f1() int {
	x:=5
	defer func() {
		x++
	}()
	return x //5
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //6 执行x++ 在返回值
}

func f3() (y int) {
	x:=5
	defer func() {
		x++
	}()
	return x //5
}
func f4() (x int) {
	defer func(y int) {
		y++
	}(x)//defer 注册要延迟执行的函数时该函数所有的参数都需要确定值
	return 5 //5
}

