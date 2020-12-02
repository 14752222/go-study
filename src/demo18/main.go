package main

import (
	"sync"
	"time"
   . "fmt"
)

var wg sync.WaitGroup

func main() {
	Println("单向管道， 管道多路复用")
	Println("------方法1------")
	fn()
	Println("------方法2------")
	fn2()
	Println("------方法3------")
	fn3()
}
func fn() {
	Println("单向管道")
	Println("默认的管道都是双向的，但是也可以创建单向管道，只读或只写")
	Println("var ch =make(chan interface{}) 双向管道")
	var ch = make(chan interface{}, 4)
	ch <- 10
	ch <- 20
	Println(<-ch)
	Println("var ch1=make(chan <- interface{}) 只写单向管道")

	var ch1 = make(chan<- interface{}, 2)
	ch1 <- 10
	ch1 <- 20
	Println("var ch1=make( <-chan interface{}) 只都单向管道")
	var ch2 = make(<-chan interface{})
	Println(ch2)
}
func fn2() {
	Println("单向管道使用案例")
	var ch = make(chan interface{}, 20)
	wg.Add(1)
	go f(ch)
	wg.Add(1)
	go f1(ch)
	wg.Wait()
}

func f(ch chan<- interface{}) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 20; i++ {
		ch <- i
		Println("写入成功", i)
		time.Sleep(time.Millisecond * 100)

	}
}
func f1(ch <-chan interface{}) {
	defer wg.Done()
	for v := range ch {
		Println("读取成功", v)
		time.Sleep(time.Millisecond * 100)
	}
}

func fn3() {
	Println("select多路管道复用 有点类似switch语句 它有case分支和一个默认分支")
	Println("select使用不需要关闭管道")
	var ch = make(chan interface{}, 10)
	var ch1 = make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	for i := 0; i < 10; i++ {
		ch1 <- Sprintf("字符串： %d", i)
	}

	for {
		select {
		case v := <-ch:
			Println("ch", v)
		case v := <-ch1:
			Println("ch1", v)
		default:
			Println("数据获取完毕")
			return
		}
	}

}
