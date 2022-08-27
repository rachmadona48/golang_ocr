package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	string := "10 Harris Room Hotel Agus  Soewadji            1 PRS         722,950                  722,950    20 Management Fee                              1  PRS         28,918                   28,918    30 Tiket KAI Horiuchi CN-GMR                   1 PRS         400,000                 400,000    40 Management Fee                              1  PRS         16,000                   16,000    50 Tiket KAI Martinus CN-GMR                   1 PRS         441,500                 441,500    60 Management Fee                              1  PRS         17,660                   17,660    70 Tiket KAI Horiuchi GMR-CN                   1 PRS         160,000                 160,000    80 Management Fee                              1  PRS          6,400                    6,400    90 Tiket KAI Martinus GMR-CN                   1 PRS         167,500                 167,500   100 Management Fee                              1  PRS          6,700                    6,700   "

	var start_line int = 10
	var last_line int = 100
	for i := 10; i <= last_line; {
		var awal int = i
		var akhir int = i + 10
		delimiter_awal := ""
		if i == start_line {
			delimiter_awal = strconv.Itoa(awal) + " "
		} else {
			delimiter_awal = " " + strconv.Itoa(awal) + " "
		}

		delimiter_akhir := strconv.Itoa(akhir) + " "
		word_awal := strings.Join(strings.Split(string, delimiter_awal)[1:], delimiter_awal)
		word_akhir := strings.Join(strings.Split(word_awal, delimiter_akhir)[:1], delimiter_akhir)
		final_word := delimiter_awal + word_akhir
		fmt.Println("final_word : ", final_word, " : akhir ", akhir)
		i = i + 10
	}
}
