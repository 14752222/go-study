package main

import . "fmt"

type Person struct {
	Name string
	Age int
	Sex string
}

//匿名结构体，不能重复写同一个类型的简写
type  person1 struct {
	 string
	 int
}

func (p *Person) setPerson(name,sex string,age int)  {
	 p.Name=name
	 p.Age=age
	 p.Sex=sex
}
func (p Person) watchInfo()  {
	Println("watchInfo")
	Println(p,&p)
	Printf("当前的值%#v 指针地址%p\n",p,&p)
}
//结构体嵌套
type  class struct {
	ClassName int
	Student []Person
	Ftp Person //嵌套Person结构体
	Person
}

func main()  {
	 Println("结构体使用 结构体嵌套 继承")
	 Println("--------方法1---------")
    fn()
	Println("--------方法2---------")
	fn1()
	Println("--------方法3--------")
	fn2()

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
    className.Student[0].watchInfo() //继承的方法
	Printf("\np 值:%#v 类型:%T  指针:%p\n",p,p,&p)
	Printf("\nclassName 值:%#v 类型:%T  指针:%p\n",className,className,&className)

}

func fn1()  {
    Println("结构体嵌套使用，切片使用")
	var p =Person{
       Name: "李四",
       Age: 22,
       Sex: "男",
	}
	var className=class{
		ClassName: 1,
		Student: make([]Person,3),
		Ftp: Person{
			Name: "张三",
			Age:20,
			Sex: "男",
		},
	}
	Println("访问结构体className.Ftp的值",className.Ftp)//访问结构体
	className.Student[0]=p;
	Printf("\np 值:%#v 类型:%T  指针:%p\n",p,p,&p)
	Printf("\nclassName 值:%#v 类型:%T  指针:%p\n",className,className,&className)

}

func fn2()  {

	type user struct {
		Name string
		Age int
	}
	type common struct {
		Person
		user
	}
	var com=common{
         Person{
         	Name: "张三",
			Sex: "张三",
         	Age:22,
		 },
		 user{
         	Name: "李四",
         	Age: 18,
		 },
	}
	//Println(com.Name) error
	Println(com.Person.Name)
	Printf("\ncom 值:%#v 类型:%T  指针:%p\n",com,com,&com)


}
