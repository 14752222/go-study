package main

import (
	. "fmt"
	"sync"
	"time"
	)
var wg sync.WaitGroup  //协程等待
var Mute sync.Mutex  //互斥锁基类
var Rmute sync.RWMutex //读写互斥锁
var count=0
func main()  {
	defer wg.Wait()
	Println("互斥锁 同一时间只有一个协程在查找资源")
	Println("go build -race main.go 编译运行查看是否有资源机战")
	Println("----方法1-----")

	for i := 0; i <20 ; i++ {
		wg.Add(1)
		go fn()
	}
	Println("----方法2-----")
	for i := 1; i < 20; i++ {
	 wg.Add(1)
	 go fn1(i)
	}
	Println("----方法3-----")
	Println("读写互斥锁")

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go fn2()
	}
	Println("----方法4-----")
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go fn3()
	}
}

func fn()  {
	Println("Mute.Unlock()//上锁")
	Mute.Lock() //上锁
	count++
	Println("count",count)
	Println("Mute.Unlock()//开锁")
	time.Sleep(time.Second)
	wg.Done()
	Mute.Unlock()//开锁

}
var  m= make(map[int]int,0)
func  fn1(num int)  {
	defer wg.Done()
	Mute.Lock()
     sum:=1
	for i := 1; i <num ; i++ {
          sum*=i
	}
	m[num]=sum
	Println("m",m)
	time.Sleep(time.Second)
	Mute.Unlock()

}

func fn2()  {
	Mute.Lock()
	Println("写入操作")
	time.Sleep(time.Second)
	Mute.Unlock()
	wg.Done()
}
func fn3()  {
	Println("读写互斥锁")
	Rmute.RLock()
	Println("写入操作")
	time.Sleep(time.Second)
	Rmute.RUnlock()
	wg.Done()
}