package main

import(
	. "fmt"
	"time"
	"sync"
	)
var wg sync.WaitGroup
func main()  {
	start:=time.Now().Unix()
	test()
	end:=time.Now().Unix()
	Println(end-start)
}

/**
统计 1到120000的素数 使用多协程
 第一个协程1-30000
 第二个协程30001-60000
 第三个协程60001-90000
 第四个协程90001-120000
**/

func test()  {
	for i:=1;i<=4;i++ {
		wg.Add(1) //开启一个子协程
		go fn(i)
	}
	wg.Wait()//等待线程执行结束
}

func fn(n int)  {
	defer  wg.Done() //等待执行结束后关闭当前协程
	for num:=(n-1)*30000+1;num<30000*n;num++ {
		if num>1{
			var flag=true
			for i := 2; i <num ; i++ {
				if num%i==0{
					flag=false
					break
				}
			}
			if flag {

			}
		}
	}
}
func fn1()  {
	Println("单线程写法")
	start:=time.Now().Unix()
	for num:=2;num<120000;num++ {
		flag:=true
		for i:=2;i<num;i++ {
			if num%i==0{
				flag=false
				break
			}
		}
		if flag {
			//Println("素数",num)
		}
	}
	end:=time.Now().Unix()
	Println(end-start)
}