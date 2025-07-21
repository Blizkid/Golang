package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 100)

	// 启动生成协程
	go generateInt(intChan)

	// 启动拉取协程
	go pullInt(intChan)
	// 主协程等待
	//fmt.Println("主协程等待拉取协程完成") // 没有占用主线程时间就会直接结束
	time.Sleep(2 * time.Second) // 等待协程完成
	fmt.Println("主协程完成")
}

func generateInt(intChan chan int) {
	for i := 0; i < 200; i++ {
		intChan <- i
		fmt.Println("生成协程生成整数:", i)
	}
	fmt.Println("生成协程完成")
	close(intChan)
}

func pullInt(intChan chan int) {
	for i := 0; i < 200; i++ {
		fmt.Println("拉取协程拉取整数:", <-intChan)
	}
	fmt.Println("拉取协程完成")
}
