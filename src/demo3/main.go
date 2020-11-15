package main

import (
	. "fmt"
	"time"
)

func main()  {
	Println("Time包使用，日期函数")
   var Date=time.Now()
	Println(Date)//2020-11-15 14:39:11.1496994 +0800 CST m=+0.003474901
   var year= Date.Year()
   var month= Date.Month()
   var day= Date.Day()
   var hour= Date.Hour()
   var minute= Date.Minute()
   var second= Date.Second()
   Printf("\n%d-%d-%d %d:%d:%d",year,month,day,hour,minute,second) //2020-11-15 15:22:8
   Printf("\n%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)//2020-11-15 15:22:08
   Println("%02d 时间少于两位数 在前面补上0")
   Println("2006 表示年")
   Println("01 表示月")
   Println("02 表示日")
   Println("03 表示时 如果需要输出24小时制 应该写15")
   Println("04 表示分")
   Println("05 表示秒")
   var str= Date.Format("2006-01-02 03:04:05")
   var str1= Date.Format("2006-01-02 15:04:05")
   Println(str) //2020-11-15 03:28:37
   Println(str1)//2020-11-15 15:28:37
   Println("获取当前时间戳 毫秒时间戳")
   unixTime:=Date.Unix()
   Println(unixTime)
   Println("获取当前时间戳 纳秒时间戳")
   unixNaTime:=Date.UnixNano()
   Println(unixNaTime)

   var unixTime1 =1605425787
   time1:=time.Unix(int64(unixTime1),0) //需要转换int64
   //time.Unix(int64(毫秒时间戳),int64(纳秒时间戳)) 需要那个写那个 不需要的就写0
   Println("//time.Unix(int64(毫秒时间戳),int64(纳秒时间戳)) 需要那个写那个 不需要的就写0")
   str2:=time1.Format("2006/01/02 ->03:04:05")
   Println(str2)//2020/11/15 ->03:36:27
   tiemStr()
   tiemConst()
   timeAdd()
   timeTicker()
   timeSleep()
   timeSleep1()
}

func tiemStr()  {
	Println("日期字符串转换时间戳")
	var str="2020-11-15 03:44:26" //1605383066
	var tep="2006-01-02 15:04:05"
	Println("模板应该同 需要转换的时间字符串一样，不然会报错")
	var timeObj,_= time.ParseInLocation(tep,str,time.Local)
	Println(timeObj)
	Println(timeObj.Unix())

}

func tiemConst()  {
	Println("golang 提供的时间间隔常量")
	Println(time.Microsecond) //毫秒
	Println(time.Second)
	Println(time.Nanosecond)
	Println(time.Minute)
	Println(time.Hour)
}

func timeAdd()  {
	Println("增加时间")
	var timeObj=time.Now()
	Println(timeObj)
    timeObj=timeObj.Add(time.Hour)
	Println(timeObj)

}

func timeTicker()  {
	Println("定时器")
	ticKer:=time.NewTicker(time.Second)//定义一个1秒间隔的定时器
	num:=0
	for i := range ticKer.C { //写俩个参数会报错 这里只返回一个参数
		Println(i)
		num++
		if num>5{
			ticKer.Stop() //停止定时器 并且在内存中销毁
			return
		}
	}
}
func timeSleep()  {
	Println("time.Sleep(time.Second) 表示休眠1秒后执行")
	time.Sleep(time.Second*3)
	Println("aaa1")
	time.Sleep(time.Second)
	Println("aaa2")
	time.Sleep(time.Second)
	Println("aaa3")
}
func timeSleep1(){
	Println("定时器方法2")
	var num=0
	for {
		time.Sleep(time.Second)
		num++
		Println(num)
		if num>10{
           break
		}
	}
}