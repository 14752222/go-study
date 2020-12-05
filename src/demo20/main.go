package main

import (
  .	"fmt"
  "reflect"
	)
func ptf() {
	Println("反射使用案例1")
	Println("GO语言中的变量分为两个不部分")
	Println("1,类型信息，预先定义的元信息 ")
	Println("2,值信息：程序运行过程中可以动态变化的")
	Println("程序在运行期间对本身修改和访问的能力")
	Println("反射可以用来查看元素类型")
}
func main() {
	//概述
	Println("简介")
	ptf()
	Println("----------方法1----------")
	fn()
	Println("----------方法2----------")
    fn1()
	Println("----------方法3----------")
    fn2()
	Println("----------方法4----------")
	fn3()

}
func fn() {
	Println("反射获取元素类型")
	type T interface{}
	type myint int
	var a T = 10
    v:=reflect.TypeOf(a)
	Println("a的类型",v)
    Println(reflect.TypeOf(2.34))//获取类型的元素
    Println(reflect.TypeOf(2!=1))//bool类型
    var b myint =10
    Println("自定义类型")
    Println("b的类型",reflect.TypeOf(b))//main.myint
	Println("b的类型",reflect.TypeOf(&b))// *main.myint

}
func fn1()  {
	Println("Name和Kind的使用")
	Println("Name类型的名称")
	Println("Kind就是指底层的类型")
	type myInt int
	type user struct {
		name string
		age int
	}
    var  a myInt=10
	var  b =true
	var  c =20.1
	var  u=&user{
		"张三",
		 10,
	}
	var  u1=user{
		"张三",
		10,
	}
	Println("a的类型名称",reflect.TypeOf(a).Name())
	Println("b的类型名称",reflect.TypeOf(b).Name())
	Println("c的类型名称",reflect.TypeOf(c).Name())
	Println("a的底层类型",reflect.TypeOf(a).Kind())
	Println("b的底层类型",reflect.TypeOf(b).Kind())
	Println("c的底层类型",reflect.TypeOf(c).Kind())
	Println("u的类型名称",reflect.TypeOf(u).Name())
	Println("u的底层类型",reflect.TypeOf(u).Kind())
	Println("u1的类型名称",reflect.TypeOf(u1).Name())
	Println("u1的底层类型",reflect.TypeOf(u1).Kind())
    Println("指针没有类型名称")
	var arr =[3]int{1,2,3}
	var slice =[]int{1,2,3}
	Println("arr的类型名称",reflect.TypeOf(arr).Name())
	Println("arr的底层类型",reflect.TypeOf(arr).Kind())
	Println("slice的类型名称",reflect.TypeOf(slice).Name())
	Println("slice的底层类型",reflect.TypeOf(arr).Kind())

}

func fn2()  {
	Println("reflect.ValueOf使用")
    var  a=10
    valueof(a)
	var b float32=3.14
	var c int64=641
	var d string ="你好golang"
	value1(b)
	value1(c)
	value1(d)


}
func valueof(x interface{})  {
	   //var num =10+x  //operation: 10 + x (mismatched types int and interface {})
	   //类型断言
	   //b:=x.(int)
	   //var  num =10+b
	   //反射实现
	    b:=reflect.ValueOf(x)
	    Println("b:=reflect.ValueOf(x) 获取原始值 b.Int() 其他类型同理")
  	   var num=10+b.Int()
	   Println("num+x=",num)
}

func value1(x interface{})  {
	 v:=reflect.ValueOf(x)
	 Println(v)
	switch v.Kind() {
	case reflect.Int64:
		Println("int类型",v.Int())
	case reflect.Float32:
		Println("Float32类型",v.Float())
	case reflect.Float64:
		Println("Float64类型",v.Float())
	case reflect.String:
		Println("String类型",v.String())
	default:
		Println("类型暂时未添加")

	}
}
func fn3()  {
	Println("通过反射设置元素的值")
	var a=10
	//setValue(&a) invalid indirect of x (type interface {})
	//Println(a) invalid indirect of x (type interface {})
	setValue(&a)
	var b float32=3.14
	var c int64=641
	var d string ="你好golang"
	setValue(&b)
	setValue(&c)
	setValue(&d)
	Println(a)
	Println(b)
	Println(c)
	Println(d)

}

func setValue(x interface{})  {
     //*x=200 invalid indirect of x (type interface {})
      Println("反射修改一个元素的值 Kind()指针只能只能看到ptr 看不到类型 所以要加上Elem()后Kind才能查看到原来的类型")
	v:=reflect.ValueOf(x)
	Println(v.Elem().Kind())
	switch v.Elem().Kind() {
	case reflect.Int64:
		Println("Int64类型",)
		v.Elem().SetInt(200)
		Println("Float32类型",)
	case reflect.Float32:
		v.Elem().SetFloat(3.12)
	case reflect.Float64:
		Println("Float64类型",)
		v.Elem().SetFloat(30.12)
	case reflect.String:
		Println("String类型",)
		v.Elem().SetString("这是一段文本")
	default:
		Println("类型暂时未添加")
	}
}