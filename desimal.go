package main

import (
	"fmt"
	"math"
)

func main() {
	var no float64
	no = 5
	res := math.Floor(no)
	if no == res {
		fmt.Println("data", true)
	} else {
		fmt.Println("data", false)
	}

}
