package main 

import . "fmt"

func main(){
  Println("切片使用")
  Println("基于数组，可拓展长度的，不需要自定义数组长度")
  Println("切片不能用赋值数组的方法来添加数据")
  Println("方法1")
  var arr =[]int{1,23,43,455,656,67}
  Printf("\narr类型：%T arr=%v 长度:%v",arr,arr,len(arr))
  Println("方法2")
  arr1:=[]int{1,23,4345,54}
  Printf("\narr类型：%T arr=%v 长度:%v",arr1,arr1,len(arr1))
  Println("方法3")
  arr2:=[]int{1:11,2:23,3:43,4:45,5:54}
  Printf("\narr类型：%T arr=%v 长度:%v",arr2,arr2,len(arr2))
 Println("----------")
 var arr3 []int
  Printf("\narr类型：%T arr=%v 长度:%v",arr3,arr3,len(arr3))
  Println("切片的默认值是nil")
  Println("---遍历切片---")
  Println("遍历方法1")
  for k, v := range arr1 {
     println("键值",k,"值",v)
  }
  Println("遍历方法2")
  for i := 0; i < len(arr1); i++ {
    println("键值",i,"值",arr1[i])
  }
  Println("注意！！！")
  var arr4=[]int{1,23,4,45,456,57}
  arrc:=arr4[:]
  Println("arr4",arr4)
  Println("arrc",arrc)
  Println("arrc:=a[:]相当于获取a数组里面的所有值")
  c:=arr4[1:4] // [23 4 45]
  Println("c",c)
  Println("c:=a[1:4] 左边包括，右边不包括 有点像 javascript数组的slice()方法")
  Println("c:=a[1:4]相当于获取a数组里面的所有值")
  d:=arr4[2:] //[4 45 456 57]
  Println("d",d)
  Println(" d:=arr4[2:] 表示从切片第二位数开始 截取后面所有")
  e:=arr4[:3] //[1 23 4]
  Println(" e:=arr4[:3] 表示获取切片索引3以前的数据")
  Println("e",e)
  Println("---基于切片----")
  var Arr=[]int{1,2,3,4,5,6,7}
  A:=Arr[1:]
  Println("A",A)
  Println("切片拥有自己的长度和容量 可以使用len函数获取长度 使用内置的cap函数求切片的容量 \n")
  Println("注意：！！！ 容量是切片第一个元素到到底层数组元素末尾的个数 \n")
  Printf("\nArr=%v A类型：%T 长度：%v 容量：%d \n",Arr,Arr,len(Arr),cap(Arr)) // A=[1 2 3 4 5 6 7] A类型：[]int 长度：7 容量：7
  Printf("\nA=%v A类型：%T 长度：%v 容量：%d \n",A,A,len(A),cap(A)) // A=[2 3 4 5 6 7] A类型：[]int 长度：6 容量：6
  B:=Arr[1:3]
  Println("B",B) // B=[2 3]
  Printf("\n B=%v B类型：%T 长度：%v 容量：%d \n",B,B,len(B),cap(B)) // A=[2 3] A类型：[]int 长度：2 容量：6
    Println("注意：！！！ 容量是切片第一个元素到到(底层数组)元素末尾的个数 \n")
    Println("注意：！！！ 容量是切片第一个元素到到(底层数组)元素末尾的个数 \n")
    Println("注意：！！！ 容量是切片第一个元素到到(底层数组)元素末尾的个数 \n")
  Println("golang切片的本质是对底层的数组的封装，它包含三个信息：底层数组的指针，切片的长度，切片的容量")


}