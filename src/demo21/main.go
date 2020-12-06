package main

import (
	. "fmt"
	"reflect"
	)

type Student struct {
	Name string `json:"name"`
	Age int  `json:"age"`
	Score int`json:"score"`
}

func (s Student) GetInfo() (str string)  {
	  str=Sprintf("姓名%v 年龄%v 成绩%v",s.Name,s.Age,s.Score)
	  return
}
func (s *Student) SetInfo(name string,age,score int)  {
	s.Name=name
	s.Age=age
	s.Score=score
}
func (s Student) Print()  {
	Println("打印")
}
func main()  {
	Println("反射使用2 结构体反射")
	Println("反射的代码非常脆弱的")
	Println("---------方法1-----------")
	fn()
}

func fn()  {
	var stu=Student{
		"张三",
		10,
		100,
	}
	//printStudent(stu)
	//printStudentFn(stu)//panic: reflect: call of reflect.Value.Call on zero Value  因为传入的是值类型而不是指针类型
	printStudentFn(&stu)
	printStudentChang(&stu)
	Printf("%#v",stu)
}
func printStudent(s interface{})  {
	t:=reflect.TypeOf(s)//类型变量
	v:=reflect.ValueOf(s) //值类型

	if t.Kind()!=reflect.Struct&&t.Elem().Kind()!=reflect.Struct{
		Println("判断传入参数是不是一个结构体")
		return
	}
	//1 通过类型变量里面的Field可以获取结构体字段
	Println("//1 通过类型变量里面的Field可以获取结构体字段")
	f:=t.Field(0)
	Printf("%#v\n",f)//reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0xf93080), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	Println(f)//{Name  string  0 [0] false}
	Println("字段名称",f.Name)
	Println("字段类型",f.Type)
	Println("字段标签",f.Tag.Get("json"))
	//2 通过类型里面的FieldByName可以获取结构体的字段
	Println("//2 通过类型里面的FieldByName可以获取结构体的字段")
	f1,ok:=t.FieldByName("Age")
	if ok {
		Println("字段名称",f1.Name)
		Println("字段类型",f1.Type)
		Println("字段标签",f1.Tag.Get("json"))
	}
	//3 通过类型里面的NumField获取到结构体有几个字段
	Println("//3 通过类型里面的NumField获取到结构体有几个字段")
	f2:=t.NumField()
	Println("结构体属性个数：",f2)
	Println("//4 通过值变量获取结构体属性对应的值")
    f3:=v.FieldByName("Name")
	f4:=v.FieldByName("Age")
	Println("结构体Name属性的值:",f3)
	Println("结构体Age属性的值:",f4)
	Println("for循环获取")
	for i := 0; i <f2; i++ {
        Printf("属性名称：%v",t.Field(i).Name)
		Printf("属性值：%v",v.Field(i))
		Printf("属性类型：%v",t.Field(i).Type)
		Printf("属性tag：%v \n",t.Field(i).Tag.Get("json"))

	}

}

func printStudentFn(s interface{})  {
	Println("要使用Method属性需要函数首字母是大写的")
	t:=reflect.TypeOf(s)//类型变量
	v:=reflect.ValueOf(s)//值变量
	if t.Kind()!=reflect.Struct&&t.Elem().Kind()!=reflect.Struct{
		Println("判断传入参数是不是一个结构体")
		return
	}
	Println("1 通过类型变量里面的Method可以获取结构体的方法")
	m:=t.Method(0)
	Println("t.Method(0)这个0和结构体方法的顺序没有关系，是按照结构体方法的名称的ascII有关系")
	Println("方法名",m.Name)
	Println("类型",m.Type)
	Println("2 通过类型变量获取结构体里面有多少个方法")
	m1,ok:=t.MethodByName("Print")
	if ok{
		Println(m1.Name)
		Println(m1.Type)
	}
	Println("3 通过<值变量>执行方法（需要注意使用值变量，并且要注意参数）v.Method(0).Call(参数没有传入nil)")
	Println("v.Method(1).Call(nil)调用方法")
	v.Method(1).Call(nil)
	m2:=v.MethodByName("GetInfo").Call(nil)
	Println(m2)
	Println("传入方法传入参数，需要使用\"值变量\"，并且要注意参数，接受参数是[]reaflect.Value的切片")
	var params []reflect.Value
	params=append(params,reflect.ValueOf("李四"))
	params=append(params,reflect.ValueOf(20))
	params=append(params,reflect.ValueOf(78))
    //panic: reflect: call of reflect.Value.Call on zero Value  因为传入的是值类型而不是指针类型
	v.MethodByName("SetInfo").Call(params)//执行方法传入参数
	m3:=v.MethodByName("GetInfo").Call(nil)
	Println(m3)
	Println("4 获取方法的数量 值变量，类型变量")
	Println("获取方法数量",v.NumMethod())
	Println("获取方法数量",v.NumMethod())


}
func printStudentChang(s interface{})  {
	Println("修改结构体属性的值")
	t:=reflect.TypeOf(s)//类型变量
	v:=reflect.ValueOf(s)//值变量
	if t.Kind()!=reflect.Ptr{
		Println("传入参数是不是一个结构体指针类型")
		return
	}else if t.Elem().Kind()!=reflect.Struct{
			Println("传入参数不是一个结构体指针类型")
			return
	}
    name:=v.Elem().FieldByName("Name")
    name.SetString("李四")
	age:=v.Elem().FieldByName("Age")
	age.SetInt(23)


}