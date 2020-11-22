package main

import (
	. "fmt"
	"encoding/json"
	)

type User struct {
	Name string
	Age int
	Sex string
	son  string  //私有属性没办法转换json数据格式
}
//结构体标签使用
type User1 struct {
	Name string `json:"name"`
	Age int     `json:"age"`
	Sex string  `json:"sex"`
	Son  string  `json:"son"`
}
func main()  {
	Println("结构体互相转换json方法使用，序列化，反序列化")
	Println("--------方法1---------")
	fn()
	Println("--------方法2---------")
	fn1()
	Println("--------方法3---------")
	fn2()
	Println("--------方法3---------")
	fn3()

}

func fn(){
	Println("结构体转化为json数据")
	var u =User{
       Name: "李四",
       Age: 21,
       Sex: "男",
       son:"隶属",
   }
  Printf("\nu 值:%#v 类型:%T  指针:%p\n",u,u,&u)
  jsonByte,_:= json.Marshal(u)
  Println("json.Marshal(u) 调用时返回俩个值 第一个是Byte类型的切片 第二个是错误信息")
  str :=string(jsonByte)
  Printf("\nstr 值:%v 类型:%T  指针:%p\n",str,str,&str)

}
func fn1 (){
	Println("定义json字符串的方法1\"{\\\"Name\\\":\\\"李四\\\",\\\"Age\\\":21,\\\"Sex\\\":\\\"男\\\"}\" ")
	Println("定义json字符串的方法2 `{\"Name\":\"李四\",\"Age\":21,\"Sex\":\"男\"}` ")
	var str="{\"Name\":\"李四\",\"Age\":21,\"Sex\":\"男\"}"
	var str1=`{"Name":"李四","Age":21,"Sex":"男"}`
	Println("json数据转化为结构体")
    var u=User{
    	Name: "李四",
    	Age: 25,
    	Sex: "男",
    	son: "123",
	}
	Println("json.Unmarshal 第二个值指针地址")
   err:=json.Unmarshal([]byte(str),&u)
   if err!=nil{
   	 Println(err)
   }

	Printf("\nstr 值:%v 类型:%T  指针:%p\n",str,str,&str)
	Printf("\nstr1 值:%v 类型:%T  指针:%p\n",str1,str1,&str1)
	Printf("\nu 值:%#v 类型:%T  指针:%p\n",u,u,&u)


}
func fn2(){
	Println("结构体tag标签使用")
	var u =User1{
		Name: "李四",
		Age: 21,
		Sex: "男",
		Son:"隶属",
	}
	Printf("\nu 值:%#v 类型:%T  指针:%p\n",u,u,&u)
	jsonByte,_:= json.Marshal(u)
	Println("json.Marshal(u) 调用时返回俩个值 第一个是Byte类型的切片 第二个是错误信息")
	str :=string(jsonByte)
	Printf("\nstr 值:%v 类型:%T  指针:%p\n",str,str,&str)
}

func fn3()  {
	type  student struct {
		Id int
		Gender string
		Name string
	}
	type class struct {
		Title string
		Students []student
	}
	var c=class{
		Title: "1班",
		Students: make([]student,0,20),
	}

   for i:=0; i<=10; i++{
   	   str :="男"
   	   if i/2==0{
   	   	 str="女"
	   }
	    u:=student{
		   Id: i,
		   Gender: str,
		   Name: Sprintf("stu_%v",i),
	   }
	   c.Students= append(c.Students, u)
   }

	Printf("\nc 值:%#v 类型:%T  指针:%p\n",c,c,&c)

   strByte,err :=json.Marshal(c)
   if err!=nil{
   	 Println(err)
   }
   jsonStr:=string(strByte)
   Printf("\nc 值:%v 类型:%T  指针:%p\n",jsonStr,jsonStr,&jsonStr)


}