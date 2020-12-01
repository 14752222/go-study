package main

import (
	. "fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func  T()  {
Println("管道 goroutine之间的通讯方式")
	Println("管道 数据读取不到，就会等待管道")


}

func main()  {
	Println("--------方法1--------")
	T()
	Println("--------方法2--------")
	fn()
	Println("--------方法3--------")
	fn1()
	Println("--------方法4--------")
	eg()
}

func fn()  {
	Println("创建管道var ch =make(chan interface{},4)")
	var ch =make(chan interface{},4)
	Println("取值遵守 先入先出 管道也是一种类型 而且是引用类型")
	Println("ch<-10  把 10 发送的到管道")
	Println("<-ch  从管道中取值")
	ch<-10
	ch<-20
	ch<-"你好"
	data,ok:=<-ch
	if ok {
		Println("取值成功",data)
		Println(ok)
	}
	Println("管道的值",ch)
	Printf("\n管道的类型%T\n",ch)


	ch<-22
	Println(<-ch)
	Println(<-ch)
	Println(<-ch)
}

func fn1()  {
	var ch =make(chan interface{},4)
	Println("管道阻塞")
	Println("管道超出会锁死")
	ch<-10
	ch<-20
	ch<-"你好"
	//ch<-"你好"
	//ch<-"你好"
    Println("超出限制会抛出错误: fatal error: all goroutines are asleep - deadlock!")
	Println(<-ch)
	Println(<-ch)
	Println(<-ch)
	Println("Println(<-ch)//取完值,在取值也会报错:fatal error: all goroutines are asleep - deadlock! ")
	Println("可以边读边写")
	var ch1=make(chan interface{},1)
	ch1<-1
	<-ch1
	ch1<-10
	<-ch1
	ch1<-20
   Println(<-ch1)

}

type Ch chan  int
func fn2(ch1 Ch)  {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		ch1<-i

		Println("写入数据成功",i)
		time.Sleep(time.Millisecond*500)
	}
	close(ch1)
}
func fn3(ch1 Ch)  {
	defer wg.Done()
	for v := range ch1 {
		Println("读取数据成功",v)
		time.Sleep(time.Millisecond*1000)
	}
}
func eg()  {
	var ch =make(chan int,10)
	wg.Add(1)
	go fn2(ch)
	wg.Add(1)
	go fn3(ch)
	wg.Wait()
}