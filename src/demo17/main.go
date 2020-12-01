package main
import(
	. "fmt"
	"sync"
)
var wg sync.WaitGroup
func T() {
	Println("管道使用")

}
func main() {
	T()
	var intChan = make(chan int, 1000)
	var parimeChan = make(chan int, 1000)
	var extChan = make(chan bool, 16)
	wg.Add(1)
	go putNum(intChan)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primNum(intChan,parimeChan,extChan)
	}
	wg.Add(1)
	go printPrimeChan(parimeChan)
	wg.Add(1)
	go func() {
		for i := 0; i <16 ; i++ {
			<-extChan
		}
		close(parimeChan)
		wg.Done()
	}()
	wg.Wait()
}
func primNum(intChan, parimeChan chan int ,extChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
           parimeChan<-num
		}
	}
	extChan<-true
	wg.Done()

}
func putNum(intChan chan int) {
	for i := 2; i < 100; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()

}

func printPrimeChan(primeChan chan int)  {
	for num := range primeChan {
        Println("素数",num)
	}
	wg.Done()

}
