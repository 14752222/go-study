package main

import . "fmt"

type num=int
type myInt int
type myFn func(int,int)int
type person struct {
	name string
	age int
	sex string
} //定义结构体后面不需要逗号

func main()  {
  Println("golang结构体使用，学习")
  Println("golang没有class概念 但是有结构体，结构体更加灵活，使用type关键字定义")
  fn()
  fn1()
  fn2()
  fn3()
  fn4()
  fn5()
  fn6()
  fn7()
}
func fn()  {
	Println("使用type关键字定义自定义类型，类型别名")
	Println("结构体首字母大写表示公有，首字母小写表示私有，变量也是一样的")
	var a num=10
	var b int=10
	var c myInt=10
	Printf("a的值%v a的类型%T\n",a,a)
	Printf("b的值%v a的类型%T\n",b,b)
	Printf("c的值%v a的类型%T\n",c,c)
	var fnc myFn = func(x , y int)int {
		return x+y
	}
	Println("自定义方法类型使用")
	var add myFn
	add=fnc
	Println(add(10,20))
}

func fn1()  {
	Println("实例化person方法1")
	var lisi person
	lisi.name="李四"
	lisi.age=20
	lisi.sex="男"
	Printf("\n普通打印 list的值%v 类型%T 指针%v \n",lisi,lisi,&lisi)
	Printf("\n#v打印 list的值%#v 类型%T 指针%v \n",lisi,lisi,&lisi)

}

func fn2()  {
	Println("使用new 实例化person方法2")
	Println("在golang在支持对结构体指针直接使用 . 来赋值访问变量 但是底层还是(*p).name来访问，赋值实现的 ")
    var p=new(person)
    p.name="张三"
    p.age=20
    p.sex="男"
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)

}
func fn3()  {
	Println("使用&符号 实例化person方法3")
	var p=&person{}
	p.name="张三"
	p.age=20
	p.sex="男"
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)
}
func fn4()  {
	Println("使用键值对 实例化person方法4")
	var p=person{
		name:"张三",
		age:20,
		sex:"男",
	}
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)
}
func fn5()  {
	Println("使用&加键值对 实例化person方法5")
	var p=&person{
		name:"张三",
		age:20,
		sex:"男",
	}
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)
}
func fn6()  {
	Println("可以不全赋值 实例化person方法6")
	var p=&person{
		name:"张三",
	}
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)
}
func fn7()  {
	Println("不写键名初始化 实例化person方法7")
	Println("需要按顺序对应，需要加逗号")
	var p=&person{
		"张三",
		20,
		"男",
	}
	Printf("\n#v打印 p的值%#v 类型%T 指针%v \n",p,p,&p)
}