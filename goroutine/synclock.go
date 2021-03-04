package main

import (
	"fmt"
	"sync"
	"time"
)

var x uint64
var wgg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wgg.Done()
}
func write() {
	rwlock.Lock()
	x += 1
	fmt.Println("start write...")
	time.Sleep(10 * time.Microsecond) // 假设写操作耗时时长
	fmt.Println("end write")
	rwlock.Unlock()

	wgg.Done()
}

func read() {
	rwlock.Lock()
	fmt.Println("start read...")
	time.Sleep(time.Microsecond) // 假设读操作耗时时长
	fmt.Println("end read")
	rwlock.Unlock()

	wgg.Done()
}
func main() {
	// 互斥锁
	//wgg.Add(2)
	//go add()
	//go add()
	//wgg.Wait()
	//fmt.Println(x)

	// 读写互斥锁，适合读多写少的操作
	for i := 0; i <= 10; i++ {
		wgg.Add(1)
		go write()
	}

	for i := 0; i <= 10; i++ {
		wgg.Add(1)
		go read()
	}

	wgg.Wait()
}
