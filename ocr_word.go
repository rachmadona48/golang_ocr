package main

import (
	"fmt"

	"github.com/lu4p/cat"
)

func main() {
	txt, _ := cat.File("sample2.docx")
	fmt.Println(txt)
}
