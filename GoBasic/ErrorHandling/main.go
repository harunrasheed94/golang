package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Reads JPMC header records and data records
*/

// type JPMCRecord struct {
// 	AllRecords []map[string]string `json:"allRecords"`
// 	RecordKeys []string            `json:"recordKeys"`
// }

const (
	HEADER string = "Header" //header
	DETAIL string = "Detail" //detail
)

func main() {

	csvFile, err := OpenCSV("/Users/haroon-kalilur/Documents/financedoc.csv")
	if err != nil {
		return
	}

	err = ProcessCSV(csvFile)
	if err != nil {
		return
	}
}

// OpenCSV opens csv file
func OpenCSV(filePath string) (*os.File, error) {

	csvfile, err := os.Open(filePath)

	if err != nil {
		fmt.Errorf("OpenCSV: %w", err)
		fmt.Println("error in opening file. Error = ", err)
		return nil, err
	}
	defer csvfile.Close()
	return csvfile, nil
}

// ProcessCSV processes csv file
func ProcessCSV(csvfile *os.File) error {
	csvReader := csv.NewReader(csvfile)
	csvReader.FieldsPerRecord = -1 //FieldsPerRecords should be set to negative so that each record can have any number of fields (i.e. variable number of columns)
	//recordMap := make(map[string]interface{})
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error in reading record. Error = ", err)
		return err
	}

	//var jpmcRecord JPMCRecord
	headerRecords := [][]string{}
	detailRecords := [][]string{}

	headerRecordsCnt := 0
	detailRecordsCnt := 0
	for _, record := range records {
		//fmt.Println(record)
		if strings.HasPrefix(record[0], "HACT") {
			headerRecords = append(headerRecords, record)
			headerRecordsCnt = len(record)
		} else {
			detailRecords = append(detailRecords, record)
			detailRecordsCnt = len(record)
		}
	}

	headerKeys := make([]string, headerRecordsCnt)
	detailKeys := make([]string, detailRecordsCnt)
	for i := 0; i < headerRecordsCnt; i++ {
		key := HEADER + strconv.Itoa(i+1)
		headerKeys[i] = key
	}
	for i := 0; i < detailRecordsCnt; i++ {
		key := DETAIL + strconv.Itoa(i+1)
		detailKeys[i] = key
	}

	fmt.Println(headerKeys)
	fmt.Println(detailKeys)
	fmt.Println(headerRecords)
	fmt.Println(detailRecords)
	return nil
}
