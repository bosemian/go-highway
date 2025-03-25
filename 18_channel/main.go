package main

import (
	"fmt"
	"sync"
	"time"
)

func getBurger(ch chan string) {
	time.Sleep(time.Second)
	ch <- "🍔 Burger delivered"
}

func getPackage(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	ch <- "📦 Package delivered"
}

func getPizza(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	ch <- "🍕 Pizza delivered"
}

func main() {
	// baseChan()
	// grabMany()
	syncM()
}

func baseChan() {
	ch := make(chan string)

	go getBurger(ch)

	msg := <-ch
	fmt.Println(msg)
}

func grabMany() {
	ch := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go getPizza(ch, &wg)
	go getPackage(ch, &wg)

	// Get first message (but note you have 2 goroutines sending to the channel)
	for i := 0; i < 2; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

	// time.Sleep(time.Second)
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No message yet!")
	// }

	wg.Wait()
}

func getSyncBurger(ch chan string, m *sync.Mutex) {
	m.Lock()
	ch <- "🍔 Burger delivered"
	m.Unlock()
}

func getSyncPackage(ch chan string, m *sync.Mutex) {
	m.Lock()
	ch <- "📦 Package delivered"
	m.Unlock()
}

func getSyncPizza(ch chan string, m *sync.Mutex) {
	m.Lock()
	ch <- "🍕 Pizza delivered"
	m.Unlock()
}

func syncM() {
	start := time.Now()
	var mu sync.Mutex
	ch := make(chan string, 3)

	go getSyncBurger(ch, &mu)
	go getSyncPizza(ch, &mu)
	go getSyncPackage(ch, &mu)

	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
	fmt.Println("End with ", time.Since(start), " seconds")
}
