package main

import . "fmt"

type User interface {
	start()
	stop()
}
type Phone struct {
	Name string
}

func (p Phone) start()  {
Println(p.Name,"启动")
}
func (p Phone) stop()  {
Println(p.Name,"关闭")
}
func (p Phone) run()  {
	Println(p.Name,"关闭")
}
func main()  {
	 Println("golang中接口的使用")
	 Println("golang中接口interface是一种类型 一种抽象的类型 ，接口是一组函数的集合 golang中的接口不能包含如何变量")
	 Println("如果接口里面有方法，就必须通过结构体或者自定义类型来实现这个接口")
	 Println("接口就是一种数据类型")
	 Println("接口中定义的都必须实现方法，变量")
     Println("------方法1-------")
	 fn()
	Println("------方法2-------")
	fn1()

}

func fn()  {
	var p=Phone{
		"小米",
	}
	Printf("\np 值:%#v 类型:%T  指针:%p\n",p,p,&p)
	p.start()

	p.stop()
}
func fn1()  {
	var p=Phone{
		"小米",
	}
  var	p1 User
   p1=p
   p1.start()
   p1.stop()
   //p1.run()//错误调用 p1.run undefined (type User has no field or method run)
   p.run()  //正确调用

}