package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf" /*go get github.com/ledongthuc/pdf*/
)

func main() {
	pdf.DebugOn = true
	// content, err := readPdf("sample.pdf") // Read local pdf file
	_, err := readPdf("sample.pdf")
	if err != nil {
		panic(err)
	}
	// ct := strings.Split(content, `\n`)
	// fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)

	words := strings.Split(buf.String(), `\n`)
	fmt.Println(words)

	return buf.String(), nil
}
