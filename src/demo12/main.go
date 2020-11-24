package main

import . "fmt"

func main()  {
	Println("空接口使用,断言")
	Println("------方法1-------")
	fn()
	Println("------方法2-------")
	fn1()
	Println("------方法3-------")
	fn2()
}
func fn()  {
	Println("空接口使用")
    Println("可以实现其他语言的any类型")
	type  A interface {}
	var c A
	var a A=10
	Printf("\n值%v 类型%T\n",a,a)

	var str="你好"
	c=str
	Printf("\n值%v 类型%T\n",c,c)
}

func fn1()  {
	Println("任意类型")
	type A interface {}
	var m1=make(map[A]A)
    m1["name"]=12
	m1[12]=1
	m1[true]=false
	m1['a']=123

	Printf("\n值%v 类型%T\n",m1,m1)

}
func fn2()  {
	Println("类型断言使用")
	type A interface {}
   var a A
	a="你好"
    Println("eg: 变量.(类型)返回俩个值 第一个是原有的值，第二个值如果是返回true 否则返回false")
	_,ok:= a.(string)
    if ok {
    	Println("断言成功")
	}else {
		Println("断言失败")
	}
	Println(" 变量.(type)断言是必须结合switch使用")
	switch a.(type) {
	case string :
		Println("字符串类型")
	case int:
		Println("数字类型")
	case bool:
		Println("布尔类型")
	case float64:
		Println("浮点类型")
	case byte:
		Println("字符类型")
	}
}
