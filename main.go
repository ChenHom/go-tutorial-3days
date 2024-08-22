package main

import (
	"fmt"
	"time"
)

// producer 將數據發送到通道
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i                            // 將數據發送到通道
		time.Sleep(100 * time.Millisecond) // 模擬生產耗時
	}
	close(ch) // 生產完成後關閉通道
}

// consumer 從通道接收數據
func consumer(ch chan int) {
	for num := range ch { // 從通道中接收數據
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(150 * time.Millisecond) // 模擬消耗耗時
	}
}

func main() {
	// 創建一個無緩衝的通道
	ch := make(chan int)

	// 並行執行 producer 和 consumer
	go producer(ch)
	go consumer(ch)

	// 等待 Goroutines 完成
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
