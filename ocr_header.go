package main

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	content, err := readPdf_header("sample.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf_header(path string) (string, error) {
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
		line_all := 0
		line_cek := 0
		line_vendor := 0
		line_deliver_to := 0
		line_apd := 0
		line_payment_term := 0
		line_req_dept := 0
		line_ppn := 0

		for _, row := range rows {
			for _, word := range row.Content {
				if no_cek == 1 {
					line_cek++
					wdata := word.S

					if strings.Contains(wdata, "Vendor") == true {
						line_vendor = line_cek
					}

					if strings.Contains(wdata, "Deliver to") == true {
						line_deliver_to = line_cek
					}

					if strings.Contains(wdata, "APD No") == true {
						line_apd = line_cek
					}

					if strings.Contains(wdata, "Payment term") == true {
						line_payment_term = line_cek

					}

					if strings.Contains(wdata, "Req Dept") == true {
						line_req_dept = line_cek

					}

					if strings.Contains(wdata, "PPN") == true {
						line_ppn = line_cek

					}

				}

			}
			line_all = line_cek
			no_cek = no_cek + 1
		}
		lineCount := 0
		var po, po_type, vendor, deliver_to, payment_term, apd, req_dept, ppn []string
		for _, row := range rows {

			for _, word := range row.Content {
				if no == 1 {
					lineCount++
					www := word.S

					/* data vendor */
					if lineCount >= line_vendor && lineCount < line_deliver_to {
						vendor = append(vendor, string(word.S))
					}

					/* data deliver to */
					if lineCount >= line_deliver_to && lineCount < line_all {
						deliver_to = append(deliver_to, string(word.S))
					}

					/* data P/O */
					if strings.Contains(www, "P/O          :") == true {
						po = append(po, string(word.S))
					}

					/* data P/O Type*/
					if strings.Contains(www, "P/O Type") == true {
						po_type = append(po_type, string(word.S))
					}

					/* data payment term */
					if lineCount >= line_payment_term && lineCount < line_apd {
						payment_term = append(payment_term, string(word.S))
					}

					/* data apd */
					if lineCount >= line_apd && lineCount < line_req_dept {
						apd = append(apd, string(word.S))
					}

					/* data req dept */
					if lineCount >= line_req_dept && lineCount < line_ppn {
						req_dept = append(req_dept, string(word.S))
					}

					/* data req dept */
					if lineCount >= line_ppn && lineCount < line_vendor {
						ppn = append(ppn, string(word.S))
					}

				}

			}

			if no == 1 {
				po2 := strings.Join(po, " ")
				fmt.Println(po2)

				po_type2 := strings.Join(po_type, " ")
				fmt.Println(po_type2)

				vendor2 := strings.Join(vendor, " ")
				vendor3 := strings.ReplaceAll(vendor2, "Vendor     :", "")
				fmt.Println("data vendor", vendor3)

				deliver_to2 := strings.Join(deliver_to, " ")
				fmt.Println(deliver_to2)

				payment_term2 := strings.Join(payment_term, " ")
				fmt.Println(payment_term2)

				apd2 := strings.Join(apd, " ")
				fmt.Println(apd2)

				req_dept2 := strings.Join(req_dept, " ")
				fmt.Println(req_dept2)

				ppn2 := strings.Join(ppn, " ")
				fmt.Println(ppn2)

				fmt.Println("number of lines:", lineCount)
			}
			no = no + 1
		}

	}
	return "", nil
}
