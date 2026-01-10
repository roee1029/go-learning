package main

import (
	"fmt"
	"time"
)

func printNumbers(num1 int, num2 int, ch chan int) {
	var biggerNum, smallerNum int = max(num1, num2), min(num1, num2)
	for i := smallerNum; i < biggerNum; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func printLetters(ch chan rune){
	for c := 'a'; c <= 'z'; c++ {
		ch <- c
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func main(){
	chInt := make(chan int)
	chStr := make(chan rune)

	go printNumbers(1, 5, chInt)
	go printLetters(chStr)

	for {
		select{
		case msg1, ok := <-chInt:
			if ok{
				fmt.Println(msg1)
			}else{
				chInt = nil
			}
			
		case msg2, ok := <-chStr:
			if ok{
				fmt.Printf("%c\n", msg2)
			}else{
				chStr = nil
			}
			
		}
		if chInt == nil && chStr == nil{
			fmt.Println("All channels closed, exiting...")
			return
		}
		
	}
		


}