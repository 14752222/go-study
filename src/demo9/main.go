package main  //package 必须在最上面

import (
	. "fmt"
	"test/calc" //"test/calc"  斜杠后面是文件夹的名字 不是go文件
	"test/tools"
	)
//

func init()  {
	Println("---init()函数是系统调用 优先执行init()函数---")
	Println("golang执行顺序 全局声明->init()函数->main()")
	Println("假如 入口导入A包 A包导入了B包 B包导入了C包 init()函数的执行顺序为")
	Println("main.in()->A.init()->B.init()->c.init()")
	Println("---init()函数---")
}

func main()  {

	Println("golang包的使用")
	Println("golang包的类型分为，内置包，自定义包，第三方包")
	Println("golang包的项目的命名 go mod init 包名")
	Println("//首字母小写不能被导出,只能导出首字母大写的方法和变量等")
	Println("一个文件夹下面直接包含的文件只能属于一个package 同样一个package的文件不能在多个文件夹下")
	Println("包名可以不和文件夹同一个名称 包名不能包含-符号")
	Println("同一个包里面不能有两个同名的方法")
	Println("包名为main的包为应用的入口程序应用的入口包 这种包编译后会得到一个可执行文件 而编译不包含main包的源代码不会得到可执行文件")
	Println(`//import (
	. "fmt" //去掉包名写法
	 "test/calc" //"test/calc"  斜杠后面是文件夹的名字 不是go文件
    T "test/tools" //重命名写法
    _ "test/tools" //匿名编译，不会参与编译 
)`)
    Println("----------方法1-----------")
	add()
	Println("----------方法2-----------")
	fn()
}
func add()  {
	Println("使用如果没有使用点导入，就需要加上包名点上需要使用的方法或者变量")
	Println(calc.Add(1,2))
	Println(calc.Sub(12,2))
	Println("使用导入User结构体 &calc.User")
	var u =&calc.User{
		Name:"张三",
		Age:12,
		Sex:"男",
		Id:1,
		CreateTime:"2006-01-02 03:04:05",
	}
	Printf("\nu 值:%#v 类型:%T  指针:%p\n",u,u,&u)


}

func fn()  {
  Println("10和2的最大值：",tools.Max(10,2))
  tools.Log_int(2)
}
