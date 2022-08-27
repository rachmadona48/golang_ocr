package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	_, err := readPdf("sample.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)
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
	var new_word []string
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		lineCount := 0

		for _, row := range rows {
			line_cek := 0
			line_seid := 0
			line_amount := 0
			for _, word1 := range row.Content {
				wdata := word1.S
				if len(wdata) > 0 {
					line_cek++
					if strings.Contains(wdata, "SEID") == true {
						line_seid = line_cek
					}

					if strings.Contains(wdata, "Amount(IDR)") == true {
						line_amount = line_cek
					}

				}

			}
			no_cek = no_cek + 1

			for _, word := range row.Content {
				www := word.S
				if len(www) > 0 {
					lineCount++
					if lineCount < line_seid && lineCount > line_amount {
						if strings.Contains(www, "__") == false {
							new_word = append(new_word, string(www))
						}
					}
				}

			}
			no = no + 1
		}

	}
	new_word2 := strings.Join(new_word, " ")
	new_word3 := strings.Join(strings.Split(new_word2, "Total")[:1], "Total")

	var start_line int = 10
	var last_line int = 780
	for i := 10; i <= last_line; {
		string := new_word3
		var awal int = i
		var akhir int = i + 10
		delimiter_awal := ""
		if i == start_line {
			delimiter_awal = strconv.Itoa(awal) + " "
		} else {
			delimiter_awal = " " + strconv.Itoa(awal) + " "
		}
		delimiter_akhir := strconv.Itoa(akhir) + " "
		final_word := ""
		if i < last_line {
			word_awal := strings.Join(strings.Split(string, delimiter_awal)[1:], delimiter_awal)
			word_akhir := strings.Join(strings.Split(word_awal, delimiter_akhir)[:1], delimiter_akhir)
			final_word = word_akhir

		} else {
			word_awal := strings.Join(strings.Split(string, delimiter_awal)[1:], delimiter_awal)
			final_word = word_awal
		}

		// fmt.Println(final_word)
		items := delimiter_awal
		strmodel := strings.Split(final_word, "    ")
		model := strmodel[0]
		words := strings.Fields(final_word)
		last := words[len(words)-4:]
		last_word := strings.Split(strings.Join(last, " "), " ")
		qty := last_word[0] + " " + last_word[1]
		price_idr, _ := strconv.Atoi(strings.ReplaceAll(last_word[2], ",", ""))
		amount, _ := strconv.Atoi(strings.ReplaceAll(last_word[3], ",", ""))
		fmt.Println("items : ", items, "| model : ", model, "| qty : ", qty, "| price_idr : ", price_idr, "| amount : ", amount)

		i = i + 10
	}

	return "", nil
}
