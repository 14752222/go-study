package main

import . "fmt"

func main()  {
 Println("空接口细节和类型断言使用细节")
	fn()
}
func fn()  {
	type phone struct {
		name string
		sex string
	}
	var user= make(map[interface{}]interface{})
	user["name"]="张三"
	user["age"]=12
	user["hobby"]=[]string{"打游戏","看电视"}
   //Println(user["hobby"][1]) //interface类型的没有索引  invalid operation: user["hobby"][1] (type interface {} does not support indexing)
    n,_:=user["hobby"].([]string)
	Println(n[0]) //正确获取方法

	Println(user["hobby"])
	var p=phone{
		name:"李四",
		sex:"男",
	}
	Println(p.name)
	user["p"]=p
	p1 ,_:=user["p"].(phone)
	Println(p1.name) //正确获取方法
	//Println(user["p"].name)user["p"].name undefined (type interface {} is interface with no methods)
	Printf("\nuser值:%#v 类型:%T  指针:%p\n",user,user,&user)


}