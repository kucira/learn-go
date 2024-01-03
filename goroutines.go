package main

import "fmt"
// import "time"
import "sync"

func first() {
	fmt.Println("first")
}

func second() {
	go first()
	fmt.Println("second")
}

func producer(ch chan int) {
  for i := 0; i < 10; i++ {
    ch <- i
  }
} 

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch := make(chan int)
		second()
		go producer(ch)
		// time.Sleep(1 * time.Second)

		for x := range ch {
			fmt.Println(x) 
		}		
		wg.Done()
	}()

	wg.Wait()
}