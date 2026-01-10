package main

import (
	"fmt"
	"time"
)

func printNumbers(num1 int, num2 int) {
	var biggerNum, smallerNum int = max(num1, num2), min(num1, num2)
	for i := smallerNum; i < biggerNum; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters(){
	for c := 'a'; c <= 'z'; c++ {
		fmt.Printf("%c\n", c)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	start := time.Now()
	printNumbers(1, 5)
	printLetters()
	fmt.Println("Took: ", time.Since(start), " seconds to run")

	done := make(chan bool)
	
	start = time.Now()
	go func ()  {
		go printNumbers(1, 5)
		done <- true
	}()

	go func() {
		printLetters()
		done <- true
	}()

	<-done
	<-done

	fmt.Println("Took: ", time.Since(start), " seconds to run")

}