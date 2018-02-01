package main

import (
	"com/bbinsurance/util"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strings"
)

type TestData struct {
	CompanyList []TestCompany
}

type TestCompany struct {
	Name         string
	EnglishName  string
	BuildTime    int
	Website      string
	Introduction string
	Id           int64
	Img          string
}

type TestData2 struct {
	CompanyList []Company
}

type Company struct {
	Id         int64
	Name       string
	Desc       string
	ThumbUrl   string
	Flags      int64
	DetailData string
}

type Detail struct {
	EnglishName string
	BuildTime   int
	Website     string
}

func main() {
	content, _ := util.FileGetContent("./company.json")
	fmt.Println(content)
	var testData TestData
	json.Unmarshal(util.StringToBytes(content), &testData)
	var testData2 TestData2
	var companyList []Company
	for i := 0; i < len(testData.CompanyList); i++ {
		var company Company
		var detail Detail
		company.Id = testData.CompanyList[i].Id
		company.Name = testData.CompanyList[i].Name
		company.Desc = testData.CompanyList[i].Introduction
		company.ThumbUrl = testData.CompanyList[i].Img
		detail.BuildTime = testData.CompanyList[i].BuildTime
		detail.Website = testData.CompanyList[i].Website
		detail.EnglishName = testData.CompanyList[i].EnglishName
		detailData, _ := json.Marshal(detail)
		company.DetailData = util.BytesToString(detailData)
		companyList = append(companyList, company)
	}
	testData2.CompanyList = companyList
	testData2Byte, _ := json.Marshal(testData2)
	util.FilePutContent("./test-company.json", util.BytesToString(testData2Byte))
}

func process() {
	flag.Parse()
	root := "/Users/jiaminchen/Documents/保险/保险/excel/"
	xlsxFileList := getFilelist(root)
	fmt.Printf("%s\n", xlsxFileList)
	var companyNameMap = make(map[string]string)

	for i := 0; i < len(xlsxFileList); i++ {
		processXlsx(xlsxFileList[i], companyNameMap)
	}
	outputFile, _ := os.Create("./company.txt")
	defer outputFile.Close()
	for k, v := range companyNameMap {
		line := fmt.Sprintf("%s\t%s\n", k, v)
		outputFile.WriteString(line)
	}
	outputFile.Sync()
}

func processXlsx(path string, companyNameMap map[string]string) {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Printf("err %s", err)
	} else {
		for _, sheet := range file.Sheets {
			for key, row := range sheet.Rows {
				if key != 0 {
					// fmt.Printf("%s %s\n", row.Cells[4].Value, row.Cells[5].Value)
					companyNameMap[row.Cells[4].Value] = row.Cells[5].Value
				}
				// for k, cell := range row.Cells {
				// 	if key != 0 && (k == 5 || k == 4) {
				// 		fmt.Printf("%d %d %d %s ", sheet_key, key, k, cell.Value)
				// 	}
				// }
			}
		}
	}
}

func getFilelist(path string) []string {
	var xlsxFileList []string

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".xlsx") {
			xlsxFileList = append(xlsxFileList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return xlsxFileList
}
