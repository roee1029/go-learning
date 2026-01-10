package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand/v2"
)

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Worker", id, "started")
	time.Sleep(time.Second * time.Duration((rand.Float64() * 1.5) + 0.5))
	fmt.Println("Worker", id, "done")
}


func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		wg.Go(func() {
			worker(i, &wg)
		})
	}

	wg.Wait()
	fmt.Println("All workers finished")

}