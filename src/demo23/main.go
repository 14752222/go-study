package main

import (
	. "fmt"
	"io/ioutil"
	"os"
	"io"
	)

func main()  {
	 Println("文件写入和读取，目录操作")
	 Println("-----文件复制-----")
	 fn()
    fn1()
	Println("-----创建目录-----")
    fn2()
	Println("-----删除文件/目录-----")
    fn3()
	Println("-----重命名文件-----")
	fn4()


}
func fn()  {
	Println("文件复制")
   byteStr,err :=ioutil.ReadFile("F:/go/src/demo23/add.txt")
	if err!=nil{
		Println(err)
		return
	}
   e1:=ioutil.WriteFile("F:/go/src/demo23/app.txt",byteStr,0666)
   if e1!=nil{
   	Println("复制失败",err)
	   return
   }
}
func fn1()  {
	Println("复制文件方法2 以文件流形式复制")
	var  self="F:/go/src/demo23/add.txt"
	var  to="F:/go/src/demo23/adds.txt"
   err:=copy1(self,to)
   if err==nil{
   	Println("拷贝成功")
   }else {
	   Println("拷贝失败")
   }
}
func copy(thisFile,toFile string)  error {
	byteStr,err :=ioutil.ReadFile(thisFile)
	if err!=nil{
		return err
	}
	e1:=ioutil.WriteFile(toFile,byteStr,0666)
	if e1!=nil{
		return e1
	}
	return nil
}
func copy1(thisFile,toFile string)  error {
   oldFile,e1 :=os.Open(thisFile)
   defer oldFile.Close() //不关闭可能会导致内存泄漏
   newFile,e2 :=os.OpenFile(toFile,os.O_CREATE|os.O_WRONLY,0666)
	defer newFile.Close() //不关闭可能会导致内存泄漏

	if e1!=nil{
   	return e1
   }
   if e2!=nil{
   	return e2
   }
   var temp=make([]byte,1024)
	for  {
		n,err:=oldFile.Read(temp)
		if err==io.EOF{
			break
		}
       if err!=nil{
       	 return err
	   }
		_,err1:= newFile.Write(temp[:n])
		if err1!=nil{
			return err1
		}
	}
   return nil
}
func fn2()  {
	err:=os.Mkdir("./add",0666)
	if err!=nil{
		Println("创建文件夹失败",err)
	}
	Println("os.Mkdir创建单级目录")
	err1:=os.MkdirAll("./add/aaa/ccc",0666) //创建多级目录
	Println("os.MkdirAll创建多级目录")

	if err1!=nil{
		Println("创建多级目录失败",err1)
	}

}
func fn3()  {
	Println("os.Remove os.RemoveAll都可以删除文件和目录")
	Println("os.Remove一次只能删除一个")
    err:= os.Remove("./add/aaa/ccc")
    if err!=nil{
     	Println("删除失败",err)
   }
	Println("os.RemoveAll一次只能删除多个")
	err1:= os.RemoveAll("./add")
	if err1!=nil{
		Println("删除多个文件失败",err)
	}
}
func fn4()  {
	Println("文件重命名")
	err:=os.Rename("./adds.txt","./renname.txt")
	 Println("只能在同一个盘使用")
	 if err!=nil{
	 	Println("文件重命名失败",err)
	 }
}