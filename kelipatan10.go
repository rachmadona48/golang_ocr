package main

import (
	"fmt"
)

func main() {
	sumOf10 := 0
	for i := 1; i < 710; i++ {
		if i%10 == 0 {
			fmt.Println("Data perelement adalah :", "kelipatan 10")
			fmt.Println("Data no urut :", i)
			sumOf10 += i
		}
	}
	fmt.Println("Jumlah data merupakan kelipatan 10 adalah :", sumOf10)
}
