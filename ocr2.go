package main

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func main() {
	content, err := ReadPlainTextFromPDF("sample2.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
	fmt.Println("=======")
}

func ReadPlainTextFromPDF(pdfpath string) (text string, err error) {
	f, r, err := pdf.Open(pdfpath)
	defer f.Close()
	if err != nil {
		return
	}

	// totalPage := r.NumPage()
	// fmt.Println(totalPage)
	// fmt.Println("+++++++++++++++++++++++++++++++++++++++")

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return
	}

	// fmt.Println(b)

	// fmt.Println("=============================oooooooooooooooooooooooooo++++++++++++++++++++++++++")

	buf.ReadFrom(b)

	text = buf.String()
	return
}
