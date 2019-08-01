package studygo

/*
	M 内核级线程
	P 处理器,主要用途是用来执行goroutine,里面有一个goroutine队列
	G 一个goroutine,有自己的栈，指令指针，正在等待的channel
	S 调度器，它维护有存储M和G的队列以及调度器的一些状态信息等

	有一个global runqueue, 每个P都有一个自己的队列
	如果当前的p的G处于阻塞调用，那么p就会转到其他的M上执行，阻塞的G如果返回，则这个G可能从其他的P中来运行，如果没有合适的P，则放入global runqueue
	global runqueue，所有的P都会定期的检查global runqueue并运行其中的goroutine

	猜测： 所有开启的goruntine都会先放入global runqueue， 然后所有的p都去里面拿出来放到自己的runqueue中
*/

/*
	当启动多个goroutine时，如果其中一个goroutine异常了，并且我们并没有对进行异常处理，那么整个程序都会终止，所以我们在编写程序时候最好每个goroutine所运行的函数都做异常处理，异常处理采用recover
*/

import (
	"fmt"
	"sync"
	"time"
)

// 使用sync包同步goroutine
func dohandle(syncWg *sync.WaitGroup) {
	i := 10
	fmt.Println()
	for i > 0 {
		fmt.Print("鄢佳仪，你个小小孩！\n")
		time.Sleep(time.Second)
		i--
	}
	syncWg.Done()
}

func StudyGoruntime() {
	fmt.Print("hello")
	var wg sync.WaitGroup

	{
		wg.Add(1)
		go dohandle(&wg)
	}
	wg.Wait()
}

// 通过channel实现goroutine之间的同步
func ChannelSync() {
	Exitchan := make(chan bool, 10) //声明并分配管道内存
	for i := 0; i < 10; i++ {
		go cal(i, i+1, Exitchan)
	}
	for j := 0; j < 10; j++ {
		<-Exitchan //取信号数据，如果取不到则会阻塞
	}
	close(Exitchan) // 关闭管道
}
func cal(a int, b int, Exitchan chan bool) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	Exitchan <- true
}

// goroutine 通过通道channel来通信，而不是通过共享来通信

// select-case实现非阻塞channel
func send(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
		fmt.Println("send data : ", i)
	}
}

func SelectCase() {
	resch := make(chan int, 20)
	strch := make(chan string, 10)
	go send(resch)
	strch <- "wd"
	select {
	case a := <-resch:
		fmt.Println("get data : ", a)
	case b := <-strch:
		fmt.Println("get data : ", b)
	default:
		fmt.Println("no channel actvie")

	}
}

// channel频率控制 通过time.Ticker实现
// limiter := time.Tick(time.Second*1)
