package main

import (
	"sort"
	. "fmt"
	)

func main(){
	Println("选择排序")
	var sliceA=[]int{1,45,6678,980,564,6547,56,6789}
	for i := 0; i <len(sliceA) ; i++ {
		for j := i+1; j < len(sliceA); j++ {
			 if sliceA[i]>sliceA[j]{
			 	swit:=sliceA[i]
			 	sliceA[i]=sliceA[j]
			 	sliceA[j]=swit
			 }
		}
	}
	Println(sliceA)
	Println("冒泡排序")
   var sliceB=[]int{1,45,6678,980,564,6547,56,6789}
	for i := 0; i < len(sliceB); i++ {
		for j := 0; j <len(sliceB)-1-i; j++ {
           if sliceB[j]>sliceB[i+1]{
           	temp:=sliceB[j]
           	sliceB[j]=sliceB[j+1]
           	sliceB[j+1]=temp
		   }
		}
	}
	Println(sliceB)
	Println("sort排序")
	var sliceC=[]int{1,45,6678,980,564,6547,56,6789}
	var sliceD=[]float64{1.3,45.3,66.1,98.0,56.4,65.47,5.6,6.789}
	var sliceE=[]string{"A","C","D","G","Q","S","M","L"}
    sort.Ints(sliceC)
	sort.Float64s(sliceD)
	sort.Strings(sliceE)
	Println("默认升序排列")
	Println(sliceC)
	Println(sliceD)
	Println(sliceE)
	Println("//降序排列")
	sort.Sort(sort.Reverse(sort.IntSlice(sliceC)))
	sort.Sort(sort.Reverse(sort.Float64Slice(sliceD)))
	sort.Sort(sort.Reverse(sort.StringSlice(sliceE)))

	Println(sliceC)
	Println(sliceD)
	Println(sliceE)

}