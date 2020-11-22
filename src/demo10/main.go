package main

import (
	. "fmt"
	"github.com/shopspring/decimal"
	)

func main()  {
	 Println("golang包的使用二")
	Println("golang包的使用第三方包")
	Println("步骤1 下载第三方包")
	Println("步骤1 下载第三方包的方法1 eg: go get github.com/shopspring/decimal 全局")
	Println("步骤1 下载第三方包的方法2 eg: go mod download 全局") //下载项目的依赖
	Println("步骤1 下载第三方包的方法3 eg: go mod vendor 复制到局部")

	Println("步骤2 导入第三方包")
	Println("步骤3 使用第三方包")
	Println("-----方法1-----")
	 fn()
	Println("-----方法2-----")
   fn1()

}
func fn()  {
	n, err := decimal.NewFromString("-123.4567")
	if err!=nil{
	  Println(err)
		return
	}
	Printf("\nn 值:%v 类型:%T  指针:%p\n",n,n,&n)
	var fnum float64 = 3.1
	var fnum1 float64 = 4.2
	Println(fnum+fnum1)

}
func fn1()  {

	var fnum float64 = 3.1
	var fnum1 float64 = 4.2
	d1:=decimal.NewFromFloat(fnum).Add(decimal.NewFromFloat(fnum1))
	d2:=decimal.NewFromFloat(fnum).Sub(decimal.NewFromFloat(fnum1))
	Println(d1)
	Println(d2)

}
