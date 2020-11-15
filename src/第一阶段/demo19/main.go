package main

import . "fmt"

type T func(int,int) int //定义一个T类型

//方法1 方法有返回值
func sum(x,y int)int  {
  return x+y
}
//方法2 有返回值的方法使用
func sub(x int,y int) int {
	return x-y
}
func sub1(x float64,y float64) float64 {
	return  x-y
}
// 切片排序
func sortSlcie(x []int) []int {
	for i := 0; i <len(x) ; i++ {
		for j := i+1; j < len(x); j++ {
			if(x[i]>x[j]){
				temp:=x[i]
				x[i]=x[j]
				x[j]=temp
			}
		}
	}
	return x
}
func sortSlice1(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j <len(x)-1-i; j++ {
			if x[j]>x[j+1] {
				temp:=x[j]
				x[j]=x[j+1]
				x[j+1]=temp
			}
		}
	}
	return x
}

//方法1 返回多个值
func add(x int,y int) (int,int,bool) {  //返回多个值
	 return x-y,x+y,x>y
}

//方法2 返回多个值
func add1(x int,y int) (num int,num1 int,falg bool) {
	num=x+y
	num1=x-y
	falg=x>y
	return
}



func main()  {
	 Println("函数使用")
	 var A=sum(10,10)
	 Println(A)
	 var B=sub(15,5)
	 Println(B)
	 var C=sub1(1.8,1.3)
	 Println(C)
	 var sliec=[]int{12,345,546,765,87,34,56}
	 var D=sortSlcie(sliec)
	 Println(D)
     var num,num1,flag=add(10,21) //函数返回多个值
	 Printf("\nnum=%v num1=%v flag=%v \n",num,num1,flag)
	 var num2,_,flag1=add(10,21) //不需要就使用匿名变量接收
	 Printf("num=%v  flag=%v \n",num2,flag1)
	var sliec1=[]int{12,345,546,765,87,34,56}
	var E=sortSlice1(sliec1)
    Println(E)

}