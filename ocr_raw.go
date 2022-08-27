package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := `\033[1mString in bold.\033[0m`

	s2, err := strconv.Unquote(`"` + s + `"`)
	if err != nil {
		panic(err)
	}
	fmt.Println("normal", s2)
}
