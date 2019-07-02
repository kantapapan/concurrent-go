package main

import (
	"fmt"
	"time"
)

func main() {

	x := 1000
	for x > 0 {
		x--

		go func() {
			fmt.Println("1")
		}()

		go func() {
			fmt.Println("2")
		}()

		go func() {
			fmt.Println("3")
		}()

		go func() {
			fmt.Println("4")
		}()

		go func() {
			fmt.Println("5")
		}()
	}

	//sleep
	time.Sleep(3 * time.Second)

}
