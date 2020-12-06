package main

import (
	. "fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	)


func main()  {
  Println("文件操作")
  Println("os.Open()打开的文件只能是只读文件，不能写入")
  Println("-----读取文件------")
  //fn()
  //fn1()
  //fn2()
  Println("-----写入文件------")
  //fn3()
  //fn4()
  fn5()
}
func fn()  {
	Println("读取文件")
	Println("方法1")
	Println("os.Open()支持绝对路径和相对路径，文件打开后必须关闭，只读文件方法")
	file,err:= os.Open("F:\\go\\src\\demo22\\main.go")
	//file,err:= os.Open("./main.go")
	defer file.Close() //关闭文件流
	if err!=nil{
		Println("打开文件失败")
	}
	//Println(file)//&{0xc00007c780}
	var tempSlice=make([]byte,128)
	_,fileErr:=file.Read(tempSlice)
	Println("一次不能读取全部值file.Read([]byte)需要配合循环使用")
	if  fileErr!=nil{
     Println("读取失败")
	return
	}
	Println("读取成功",string(tempSlice))
	Println("循环遍历数据")
	strFn()

}
func strFn()  {
	file,err:= os.Open("F:\\go\\src\\demo22\\main.go")
	var strSlice []byte
	var tempSlice1=make([]byte,128)

	for  {
		n,fileErr:=file.Read(tempSlice1)
		if err==io.EOF{
			Println("err==io.EOF表示读取完毕")
			break
		}else if  fileErr!=nil{
			Println("读取失败")
			break
		}

		//strSlice=append(strSlice,tempSlice1...)会导致文件没有读取完
		strSlice=append(strSlice,tempSlice1[:n]...)
	}
	Println(string(strSlice))
}
func fn1()  {
	Println("方法2 通过bufio读取文件")
    file,err:=os.Open("./main.go")
    defer file.Close()
    if err!=nil{
    	Println("读取失败")
		return
	}
	reader:=bufio.NewReader(file)
	//line,err:=reader.ReadString("\n")

	Println("reader.ReadString(byte类型字符)")
	var strTemp string
	for  {
		line,err:=reader.ReadString('\n')
	if err==io.EOF{
		strTemp+=line //可能结束是还有返回值
		Println("读取完毕")
		break
	} else if err!=nil{
			Println(err)
			return
	}
	 strTemp+=line
	}
   Println(strTemp)
}
func fn2()  {
Println("方法3 通过iotil一次性读取完毕，不是前面两种方法以流的形式读取文件")
Println("适合小文件读取")
 str,err:=ioutil.ReadFile("F:/go/src/demo22/temp.txt")
if err!=nil{
	Println("读取失败")
	return
}
Println(string(str))
}
func fn3()  {
  Println("文件写入方法1")
  Println("打开文件的模式 可以写多个用|隔开")
  Println("1、os.O_WRONLY 只写")
  Println("2、os.O_CREATE 创建文件")
  Println("3、os.O_RDONLY 只读")
  Println("4、os.O_RDWR 读写")
  Println("5、os.O_TRUNC 清空")
  Println("6、os.O_APPEND 追加")
  file,err:=os.OpenFile("./text.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
  defer file.Close()
  if err!=nil{
    Println(err)
    return
  }
// 写入文件
  Println("方法1 file.WriteString(写入的字符串)")
  Println("方法2 file.Write([]byte(str) 需要传入一个byte类型的切片")
  file.WriteString("写入的字符串")
	for i := 0; i <10 ; i++ {
		file.WriteString(Sprintf("%v",i)+"写入的字符串\r\n") //记事本直接写\n是没法识别的
	}
	var str="//记事本直接写\\n是没法识别的"
	file.Write([]byte(str))



}
func fn4()  {
  Println("写入文件方法2")
  file,err:=os.OpenFile("./text.txt",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)
  defer file.Close()
  if err!=nil{
  	Println(err)
	return
  }
  writer:=bufio.NewWriter(file)
  writer.WriteString("写文件golang你好")
  Println(" writer.WriteString表示写入到缓存里面")
  writer.Flush()
  Println("writer.Flush()表示将缓存内容写入文件中")
}
func fn5()  {
 Println("写入文件方法3")
 str:="这是一段文本"
 Println("ioutil.WriteFile(./add.text(文件名),[]byte(str) (需要转换byte类型字符串),0666（权限）)")
 err:=ioutil.WriteFile("./add.text",[]byte(str),0666)
 if err!=nil{
 	Println("写入失败",err)
 	return
 }

}