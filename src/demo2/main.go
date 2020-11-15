package main

import (
	. "fmt"
    "errors"
	)


func main()  {
 Println("panic和recover使用")
 Println("go暂时没有异常处理机制 可以使用panic和recover来模拟处理错误")
 Println("panic可以在如何地方抛出 但是recover只能在defer调用的函数中有效！！！")
 //没有使用recover 代码会被直接结束掉
 //fn()
 //fn1()
 //Println("结束")
 fn()
 fn2()
 Println("结束")

 Println("----异常使用start----")
 Println(fn3(10,0))
 Println(fn3(10,5))
 Println("----异常使用end----")

myFn()
Println("继续执行")
}

func fn()  {
	Println("fn...")
}
func fn1()  {
	panic("出现异常问题error...")
}
func fn2()  {
	defer func() {
		err :=recover()
		if err !=nil{
			Println(err)
		}
	}()
	panic("出现异常问题error...")
}
func fn3(x,y int) int {
	defer func() {
		err :=recover()
		if err !=nil{
			Println("error",err)
		}
	}()
	return x/y
}

func readFile(fileName string) error {
   if fileName=="main.go"{
   	 return  nil
   }else {
   	return  errors.New("读取文件失败")
   }
}
func myFn()  {
	defer func() {
		err:=recover()
		if err!=nil{
			Println("推送失败记录")
		}
	}()
	var err=readFile("xx.go")
	if err!=nil{
		panic(err)
	}
}
