package main

import . "fmt"

type myInt int
func(m myInt) ptf() {
  Println("自定义类型里面的自定义方法 值=",m)
}

type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person) pt()  {
  Println("结构体方法")
  Printf("\n姓名:%v 年龄:%v 性别:%v\n",p.Name,p.Age,p.Sex)
}

func (p Person) set(name string,age int)  {
	Println("\n不会修改原有的值:")
	p.Name=name
	p.Age=age
}
func (p *Person) set1(name string,age int)  {
	Println("\n会修改原有的值:")
    p.Name=name
    p.Age=age
}

func main()  {
	Println("结构体方法使用")
	Println("golang没有类 但是有结构体 方法使用接收者函数")
	fn()
	fn1()
	fn2()
	Println("自定义类型的方法使用")
	var num myInt=10
	num.ptf()
}
func  fn()  {
	var p=&Person{
		Name:"张三",
		Age: 20,
		Sex: "男",
	}
	p2:=p;
	p2.Name="李四"
	Printf("p 值:%#v 类型:%T  指针:%p",p,p,&p)
	Println("\n结构体类型不是引用类型")
	Printf("p2 值:%#v 类型:%T  指针:%p",p2,p2,&p2)
}
func fn1()  {
	var p1=&Person{
		Name: "李四",
		Age:20,
		Sex: "男",
	}
    p1.pt()
}
func fn2()  {
	var p1=&Person{
		Name: "李四",
		Age:20,
		Sex: "男",
	}
	p1.pt()
	p1.set("王麻子",20)
	Printf("未使用指针函数p1的值 %#v 类型%T 指针%p",p1,p1,&p1)
	p1.set1("王麻子1",20)
	Printf("使用指针函数p1的值 %#v 类型%T 指针%p",p1,p1,&p1)
}
func fn3()  {

}




func test() struct{} {
	type user struct {

	}
	return user{}
}

