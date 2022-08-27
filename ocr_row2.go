package main

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

func main() {
	content, err := readPdf("sample.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	no := 1
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()

		for _, row := range rows {
			// println(">>>> row: ", row.Position)
			println("page no :", no)
			for _, word := range row.Content {
				if no == 1 {
					fmt.Println(word.S)
				}

			}
			no = no + 1
		}
	}
	return "", nil
}
