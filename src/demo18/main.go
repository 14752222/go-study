package main

import (
	. "fmt"
	"sort"
	"strings"
	)
func main(){
	Println("map 使用2")
	Println("切片中存放map")
	Println("未初始化的map 值是nil")
	var userTable=make([]map[string]string,3,3)
	Println(userTable)
	Println("赋值方法")
	userTable[0]= make(map[string]string)
	userTable[0]["name"]="张三"
	userTable[0]["age"]="18"
	userTable[0]["sex"]="男"
	userTable[1]= make(map[string]string)
	userTable[1]["name"]="李四"
	userTable[1]["age"]="20"
	userTable[1]["sex"]="男"
	userTable[2]= make(map[string]string)
	userTable[2]["name"]="王五"
	userTable[2]["age"]="28"
	userTable[2]["sex"]="男"
	Println(userTable[0])
	Println(userTable[1])
	Println(userTable[2])
	Println("遍历切片map")
	for _, val := range userTable {
		for k, v := range val {
          Printf("键:%v 值: %v",k,v)
		}
		Println(" ")
	}
	Println("map值为切片")
	var userInfo=make(map[int][]int)
    userInfo[0]=[]int{10,20}
	userInfo[1]=[]int{15,25}
    Println(userInfo)
	for _, val := range userInfo {
		for k, v := range val {
			Printf("键:%v  值: %v ",k,v)
		}
		Println(" ")
	}
	var mapA=make(map[int]int)
	mapA[0]=10
	mapA[1]=20
	mapA[2]=30
	mapA[3]=40
	var mapB=mapA;
	Println("未修改前")
	Println(mapA)
	Println(mapB)
	mapB[0]=55
	Println("修改后")
	Println("mapA",mapA)
	Println("mapB",mapB)
    Println("排序map类型")
	var mapC=make(map[int]int)
    mapC[1]=10
	mapC[2]=123
	mapC[23]=23
	mapC[27]=34
	mapC[13]=234
	mapC[12]=134
   Println(mapC)
	for k, v := range mapC {
		Println(k,v)
	/**
	打印
	13 234
	12 134
	1 10
	2 123
	23 23
	27 34**/
	}
	Println("按照key 升序输出map")
	var keySlice []int
	//获取mapC的key
	for k, _ := range mapC {
		 //新增keySlice值
         keySlice=append(keySlice,k)
	}
	//排序keySlice
	sort.Ints(keySlice)
	Println(keySlice)

	for _, v := range keySlice {
		Println(mapC[v])
	}
   Println("统计一个单词出现的次数")
   var str="is this the is isNaN"
   var strSlice= strings.Split(str," ")
   Println(strSlice)
   var mapD=make(map[string]int)
	for _, v := range strSlice {
         mapD[v]+=1
	}
	Println(mapD)

}
