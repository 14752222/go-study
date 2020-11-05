package main

import . "fmt"
func main() {
	//println(123)
	Println("make函数创建切片")
	var slice = make([]int, 4, 8)
	Printf("值%v 长度%v 容量%v %T", slice, len(slice), cap(slice), slice)
	Println("修改切片的值")
	slice[0] = 10
	slice[1] = 11
	slice[2] = 12
	slice[3] = 13
	Printf("\n值%v 长度%v 容量%v %T", slice, len(slice), cap(slice), slice)
	var slice1 []int //值[] 长度0 容量0 []int
	Println("初始化未赋值，然后直接使用索引赋值会报错，需要使用append方法来赋值，append会自动扩容 倍数扩容")
	Printf("\n值%v 长度%v 容量%v %T", slice1, len(slice1), cap(slice1), slice1)
	slice1=append(slice1,12) //值[12] 长度1 容量1 []int
	Printf("\n值%v 长度%v 容量%v %T", slice1, len(slice1), cap(slice1), slice1)
	slice1=append(slice1,24)
	slice1=append(slice1,48,96)
	Println("追加多个append(slice1,1,2,3,45)")
	slice1=append(slice1,1,2,3,45)
	Printf("\n值%v 长度%v 容量%v %T", slice1, len(slice1), cap(slice1), slice1)
    println("append方法合并切片")
    var	slice2 []int
    var	slice3 []int
   slice2=append(slice2,1,2,3,4)
   slice3=append(slice3,5,6,7,8)
	Printf("\n值%v 长度%v 容量%v %T", slice2, len(slice2), cap(slice2), slice2)
	Printf("\n值%v 长度%v 容量%v %T", slice3, len(slice3), cap(slice3), slice3)
   Println("append 方法 合并个切片")
   Println("合并切片固定语法 append(slice2,slice3...)")
   slice2=append(slice2,slice3...)
	Printf("\n值%v 长度%v 容量%v %T", slice2, len(slice2), cap(slice2), slice2)
	Println("\n切片扩容策略 按倍数扩容")
   var slice4 []int
	for i := 0; i <10 ; i++ {
		slice4=append(slice4,i)
		Printf("\n值%v 长度%v 容量%v %T", slice4, len(slice4), cap(slice4), slice4)

	}
	Println("\n")
	Println("切片是引用类型")
	var cSlice=[]int{1,2,3,4}
	Printf("\n 未修改cSlice值%v 长度%v 容量%v %T", cSlice, len(cSlice), cap(cSlice), cSlice)
    var cSlice1=cSlice
    cSlice1[0]=10
	Printf("\n cSlice值%v 长度%v 容量%v %T", cSlice, len(cSlice), cap(cSlice), cSlice)
	Printf("\n cSlice1值%v 长度%v 容量%v %T", cSlice1, len(cSlice1), cap(cSlice1), cSlice1)

    Println("copy函数使用 copy(需要复制的切片,被复制的切片)")
    var cSlice2=make([]int,4,4)
    copy(cSlice2,cSlice)
	Printf("\n cSlice值%v 长度%v 容量%v %T", cSlice, len(cSlice), cap(cSlice), cSlice)
	Printf("\n cSlice2值%v 长度%v 容量%v %T", cSlice2, len(cSlice2), cap(cSlice2), cSlice2)
   cSlice2[1]=999
	Printf("\n 重新赋值后的cSlice2值%v 长度%v 容量%v %T", cSlice2, len(cSlice2), cap(cSlice2), cSlice2)
	Printf("\n 重新赋值后的cSlice值%v 长度%v 容量%v %T", cSlice, len(cSlice), cap(cSlice), cSlice)
    Println("\n切片删除元素 golang中没有删除元素的方法")
    var sliceA=[]int{1,2,34,54,56,7,8,89}
	Printf("\n  sliceA初始值%v 长度%v 容量%v %T", sliceA, len(sliceA), cap(sliceA), sliceA)
	Println("\n使用append方法删除 !!!append 合并切片第二个元素后面加上(...)")
    sliceA=append(sliceA[:2],sliceA[3:]...)
	Printf("\n 删除索引为2SliceA值%v 长度%v 容量%v %T", sliceA, len(sliceA), cap(sliceA), sliceA)

   str :="你好golang"
   runStr:=[]rune(str)
	Printf("\n runStr值%v 长度%v 容量%v %T", runStr, len(runStr), cap(runStr), runStr)
	Printf("\n str值%v 长度%v  %T", str, len(str) , str)

	runStr[1]='?'
	Printf(string(runStr))


}
