package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile("/Users/jiaminchen/Documents/保险/保险/excel/个人意外.xlsx")
	if err != nil {
		fmt.Printf("err %s", err)
	} else {
		for sheet_key, sheet := range file.Sheets {
			for key, row := range sheet.Rows {
				for k, cell := range row.Cells {
					if key == 0 {
						fmt.Printf("%d %d %d %s ", sheet_key, key, k, cell.Value)
					}
				}
				fmt.Printf("\n")
			}
		}
	}
}
