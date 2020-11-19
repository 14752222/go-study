package main

import . "fmt"

type Person struct {
	Name string
	Age int
	Sex string
}

func (p *Person) setPerson(name,sex string,age int)  {
	 p.Name=name
	 p.Age=age
	 p.Sex=sex
}
func (p Person) watchInfo()  {
	Println("watchInfo")
	Println(p)
}
type  class struct {
	ClassName int
	Student []Person
	Ftp Person
	Person
}

func main()  {
	 Println("结构体使用 结构体嵌套 继承")
    fn()
}

func fn()  {
	var p=&Person{
		Name:"张三",
		Age:20,
		Sex:"男",
	}
	var p1=Person{}
	p1.Sex="男"
	p1.Name="李四"
	p1.Age=21
	var className=&class{
		ClassName: 1,
		Student: []Person{},
	}
	className.Student=make([]Person,4,4)
    className.Student[0]=p1
	className.Student[1]=*p
	className.Student[2]=p1
	className.Student[3]=*p
    className.watchInfo()
	Printf("\np 值:%#v 类型:%T  指针:%p\n",p,p,&p)
	Printf("\nclassName 值:%#v 类型:%T  指针:%p\n",className,className,&className)

}
