package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// string to int
	i, _ := strconv.Atoi(strings.ReplaceAll("722,950", ",", ""))

	fmt.Println(i)
}
