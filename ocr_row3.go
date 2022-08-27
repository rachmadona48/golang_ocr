package main

import (
	"fmt"
	"strings"

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

	no_cek := 1
	no := 1
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		line_cek := 0
		line_apd := 0
		line_payment_term := 0

		for _, row := range rows {
			for _, word := range row.Content {
				if no_cek == 1 {
					line_cek++
					wdata := word.S
					if strings.Contains(wdata, "APD No") == true {
						line_apd = line_cek
					}

					if strings.Contains(wdata, "Payment term") == true {
						line_payment_term = line_cek

					}

				}

			}
			no_cek = no_cek + 1
		}

		fmt.Println("line_apd:", line_apd)
		fmt.Println("line_payment_term:", line_payment_term)

		lineCount := 0
		for _, row := range rows {

			for _, word := range row.Content {
				if no == 1 {
					lineCount++

					if lineCount >= line_payment_term && lineCount < line_apd {
						fmt.Println(word.S)

					}

				}

			}
			fmt.Println("number of lines:", lineCount)
			no = no + 1
		}

	}
	return "", nil
}
