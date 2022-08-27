package main

import (
	"fmt"

	"github.com/lu4p/cat"
)

func main() {
	txt, _ := cat.File("sample.pdf")
	fmt.Println(txt)

	// s2, err := strconv.Unquote(`"` + txt + `"`)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("normal", s2)
}
