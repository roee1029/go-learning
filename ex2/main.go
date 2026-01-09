package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%6 == 0 && i%7 == 0 {
			fmt.Println("Six Seven!")
		}
		if i % 6 == 0 {
			fmt.Println("six")
		}else if i % 7 == 0 {
			fmt.Println("seven")
		}else{
			fmt.Println(i)
		}
	}
}