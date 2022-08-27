package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	string := "10 Harris Room Hotel Agus  Soewadji            1 PRS         722,950                  722,950    20 Management Fee                              1  PRS         28,918                   28,918    30 Tiket KAI Horiuchi CN-GMR                   1 PRS         400,000                 400,000    40 Management Fee                              1  PRS         16,000                   16,000    50 Tiket KAI Martinus CN-GMR                   1 PRS         441,500                 441,500    "
	var awal, akhir int = 20, 30
	delimiter_awal := strconv.Itoa(awal) + " "
	delimiter_akhir := strconv.Itoa(akhir) + " "
	// leftOfDelimiter := strings.Split(string, delimiter_awal)[0]
	word_awal := strings.Join(strings.Split(string, delimiter_awal)[1:], delimiter_awal)
	word_akhir := strings.Join(strings.Split(word_awal, delimiter_akhir)[:1], delimiter_akhir)
	final_word := delimiter_awal + word_akhir

	// fmt.Println("Left of Delimiter: ", leftOfDelimiter)
	// fmt.Println("word_awal : ", word_awal)
	// fmt.Println("word_akhir : ", word_akhir)
	fmt.Println("final_word : ", final_word)
}
