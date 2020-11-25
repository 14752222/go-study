package main

import (
	. "fmt"
	"time"
	"sync"
	"runtime"
	)

var wg sync.WaitGroup

func T()  {
	Println("进程/线程\n进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。\n\n线程是进程的一个执行实体，是 CPU 调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。\n\n一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发执行。\n并发/并行\n多线程程序在单核心的 cpu 上运行，称为并发；多线程程序在多核心的 cpu 上运行，称为并行。\n\n并发与并行并不相同，并发主要由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行，Go程序可以设置使用核心数，以发挥多核计算机的能力。\n协程/线程\n协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。\n\n线程：一个线程上可以跑多个协程，协程是轻量级的线程。\n\n优雅的并发编程范式，完善的并发支持，出色的并发性能是Go语言区别于其他语言的一大特色。使用Go语言开发服务器程序时，就需要对它的并发机制有深入的了解。")
	Println("主线程执行结束后，如果不使用休眠，进程就结束了，协程也将会结束，不管代码有没有执行完成")
}
func main()  {
	cpuNum:=runtime.NumCPU() //获取当前计算机可以使用的核心数
	Println("cpu核心数",cpuNum)
	runtime.GOMAXPROCS(cpuNum)//设置自己可以使用的cpu核心数
	T()
	wg.Add(1)//协程加1
	go test() //开启一个协程
	for i := 0; i < 10; i++ {
         Println("main",i)
		time.Sleep(time.Millisecond*50)
	}
    wg.Wait()//等待所有的记录的goroutine执行完毕
}

func test()  {
	defer  wg.Done()//表示结束当前协程
	for i := 0; i < 10; i++ {
		Println("test",i)
		time.Sleep(time.Millisecond*100)
	}
}