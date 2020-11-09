package main

import . "fmt"

func main()  {
  Println("map 学习使用")
  Println("map 是一种无序的键值对数据 ！！！是引用类型数据")
  Println("创建方法1")
  var mapA=make(map[string]int)
  mapA["age"]=18
  mapA["class"]=10
  Println("访问map  mapA[\"age\"] ")
  Printf("\n mapA[age]=%v",mapA["age"])
  Println("创建方法2 不使用make函数")
  var mapB=map[string]string{
  	"name":"张三",
  	"fastName":"张",
  	"lastName":"三",
  }
  Println("map赋值最后一个也必须加上逗号")
  Println(mapB)
  Println("方法3")
  mapC:=map[string]string{
	  "name":"李四",
	  "fastName":"李",
	  "lastName":"四",
  }
  Println(mapC)
 Println("//遍历 map for range")
	for key, val := range mapC {
     Printf("\n key=%v value=%v",key,val)
	}
	Println("创建修改方法1")
	var user=make(map[string]string)
	user["name"]="王五"
	user["age"]="20"
	Println(user)
	Println("创建修改方法2")
	var user1=map[string]string{
		"name":"李四",
		"age":"30",
	}
	user1["age"]="18"
	Println(user1)
   Println("查找是否存在")
	Println("v,ok:=user1[\"sex\"] // ok存在返回true 不存在返回false v 指定键的值,若不存在返回空")
	v,ok:=user1["age"]
	if ok{
		Println(v)
	}else {
		Println("不存在")
	}
	Println("删除map键值对 delete函数")
	var user2=map[string]string{
		"name":"李四",
		"age":"30",
		"sex":"男",
	}
	Printf("\n删除前的值: %v \n",user2)

	delete(user2,"sex")
	Printf("\n删除后的值: %v \n",user2)




}