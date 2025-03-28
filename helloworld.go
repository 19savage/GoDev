package main

// 这是一段未来的Golang框架试用技术样例。 希望在不远的将来我能够从容地实现RPC框架下的众多算法分发和安全模式设计的对应内容。
// 2025-03-28 AnJayZhou ——TenCent. 我不知道我的未来会变成什么样子，但我希望未来我在入职之后不要再变成简单的面向Copilot编程了
// 在接下来的这一段时间中，我将基于Golang进行一组简单的Golang框架内容设计和利用。



// Golang的并行化，协程章节。

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Golang并发编程示例")

	// 1. 使用基本的goroutine
	go func() {
		fmt.Println("这是一个并发执行的goroutine")
	}()

	// 2. 使用WaitGroup等待多个goroutine完成
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutine %d 正在执行\n", id)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("goroutine %d 执行完毕\n", id)
		}(i)
	}
	wg.Wait()

	// 3. 使用channel进行goroutine间通信
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "通过channel发送的消息"
	}()
	msg := <-ch
	fmt.Println(msg)

	// 4. 使用select处理多个channel
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 500)
		ch1 <- "来自channel 1的消息"
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- "来自channel 2的消息"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	// 5. 使用互斥锁保护共享资源
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
			fmt.Printf("计数器当前值: %d\n", counter)
		}()
	}

	// 等待所有goroutine完成
	time.Sleep(time.Second)
	fmt.Println("并发示例执行完毕")
}
