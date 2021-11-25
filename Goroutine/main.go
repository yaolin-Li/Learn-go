package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int32
var mtx sync.Mutex // 互斥锁
var ch = make(chan int, 1) // 这里是1，因为在最后一轮，即获得结果时，没人来拿。如果不留buff则会卡住

func say(id string) {
	time.Sleep(time.Second)
	fmt.Println("I am done! id:" + id)
	wg.Done() // 任务完成
}
func player(name string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch //怎样从通道里拿值

		if !ok { //通道关闭
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)

		if n%10 == 0 {
			//把球打飞
			fmt.Printf("%s misses the ball! %s lose\n", name, name)
			close(ch)
			return
		}
		ball++
		fmt.Printf("%s receives ball %d\n", name, ball)
		ch <- ball // 把球传给对手
	}

}
func AtomicIncCounter() {
	defer wg.Done()
	// race condition here 竞争条件 i++的不安全性，主存和线程存储的冲突
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1) // 同一个thread执行
	}

}
func UnsafeIncCounter() {
	defer wg.Done()
	// race condition here 竞争条件 i++的不安全性，主存和线程存储的冲突
	for i := 0; i < 10000; i++ {
		/*
			mtx.Lock()
			counter++
			// 这不是原子性的
			// temp := counter 	读
			// counter = temp + 1	计算
			// counter = temp		写入
			mtx.Unlock()
		*/

		atomic.AddInt32(&counter, 1) // 同一个thread执行
	}

}
func ChannelIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		count := <-ch
		count++
		ch <- count
	}
}
func main() {
	wg.Add(2) //总共有两个任务

	/*
		go func(id string) {
			fmt.Println(id)
			wg.Done()
		}("here")

		go say("world")


		ch := make(chan int, 0) // 无缓存通道

		//	ch <- 0 因为没有人接着，所以会永远卡在这

		go player("heli", ch)
		go player("lilei", ch)

		ch <- 0 // 先有人接，才能放入
	*/
	go ChannelIncCounter()
	go ChannelIncCounter()

	ch <- 0
	wg.Wait() // 等待所有的任务完成，卡住，如果不是wg不是0

	fmt.Println(<- ch)
}
